package example

import (
	"encoding/json"

	"net/http"

	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
	"github.com/ozyab09/metrics-logger/go_metrics_logger/metrics"
)

type Data struct {
	Message string `json:"message"`
}

type FileData struct {
	Message string `json:"message"`
	File    string `json:"file"`
}

func main() {

	logger := logger.NewDefaultLogger()
	//fileLogger, err := log.NewFileLogger("output.log")

	metrics := metrics.NewMetrics()

	http.HandleFunc("/handle1", func(w http.ResponseWriter, r *http.Request) {
		var data Data
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			//log.Printf("Error handling /handle1: %v", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		//log.Printf("/handle1 received: %s", data.Message)
	})

	http.HandleFunc("/handle2", func(w http.ResponseWriter, r *http.Request) {
		var data Data
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	})

	//log.Println("Server started on :8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
