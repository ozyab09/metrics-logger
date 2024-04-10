# Go-metrics-logger
Go-metrics-logger is a 
logging library in Go that allows you to 
select a logging location and also generates
metrics based on the message.
## Features
- Allows to log messages with different logging levels
- Generates metrics based on transmitted messages
- Provides several logging methods, and also allows users to make their own custom logger and use it
- Receives only the necessary fields in the message and itself converts the structure into a logging format
- Sends metrics on a given port, making it possible to build graphs
- Metrics contain many labels, making it possible to build a variety of graphs
## Install
```shell
go get github.com/ozyab09/metrics-logger/go_metrics_logger@latest
```
## Examples
### Log messages in console
```go
package main

import (
	"context"
	"time"

	"github.com/ozyab09/metrics-logger/go_metrics_logger"
	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
)

func main() {
	ctx := context.Background()

	// create metrics logger with default logger which log essages in console
	// and forward metrics in port 8081 and path /metrics
	ml := go_metrics_logger.NewMetricsLogger(ctx, logger.NewDefaultLogger(), "/metrics", ":8081")

	// log message with level info
	ml.Info(ctx, logger.Message{
		EventID:       12,
		OperationName: "http request",
		ComponentName: "test example",
		EventStatus:   logger.StatusSuccess,
		Description:   "example log message",
		TS:            time.Now(),
		Latency:       time.Second,
		Headers:       `{"Accept-Encoding": "gzip, deflate, br", "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"}`,
		Payload:       "",
	})

	// Sleep to have time to check metrics
	time.Sleep(time.Minute * 2)
}
```
### Log messages into file
```go
package main

import (
	"context"
	"os"
	"time"

	"github.com/ozyab09/metrics-logger/go_metrics_logger"
	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
)

func main() {
	ctx := context.Background()

	// create logger which logs mesages into file
	fileLogger, err := logger.NewFileLogger("logs.txt")
	if err != nil {
		os.Exit(-1)
	}

	// create metrics logger with file logger which log messages into file
	// and forward metrics in port 8081 and path /metrics
	ml := go_metrics_logger.NewMetricsLogger(
		ctx,
		fileLogger,
		"/metrics",
		":8081",
	)

	// log message with level error
	ml.Error(ctx, logger.Message{
		EventID:       12,
		OperationName: "http request",
		ComponentName: "test example",
		EventStatus:   logger.StatusError,
		Description:   "example error log message",
		TS:            time.Now(),
		Headers:       `{"Accept-Encoding": "gzip, deflate, br", "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"}`,
		Payload:       "",
	})

	// Sleep to check metrics
	time.Sleep(time.Minute * 2)
}
```
### Log messages with custom logger implementation
Implement custom logger which throttle messages with low logging level
```go
package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ozyab09/metrics-logger/go_metrics_logger"
	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
)

type CustomLogger struct {
	mu sync.Mutex
}

func (c *CustomLogger) Log(level string, message string) {
	if level == go_metrics_logger.ErrorLevel || level == go_metrics_logger.FatalLevel {
		log.Printf("[%s]: %s", level, message)
	}
}

func (c *CustomLogger) Close() error {
	return nil
}

func main() {
	ctx := context.Background()

	// use custom logger which throttle messages with
	// low logging level
	customLogger := &CustomLogger{}

	// create metrics logger with custom logger which
	// forward metrics in port 8081 and path /metrics
	ml := go_metrics_logger.NewMetricsLogger(
		ctx,
		customLogger,
		"/metrics",
		":8081",
	)

	// log message with level error
	ml.Fatal(ctx, logger.Message{
		EventID:       12,
		OperationName: "http request",
		ComponentName: "test example",
		EventStatus:   logger.StatusError,
		Description:   "example error log message",
		TS:            time.Now(),
		Latency:       time.Microsecond,
		Headers:       `{"Accept-Encoding": "gzip, deflate, br", "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"}`,
		Payload:       "",
	})

	// Sleep to check metrics
	time.Sleep(time.Minute * 2)
}

```