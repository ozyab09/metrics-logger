package metrics

import (
	"time"

	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics - metrics counter interface
type Metrics interface {
	// IncOp - counting the number of operations
	IncOp(opName string, componentName string, status logger.Status)
	// Observe - calculation of operation execution time
	Observe(opName string, componentName string, status logger.Status, d time.Duration)
	// Close - close metrics counter
	Close()
}

type metricsImpl struct {
	opCount   *prometheus.CounterVec
	opLatency *prometheus.HistogramVec
}

// NewMetrics - return new metrics counter implementation
func NewMetrics() Metrics {
	opCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_op_count",
			Help: "Operation counter metric",
		},
		[]string{"op_name", "component_name", "status"},
	)
	opLatency := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "service_op_latency",
			Help: "Latency of operations",
		},
		[]string{"op_name", "component_name", "status"},
	)
	prometheus.MustRegister(opCount, opLatency)

	return &metricsImpl{opCount: opCount, opLatency: opLatency}
}

// IncOp - counting the number of operations
func (i *metricsImpl) IncOp(opName string, componentName string, status logger.Status) {
	i.opCount.WithLabelValues(
		opName, componentName, string(status),
	).Inc()
}

// Observe - calculation of operation execution time
func (i *metricsImpl) Observe(opName string, componentName string, status logger.Status, d time.Duration) {
	i.opLatency.
		WithLabelValues(opName, componentName, string(status)).
		Observe(d.Seconds())
}

// Close - close metrics counter
func (i *metricsImpl) Close() {
	prometheus.Unregister(i.opCount)
	prometheus.Unregister(i.opLatency)
}
