package orm

import (
	"context"
	"database/sql"
)

//Ping database ping
func (s *Stmt) Ping() error {
	return s.Call("ping", func(ctx context.Context, db *sql.DB) error {
		return db.PingContext(ctx)
	})
}
