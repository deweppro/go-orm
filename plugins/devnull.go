package plugins

import (
	"io"

	"github.com/deweppro/go-logger"
)

var (
	DevNullLog    logger.Logger = new(devNullLogger)
	DevNullMetric MetricGetter  = new(devNullMetric)
)

type devNullMetric struct{}

func (devNullMetric) ExecutionTime(_ string, call func()) { call() }

type devNullLogger struct{}

func (devNullLogger) SetOutput(io.Writer)                           {}
func (devNullLogger) Fatalf(string, ...interface{})                 {}
func (devNullLogger) Errorf(string, ...interface{})                 {}
func (devNullLogger) Warnf(string, ...interface{})                  {}
func (devNullLogger) Infof(string, ...interface{})                  {}
func (devNullLogger) Debugf(string, ...interface{})                 {}
func (devNullLogger) SetLevel(v uint32)                             {}
func (devNullLogger) Close()                                        {}
func (devNullLogger) GetLevel() uint32                              { return 0 }
func (v devNullLogger) WithFields(_ logger.Fields) logger.LogWriter { return v }
