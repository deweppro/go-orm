/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package plugins

import (
	"fmt"
	"io"
	"time"
)

const (
	LogLevelSilent = iota
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
)

type (
	log struct {
		level int
		log   Logger
	}
	Logger interface {
		LogLevel() int
		Info(msg string, args ...interface{})
		Warn(msg string, args ...interface{})
		Error(msg string, args ...interface{})
		Debug(msg string, args ...interface{})
	}
)

func newLog(l Logger) *log {
	_log := &log{
		level: LogLevelSilent,
		log:   l,
	}
	_log.setup()
	return _log
}

func (l *log) setup() {
	if l.log != nil {
		l.level = l.log.LogLevel()
	}
}

func (l *log) Info(msg string, args ...interface{}) {
	if l.level >= LogLevelInfo && l.log != nil {
		l.log.Info(msg, args...)
	}
}

func (l *log) Warn(msg string, args ...interface{}) {
	if l.level >= LogLevelWarn && l.log != nil {
		l.log.Warn(msg, args...)
	}
}

func (l *log) Error(msg string, args ...interface{}) {
	if l.level >= LogLevelError && l.log != nil {
		l.log.Error(msg, args...)
	}
}

func (l *log) Debug(msg string, args ...interface{}) {
	if l.level >= LogLevelDebug && l.log != nil {
		l.log.Debug(msg, args...)
	}
}

var (
	_ Logger = (*simpleLog)(nil)

	StdOutLog = &simpleLog{Level: LogLevelDebug, Writer: StdOutWriter}
)

type simpleLog struct {
	Level  int
	Writer io.Writer
}

func (l *simpleLog) LogLevel() int {
	return l.Level
}

func (l *simpleLog) currentTime() string {
	return time.Now().Format(time.RFC3339)
}

func (l *simpleLog) Info(msg string, args ...interface{}) {
	_, _ = l.Writer.Write([]byte("[INFO] " + l.currentTime() + " - " + fmt.Sprintf(msg, args...))) //nolint:errcheck
}

func (l *simpleLog) Warn(msg string, args ...interface{}) {
	_, _ = l.Writer.Write([]byte("[WARN] " + l.currentTime() + " - " + fmt.Sprintf(msg, args...))) //nolint:errcheck
}

func (l *simpleLog) Error(msg string, args ...interface{}) {
	_, _ = l.Writer.Write([]byte("[ERRO] " + l.currentTime() + " - " + fmt.Sprintf(msg, args...))) //nolint:errcheck
}

func (l *simpleLog) Debug(msg string, args ...interface{}) {
	_, _ = l.Writer.Write([]byte("[DEBG] " + l.currentTime() + " - " + fmt.Sprintf(msg, args...))) //nolint:errcheck
}
