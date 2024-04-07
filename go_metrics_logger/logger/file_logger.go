package logger

import (
	"fmt"
	"os"
	"sync"
)

type FileLogger struct {
	mu sync.Mutex
	f  *os.File
}

func NewFileLogger(fileName string) (Logger, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return &FileLogger{f: file}, nil
}

func (d *FileLogger) Log(level string, message string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.f.WriteString(fmt.Sprintf("[%s], %s\n", level, message))
	if err != nil {
		fmt.Println(err)
	}
}

func (d *FileLogger) Close() error {
	return d.f.Close()
}
