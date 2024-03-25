package metrics

import (
	"github.com/ozyab09/metrics-logger/go_metrics_logger/logger"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsPublisher struct {
	m          Metrics
	handlePath string
	handlePort string
}

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

func (p *MetricsPublisher) IncOp(opName string, componentName string, status logger.Status) {
	p.m.IncOp(opName, componentName, status)
}

func (p *MetricsPublisher) Observe(opName string, componentName string, status logger.Status, d time.Duration) {
	p.m.Observe(opName, componentName, status, d)
}

func (p *MetricsPublisher) Close() {
	p.m.Close()
}
