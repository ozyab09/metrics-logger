package go_metrics_logger

import (
	"context"
	"encoding/json"
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
		Debug(ctx context.Context, message Message)
		Info(ctx context.Context, message Message)
		Warn(ctx context.Context, message Message)
		Error(ctx context.Context, message Message)
		Fatal(ctx context.Context, message Message)
		Close()
	}

	metricsLoggerImpl struct {
		exec messageExecutor
		m    Metrics
	}
)

func NewMetricsLogger(ctx context.Context, logger Logger, metricsPath string, metricsPort string) MetricsLogger {
	if logger == nil {
		logger = NewDefaultLogger()
	}

	return &metricsLoggerImpl{
		exec: newMessageExecutorImpl(ctx, logger, loggerWorkersNum, loggerBufferSize),
		m:    newMetricsPublisher(metricsPath, metricsPort),
	}
}

func (i *metricsLoggerImpl) Debug(_ context.Context, message Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(DebugLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Info(_ context.Context, message Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(InfoLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Warn(_ context.Context, message Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(WarnLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Error(_ context.Context, message Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(ErrorLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) Fatal(_ context.Context, message Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(FatalLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) generateMetrics(message Message) {
	i.m.IncOp(message.OperationName, message.ComponentName, message.EventStatus)
	i.m.Observe(message.OperationName, message.ComponentName, message.EventStatus, message.Latency)
}

func (i *metricsLoggerImpl) Close() {
	i.m.Close()
}
