/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package plugins

import (
	"time"
)

type (
	metric struct {
		metrics MetricWriter
	}
	Metrics interface {
		ExecutionTime(name string, call func())
	}
	MetricWriter interface {
		Metric(name string, time time.Duration)
	}
)

var StdOutMetric = NewMetric(StdOutWriter)

func NewMetric(m MetricWriter) *metric {
	return &metric{metrics: m}
}

func (m *metric) ExecutionTime(name string, call func()) {
	if m.metrics == nil {
		call()
		return
	}

	t := time.Now()
	call()
	m.metrics.Metric(name, time.Now().Sub(t))
}
