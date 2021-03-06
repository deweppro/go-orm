package plugins

import (
	"time"
)

type (
	metric struct {
		metrics MetricWriter
	}
	//MetricGetter interface
	MetricGetter interface {
		ExecutionTime(name string, call func())
	}
	//MetricWriter interface
	MetricWriter interface {
		Metric(name string, time time.Duration)
	}
)

//StdOutMetric simple stdout metrig writer
var StdOutMetric = NewMetric(StdOutWriter)

//NewMetric init new metric
func NewMetric(m MetricWriter) MetricGetter {
	return &metric{metrics: m}
}

//ExecutionTime calculating the execution time
func (m *metric) ExecutionTime(name string, call func()) {
	if m.metrics == nil {
		call()
		return
	}

	t := time.Now()
	call()
	m.metrics.Metric(name, time.Since(t))
}
