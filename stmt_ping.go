/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package orm

import (
	"context"
	"database/sql"
)

func (s *Stmt) Ping() error {
	return s.Call("ping", func(conn *sql.Conn, ctx context.Context) error {
		return conn.PingContext(ctx)
	})
}
