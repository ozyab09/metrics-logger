package logger

import (
	"log"
	"sync"
)

// Logger - message logging interface
type Logger interface {
	// Log - message logging
	Log(level string, message string)
	// Close - close logger
	Close() error
}

// DefaultLogger - default message logger implementation
type DefaultLogger struct {
	mu sync.Mutex
}

// NewDefaultLogger - return new default logger
func NewDefaultLogger() Logger {
	return &DefaultLogger{}
}

// Log - log message in console
func (d *DefaultLogger) Log(level string, message string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	log.Printf("[%s] %s\n", level, message)
}

// Close - close logger
func (d *DefaultLogger) Close() error {
	return nil
}
