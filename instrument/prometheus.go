package instrument

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// IPrometheus ...
type IPrometheus interface {
	NewCounter(module string, name string, help string) prometheus.Counter
}

// Prometheus ...
type Prometheus struct{}

func (p *Prometheus) generateName(module string, name string) string {
	return "m9_" + module + "_" + name
}

// NewCounter ...
func (p *Prometheus) NewCounter(module string, name string, help string) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Name: p.generateName(module, name),
		Help: help,
	})
}

// NewPrometheus ...
func NewPrometheus() IPrometheus {
	return &Prometheus{}
}
