package go_metrics_logger

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
	"github.com/ozyab09/metrics-logger/go_metrics_logger/metrics"
)

const (
	loggerWorkersNum = 512
	loggerBufferSize = 65536
)

const (
	// DebugLevel - debug logging level
	DebugLevel = "DEBUG"
	// InfoLevel - info logging level
	InfoLevel = "INFO"
	// WarnLevel - warn logging level
	WarnLevel = "WARN"
	// ErrorLevel - error logging level
	ErrorLevel = "ERROR"
	// FatalLevel - fatal logging level
	FatalLevel = "FATAL"
)

type (
	// MetricsLogger - logging messages and generate metrics interface
	MetricsLogger interface {
		// Debug - log message with level debug and generate metrics
		Debug(ctx context.Context, message logger.I_Message)
		// Info - log message with level info and generate metrics
		Info(ctx context.Context, message logger.I_Message)
		// Warn - log message with level warn and generate metrics
		Warn(ctx context.Context, message logger.I_Message)
		// Error - log message with level error and generate metrics
		Error(ctx context.Context, message logger.I_Message)
		// Fatal - log message with level fatal and generate metrics
		Fatal(ctx context.Context, message logger.I_Message)
		// Close - close metrics logger
		Close() error
	}

	metricsLoggerImpl struct {
		exec logger.MessageExecutor
		m    metrics.Metrics
	}
)

// NewMetricsLogger - return new metrics-logger implementation
func NewMetricsLogger(ctx context.Context, lg logger.Logger, metricsPath string, metricsPort string) MetricsLogger {
	if lg == nil {
		lg = logger.NewDefaultLogger()
	}

	return &metricsLoggerImpl{
		exec: logger.NewMessageExecutorImpl(ctx, lg, loggerWorkersNum, loggerBufferSize),
		m:    metrics.NewMetricsPublisher(metricsPath, metricsPort),
	}
}

// Debug - log message with level debug and generate metrics
func (i *metricsLoggerImpl) Debug(_ context.Context, message logger.I_Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(DebugLevel, string(data))
	}
	i.generateMetrics(message)
}

// Info - log message with level info and generate metrics
func (i *metricsLoggerImpl) Info(_ context.Context, message logger.I_Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(InfoLevel, string(data))
	}
	i.generateMetrics(message)
}

// Warn - log message with level warn and generate metrics
func (i *metricsLoggerImpl) Warn(_ context.Context, message logger.I_Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(WarnLevel, string(data))
	}
	i.generateMetrics(message)
}

// Error - log message with level error and generate metrics
func (i *metricsLoggerImpl) Error(_ context.Context, message logger.I_Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(ErrorLevel, string(data))
	}
	i.generateMetrics(message)
}

// Fatal - log message with level fatal and generate metrics
func (i *metricsLoggerImpl) Fatal(_ context.Context, message logger.I_Message) {
	data, err := json.Marshal(message)
	if err == nil {
		i.exec.ExecuteMessage(FatalLevel, string(data))
	}
	i.generateMetrics(message)
}

func (i *metricsLoggerImpl) generateMetrics(message logger.I_Message) {
	i.m.IncOp(message.GetOperationName(), message.GetComponentName(), message.GetEventStatus())
	if message.GetLatency() != time.Duration(0) {
		i.m.Observe(message.GetOperationName(), message.GetComponentName(), message.GetEventStatus(), message.GetLatency())
	}
	i.m.Trigger(message.GetOperationName(), message.GetComponentName(), message.GetEventStatus())
}

// Close - close metrics logger
func (i *metricsLoggerImpl) Close() error {
	i.m.Close()
	return i.exec.Close()
}
