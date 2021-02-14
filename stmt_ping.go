package orm

import (
	"context"
	"database/sql"
)

//Ping database ping
func (s *Stmt) Ping() error {
	return s.Call("ping", func(conn *sql.Conn, ctx context.Context) error {
		return conn.PingContext(ctx)
	})
}
