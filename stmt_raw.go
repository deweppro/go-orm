package orm

import (
	"context"
	"database/sql"
	"errors"
)

//Call basic query execution
func (s *Stmt) Call(name string, fn func(*sql.Conn, context.Context) error) error {
	ctx, cncl := context.WithCancel(context.Background())
	defer cncl()

	pool, err := s.db.Pool(s.name)
	if err != nil {
		return err
	}

	conn, err := pool.Conn(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if er := conn.Close(); er != nil && !errors.Is(er, sql.ErrConnDone) {
			s.plug.Logger.Error("close connection: %s", er.Error())
		}
	}()

	s.plug.Metrics.ExecutionTime(name, func() { err = fn(conn, ctx) })

	return err
}

//Tx the basic execution of a query in a transaction
func (s *Stmt) Tx(name string, fn func(*sql.Tx, context.Context) error) error {
	return s.Call(name, func(conn *sql.Conn, ctx context.Context) error {
		tx, err := conn.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		defer func() {
			if err := tx.Rollback(); err != nil {
				s.plug.Logger.Error("tx rollback: %s", err.Error())
			}
		}()

		return fn(tx, ctx)
	})
}
