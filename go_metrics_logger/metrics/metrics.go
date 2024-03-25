package metrics

import (
	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	IncOp(opName string, componentName string, status logger.Status)
	Observe(opName string, componentName string, status logger.Status, d time.Duration)
	Close()
}

type metricsImpl struct {
	opCount   *prometheus.CounterVec
	opLatency *prometheus.HistogramVec
}

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

func (i *metricsImpl) IncOp(opName string, componentName string, status logger.Status) {
	i.opCount.WithLabelValues(
		opName, componentName, string(status),
	).Inc()
}

func (i *metricsImpl) Observe(opName string, componentName string, status logger.Status, d time.Duration) {
	i.opLatency.
		WithLabelValues(opName, componentName, string(status)).
		Observe(d.Seconds())
}

func (i *metricsImpl) Close() {
	prometheus.Unregister(i.opCount)
	prometheus.Unregister(i.opLatency)
}
