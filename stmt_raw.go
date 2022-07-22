package orm

import (
	"context"
	"database/sql"

	"github.com/deweppro/go-errors"
)

//Ping database ping
func (s *Stmt) Ping() error {
	return s.CallContext("ping", context.Background(), func(ctx context.Context, db *sql.DB) error {
		return db.PingContext(ctx)
	})
}

//CallContext basic query execution
func (s *Stmt) CallContext(name string, ctx context.Context, callFunc func(context.Context, *sql.DB) error) error {
	pool, err := s.db.Pool(s.name)
	if err != nil {
		return err
	}

	s.plug.Metrics.ExecutionTime(name, func() { err = callFunc(ctx, pool) })

	return err
}

//TxContext the basic execution of a query in a transaction
func (s *Stmt) TxContext(name string, ctx context.Context, callFunc func(context.Context, *sql.Tx) error) error {
	return s.CallContext(name, ctx, func(ctx context.Context, db *sql.DB) error {
		dbx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		err = callFunc(ctx, dbx)
		if err != nil {
			return errors.Wrap(
				errors.WrapMessage(err, "execute tx"),
				errors.WrapMessage(dbx.Rollback(), "rollback tx"),
			)
		}

		return dbx.Commit()
	})
}
