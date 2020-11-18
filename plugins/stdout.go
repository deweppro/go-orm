/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package plugins

import (
	"fmt"
	"io"
	"os"
	"time"
)

type stdout struct {
	Writer io.Writer
}

var StdOutWriter = &stdout{Writer: os.Stdout}

func (s *stdout) currentTime() string {
	return time.Now().Format(time.RFC3339)
}

func (s *stdout) Write(p []byte) (n int, err error) {
	return s.Writer.Write(p)
}

func (s *stdout) Metric(name string, t time.Duration) {
	_, _ = s.Write([]byte("[MTRC] " + s.currentTime() + " - " + fmt.Sprintf("%s: %s", name, t))) //nolint:errcheck
}
