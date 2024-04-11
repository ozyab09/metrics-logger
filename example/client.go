package example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func sendData(wg *sync.WaitGroup, data []byte, url string) {
	defer wg.Done()
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error sending data to %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Sent data to %s, response status: %s\n", url, resp.Status)
}

func InitRequest() {
	var wg sync.WaitGroup
	urls := []string{"http://localhost:8080/handle1", "http://localhost:8080/handle2"}

	data := Data{Message: "Hello from client!"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling ", err)
		return
	}

	for _, url := range urls {
		for i := 0; i < 5; i++ { // Для примера каждый обработчик вызывается 5 раз
			wg.Add(1)
			go sendData(&wg, jsonData, url)
		}
	}

	wg.Wait()
	fmt.Println("All data sent")
}
