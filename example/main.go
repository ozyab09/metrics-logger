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

func FatalInit(ml go_metrics_logger.MetricsLogger, ctx context.Context) {

	for i := 0; i < 100000; i++ {
		ml.Fatal(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Mozilla"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Fatal(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Chrome"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Fatal(ctx, logger.Message{
			EventID:       12,
			OperationName: "http responce",
			ComponentName: "test",
			EventStatus:   logger.StatusError,
			Description:   "example fatal log",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Edge"}`,
			Payload:       "",
		})
	}
}

func ErrorInit(ml go_metrics_logger.MetricsLogger, ctx context.Context) {

	for i := 0; i < 100000; i++ {
		ml.Error(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Mozilla"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Error(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Chrome"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Error(ctx, logger.Message{
			EventID:       12,
			OperationName: "http responce",
			ComponentName: "test",
			EventStatus:   logger.StatusError,
			Description:   "example Error log",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Edge"}`,
			Payload:       "",
		})
	}
}

func WarnInit(ml go_metrics_logger.MetricsLogger, ctx context.Context) {

	for i := 0; i < 100000; i++ {
		ml.Warn(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Mozilla"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Warn(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Chrome"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Warn(ctx, logger.Message{
			EventID:       12,
			OperationName: "http responce",
			ComponentName: "test",
			EventStatus:   logger.StatusError,
			Description:   "example Warn log",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Edge"}`,
			Payload:       "",
		})
	}
}

func InfoInit(ml go_metrics_logger.MetricsLogger, ctx context.Context) {

	for i := 0; i < 100000; i++ {
		ml.Info(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Mozilla"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Info(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Chrome"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Info(ctx, logger.Message{
			EventID:       12,
			OperationName: "http responce",
			ComponentName: "test",
			EventStatus:   logger.StatusError,
			Description:   "example Info log",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Edge"}`,
			Payload:       "",
		})
	}
}

func DebugInit(ml go_metrics_logger.MetricsLogger, ctx context.Context) {

	for i := 0; i < 100000; i++ {
		ml.Debug(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Mozilla"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Debug(ctx, logger.Message{
			EventID:       12,
			OperationName: "http request",
			ComponentName: "test example",
			EventStatus:   logger.StatusError,
			Description:   "example error log message",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Chrome"}`,
			Payload:       "",
		})
		time.Sleep(time.Microsecond * 2000000)
		ml.Debug(ctx, logger.Message{
			EventID:       12,
			OperationName: "http responce",
			ComponentName: "test",
			EventStatus:   logger.StatusError,
			Description:   "example Debug log",
			TS:            time.Now(),
			Latency:       time.Microsecond,
			Headers:       `{"Edge"}`,
			Payload:       "",
		})
	}
}

func main() {

	ctx := context.Background()
	customLogger := &CustomLogger{}
	ml := go_metrics_logger.NewMetricsLogger(
		ctx,
		customLogger,
		"/metrics",
		":8085",
	)
	go FatalInit(ml, ctx)

	go DebugInit(ml, ctx)

	go InfoInit(ml, ctx)

	go ErrorInit(ml, ctx)

	go WarnInit(ml, ctx)

	// Sleep to check metrics
	time.Sleep(time.Minute * 20000)
}
