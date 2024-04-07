package logger

import (
	"log"
	"sync"
)

type Logger interface {
	Log(level string, message string)
	Close() error
}

type DefaultLogger struct {
	mu sync.Mutex
}

func NewDefaultLogger() Logger {
	return &DefaultLogger{}
}

func (d *DefaultLogger) Log(level string, message string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	log.Printf("[%s] %s\n", level, message)
}

func (d *DefaultLogger) Close() error {
	return nil
}
