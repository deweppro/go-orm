# go-orm

## Deprecated. Use https://github.com/deweppro/go-sdk

# Examples

## Init connection

```go
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
	defer conn.Close()
	if err := conn.Reconnect(); err != nil {
		panic(err.Error())
	}

	db := orm.NewDB(conn, orm.Plugins{Logger: plugins.StdOutLog, Metrics: plugins.StdOutMetric})
	pool := db.Pool("main_db")

	if err := pool.Ping(); err != nil {
		panic(err.Error())
	}

	// use pool[main_db] here
	err := pool.CallContext("query name", context.Background(), func(ctx context.Context, db *sql.DB) error {...}
	err := pool.TxContext("query name", context.Background(), func(context.Context, *sql.Tx) error) error {...}
}
```

## Basic query

```go
package main

import (
	"context"
	"database/sql"

	"github.com/deweppro/go-orm"
	"github.com/deweppro/go-orm/plugins"
	"github.com/deweppro/go-orm/schema/mysql"
)

func main() {
	...

	var userName string
	err := pool.CallContext("user_name", context.Background(), func(ctx context.Context, db *sql.DB) error {
		rows, err := db.QueryContext(ctx, "select `name` from `users` where `id`=?", 10)
		if err != nil {
			return err
		}
		defer rows.Close()

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

```

## Lite query

```go
package main

import (
	"context"
	"fmt"

	"github.com/deweppro/go-orm"
	"github.com/deweppro/go-orm/plugins"
	"github.com/deweppro/go-orm/schema/mysql"
)

func main() {
	...

	var userName string
	err := pool.QueryContext("user_name", context.Background(), func(q orm.Querier) {
		q.SQL("select `name` from `users` limit 1")
		q.Bind(func(bind orm.Scanner) error {
			return bind.Scan(&userName)
		})
	})

	err = pool.ExecContext("user_name", context.Background(), func(e orm.Executor) {
		e.SQL("insert into `users` (`id`, `name`) values (?, ?);")
		e.Params(3, "cccc")

		e.Bind(func(result orm.Result) error {
			fmt.Printf("RowsAffected=%d LastInsertId=%d", result.RowsAffected, result.LastInsertId)
			return nil
		})
	})

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
}
```
