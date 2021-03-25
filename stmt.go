package orm

import (
	"context"
	"database/sql"
)

type (
	//Stmt statement model
	Stmt struct {
		name   string
		db     dbPool
		plug   Plugins
		models map[string]ormModel
	}
	dbPool interface {
		Dialect() string
		Pool(string) (*sql.DB, error)
	}
	//StmtInterface statement interface
	StmtInterface interface {
		Call(string, func(*sql.Conn, context.Context) error) error
		Tx(string, func(*sql.Tx, context.Context) error) error
		Ping() error
	}
)

//newStmt init new statement
func newStmt(name string, db dbPool, p Plugins) *Stmt {
	return &Stmt{
		name: name,
		db:   db,
		plug: p,
	}
}
