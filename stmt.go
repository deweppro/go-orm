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

type (
	Stmt struct {
		name   string
		db     dbPool
		plug   *Plugins
		models map[string]ormModel
	}
	dbPool interface {
		Dialect() string
		Pool(string) (*sql.DB, error)
	}
	StmtInterface interface {
		Call(string, func(*sql.Conn, context.Context) error) error
		Tx(string, func(*sql.Tx, context.Context) error) error
		Ping() error
	}
)

func newStmt(name string, db dbPool, p *Plugins) *Stmt {
	return &Stmt{
		name: name,
		db:   db,
		plug: p,
	}
}
