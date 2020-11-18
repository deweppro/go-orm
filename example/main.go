/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"context"
	"database/sql"

	"github.com/deweppro/go-orm"
	"github.com/deweppro/go-orm/plugins"
	"github.com/deweppro/go-orm/schema/mysql"
)

func main() {

	conn, err := mysql.New(&mysql.Config{Pool: []mysql.Item{}})
	if err != nil {
		panic(err.Error())
	}

	db := orm.NewDB(conn, &orm.Plugins{Logger: plugins.StdOutLog, Metrics: plugins.StdOutMetric})
	pool := db.Pool("")

	if err = pool.Ping(); err != nil {
		panic(err.Error())
	}

	err = pool.Call("demo_metric", func(conn *sql.Conn, ctx context.Context) error {
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}
