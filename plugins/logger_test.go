package plugins

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	w := &bytes.Buffer{}

	tl := newLog(&simpleLog{Level: LogLevelSilent, Writer: w})
	tl.Info("info=msg-%d", 0)
	tl.Warn("warn=msg-%d", 0)
	tl.Debug("debug=msg-%d", 0)
	result := w.String()
	require.NotContains(t, result, "info=msg-0")
	require.NotContains(t, result, "warn=msg-0")
	require.NotContains(t, result, "debug=msg-0")

	w.Reset()

	tl = newLog(&simpleLog{Level: LogLevelInfo, Writer: w})
	tl.Info("info=msg-%d", 1)
	tl.Warn("warn=msg-%d", 1)
	tl.Debug("debug=msg-%d", 1)
	result = w.String()
	require.Contains(t, result, "info=msg-1")
	require.Contains(t, result, "warn=msg-1")
	require.NotContains(t, result, "debug=msg-1")

	w.Reset()

	tl = newLog(&simpleLog{Level: LogLevelWarn, Writer: w})
	tl.Info("info=msg-%d", 2)
	tl.Warn("warn=msg-%d", 2)
	tl.Debug("debug=msg-%d", 2)
	result = w.String()
	require.NotContains(t, result, "info=msg-2")
	require.Contains(t, result, "warn=msg-2")
	require.NotContains(t, result, "debug=msg-2")

	w.Reset()

	tl = newLog(&simpleLog{Level: LogLevelError, Writer: w})
	tl.Info("info=msg-%d", 2)
	tl.Warn("warn=msg-%d", 2)
	tl.Error("erro=msg-%d", 2)
	tl.Debug("debug=msg-%d", 2)
	result = w.String()
	require.NotContains(t, result, "info=msg-2")
	require.NotContains(t, result, "warn=msg-2")
	require.Contains(t, result, "erro=msg-2")
	require.NotContains(t, result, "debug=msg-2")

	w.Reset()

	tl = newLog(&simpleLog{Level: LogLevelDebug, Writer: w})
	tl.Info("info=msg-%d", 3)
	tl.Warn("warn=msg-%d", 3)
	tl.Debug("debug=msg-%d", 3)
	result = w.String()
	require.Contains(t, result, "info=msg-3")
	require.Contains(t, result, "warn=msg-3")
	require.Contains(t, result, "debug=msg-3")
}
