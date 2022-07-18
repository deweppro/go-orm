package main

import (
	"context"
	"database/sql"

	"github.com/deweppro/go-orm"
	"github.com/deweppro/go-orm/plugins"
	"github.com/deweppro/go-orm/schema/mysql"
)

func main() {
	conn := mysql.New(&mysql.Config{
		Pool: []mysql.Item{
			{
				Name:     "main_db",
				Host:     "127.0.0.1",
				Port:     3306,
				Schema:   "test_table",
				User:     "demo",
				Password: "1234",
			},
		}})
	defer conn.Close() //nolint: errcheck
	if err := conn.Reconnect(); err != nil {
		panic(err.Error())
	}

	db := orm.NewDB(conn, orm.Plugins{Logger: plugins.StdOutLog, Metrics: plugins.StdOutMetric})
	pool := db.Pool("main_db")

	if err := pool.Ping(); err != nil {
		panic(err.Error())
	}

	var userName string
	err := pool.CallContext("user_name", context.Background(), func(ctx context.Context, db *sql.DB) error {
		rows, err := db.QueryContext(ctx, "select `name` from `users` where `id`=?", 10)
		if err != nil {
			return err
		}
		defer rows.Close() //nolint: errcheck

		for rows.Next() {
			if err = rows.Scan(&userName); err != nil {
				return err
			}
		}
		if err = rows.Close(); err != nil {
			return err
		}
		if err = rows.Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}
