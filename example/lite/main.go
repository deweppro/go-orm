package main

import (
	"context"
	"fmt"

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

	if err := conn.Reconnect(); err != nil {
		panic(err.Error())
	}

	db := orm.NewDB(conn, orm.Plugins{Logger: plugins.StdOutLog, Metrics: plugins.StdOutMetric})
	pool := db.Pool("main_db")

	if err := pool.Ping(); err != nil {
		panic(err.Error())
	}

	var userName string
	err := pool.QueryContext("user_name", context.Background(), func(q orm.Querier) {
		q.SQL("select `name` from `users` limit 1")
		q.Bind(func(bind orm.Scanner) error {
			return bind.Scan(&userName)
		})
	})
	if err != nil {
		panic(err.Error())
	}

	err = pool.ExecContext("user_name", context.Background(), func(e orm.Executor) {
		e.SQL("insert into `users` (`id`, `name`) values (?, ?);")
		e.Params(3, "cccc")

		e.Bind(func(result orm.Result) error {
			fmt.Printf("RowsAffected=%d LastInsertId=%d", result.RowsAffected, result.LastInsertId)
			return nil
		})
	})
	if err != nil {
		panic(err.Error())
	}

	err = pool.TransactionContext("", context.Background(), func(v orm.Tx) {
		v.Exec(func(e orm.Executor) {
			e.SQL("insert into `users` (`id`, `name`) values (?, ?);")
			e.Params(3, "cccc")

			e.Bind(func(result orm.Result) error {
				fmt.Printf("RowsAffected=%d LastInsertId=%d", result.RowsAffected, result.LastInsertId)
				return nil
			})
		})
		v.Query(func(q orm.Querier) {
			q.SQL("select `name` from `users` limit 1")
			q.Bind(func(bind orm.Scanner) error {
				return bind.Scan(&userName)
			})
		})
	})
	if err != nil {
		panic(err.Error())
	}
}
