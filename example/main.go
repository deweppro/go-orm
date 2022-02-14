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

	db := orm.NewDB(conn, orm.Plugins{Logger: plugins.StdOutLog, Metrics: plugins.StdOutMetric})
	pool := db.Pool("")

	if err = pool.Ping(); err != nil {
		panic(err.Error())
	}

	err = pool.Call("demo_metric", func(ctx context.Context, db *sql.DB) error {
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}
