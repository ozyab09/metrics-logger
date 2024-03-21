package go_metrics_logger

import (
	"fmt"
	"sync"
)

type Logger interface {
	Log(level string, message string)
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
	fmt.Printf("[%s] %s", level, message)
}
