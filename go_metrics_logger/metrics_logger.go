package go_metrics_logger

import (
	"context"
	"encoding/json"

	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
	"github.com/ozyab09/metrics-logger/go_metrics_logger/metrics"
)

const (
	loggerWorkersNum = 512
	loggerBufferSize = 65536
)

const (
	DebugLevel = "DEBUG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARN"
	ErrorLevel = "ERROR"
	FatalLevel = "FATAL"
)

type (
	MetricsLogger interface {
		Debug(ctx context.Context, message logger.Message)
		Info(ctx context.Context, message logger.Message)
		Warn(ctx context.Context, message logger.Message)
		Error(ctx context.Context, message logger.Message)
		Fatal(ctx context.Context, message logger.Message)
		Close() error
	}

	metricsLoggerImpl struct {
		exec logger.MessageExecutor
		m    metrics.Metrics
	}
)

func NewMetricsLogger(ctx context.Context, lg logger.Logger, metricsPath string, metricsPort string) MetricsLogger {
	if lg == nil {
		lg = logger.NewDefaultLogger()
	}

	return &metricsLoggerImpl{
		exec: logger.NewMessageExecutorImpl(ctx, lg, loggerWorkersNum, loggerBufferSize),
		m:    metrics.NewMetricsPublisher(metricsPath, metricsPort),
	}
}

func (i *metricsLoggerImpl) Debug(_ context.Context, message logger.Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(DebugLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Info(_ context.Context, message logger.Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(InfoLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Warn(_ context.Context, message logger.Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(WarnLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Error(_ context.Context, message logger.Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(ErrorLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Fatal(_ context.Context, message logger.Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(FatalLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) generateMetrics(message logger.Message) {
	i.m.IncOp(message.OperationName, message.ComponentName, message.EventStatus)
	i.m.Observe(message.OperationName, message.ComponentName, message.EventStatus, message.Latency)
}

func (i *metricsLoggerImpl) Close() error {
	i.m.Close()
	return i.exec.Close()
}
