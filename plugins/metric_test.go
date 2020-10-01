/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package plugins

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMetric(t *testing.T) {
	w := &bytes.Buffer{}
	tl := &stdout{Writer: w}

	demo1 := NewMetric(nil)
	demo1.ExecutionTime("hello1", func() {
		return
	})

	demo2 := NewMetric(tl)
	demo2.ExecutionTime("hello2", func() {
		return
	})

	result := w.String()
	require.NotContains(t, result, "hello1")
	require.Contains(t, result, "hello2")
}
