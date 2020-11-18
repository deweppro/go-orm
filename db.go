/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package orm

import (
	"github.com/deweppro/go-orm/plugins"
	"github.com/deweppro/go-orm/schema"
)

type (
	DB struct {
		conn schema.Connector
		plug *Plugins
	}
	Plugins struct {
		Logger  plugins.Logger
		Metrics plugins.Metrics
	}
)

func NewDB(c schema.Connector, p *Plugins) *DB {
	plug := &Plugins{
		Logger:  plugins.StdOutLog,
		Metrics: plugins.StdOutMetric,
	}
	if p != nil {
		if p.Metrics != nil {
			plug.Metrics = p.Metrics
		}
		if p.Logger != nil {
			plug.Logger = p.Logger
		}
	}
	return &DB{
		conn: c,
		plug: plug,
	}
}

func (d *DB) Pool(name string) StmtInterface {
	return newStmt(name, d.conn, d.plug)
}
