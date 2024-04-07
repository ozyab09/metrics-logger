package metrics

import (
	"net/http"
	"time"

	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsPublisher - publish metrics in public port
type MetricsPublisher struct {
	m          Metrics
	handlePath string
	handlePort string
}

// NewMetricsPublisher - return new metrics publisher implementation
func NewMetricsPublisher(path string, port string) *MetricsPublisher {
	p := &MetricsPublisher{
		m:          NewMetrics(),
		handlePath: path,
		handlePort: port,
	}
	go p.init()

	return p
}

func (p *MetricsPublisher) init() {
	http.Handle(p.handlePath, promhttp.Handler())
	http.ListenAndServe(p.handlePort, nil)
}

// IncOp - counting the number of operations
func (p *MetricsPublisher) IncOp(opName string, componentName string, status logger.Status) {
	p.m.IncOp(opName, componentName, status)
}

// Observe - calculation of operation execution time
func (p *MetricsPublisher) Observe(opName string, componentName string, status logger.Status, d time.Duration) {
	p.m.Observe(opName, componentName, status, d)
}

// Close - close metrics counter
func (p *MetricsPublisher) Close() {
	p.m.Close()
}
