package orm

import (
	"context"
	"database/sql"
)

//Call basic query execution
func (s *Stmt) Call(name string, callFunc func(context.Context, *sql.DB) error) error {
	ctx, cncl := context.WithCancel(context.Background())
	defer cncl()

	pool, err := s.db.Pool(s.name)
	if err != nil {
		return err
	}

	s.plug.Metrics.ExecutionTime(name, func() { err = callFunc(ctx, pool) })

	return err
}

//Tx the basic execution of a query in a transaction
func (s *Stmt) Tx(name string, callFunc func(context.Context, *sql.Tx) error) error {
	return s.Call(name, func(ctx context.Context, db *sql.DB) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		defer func() {
			if err := tx.Rollback(); err != nil {
				s.plug.Logger.Errorf("tx rollback: %s", err.Error())
			}
		}()

		return callFunc(ctx, tx)
	})
}
