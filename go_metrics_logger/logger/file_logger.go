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
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return &FileLogger{f: file}, nil
}

func (d *FileLogger) Log(level string, message string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.f.WriteString(fmt.Sprintf("[%s], %s", level, message))
}
