package go_metrics_logger

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricsPublisher struct {
	m          Metrics
	handlePath string
	handlePort string
}

func newMetricsPublisher(path string, port string) *metricsPublisher {
	p := &metricsPublisher{
		m:          NewMetrics(),
		handlePath: path,
		handlePort: port,
	}
	go p.init()

	return p
}

func (p *metricsPublisher) init() {
	http.Handle(p.handlePath, promhttp.Handler())
	http.ListenAndServe(p.handlePort, nil)
}

func (p *metricsPublisher) IncOp(opName string, componentName string, status Status) {
	p.m.IncOp(opName, componentName, status)
}

func (p *metricsPublisher) Observe(opName string, componentName string, status Status, d time.Duration) {
	p.m.Observe(opName, componentName, status, d)
}

func (p *metricsPublisher) Close() {
	p.m.Close()
}
