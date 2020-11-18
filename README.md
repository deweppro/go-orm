# go-orm

[![Coverage Status](https://coveralls.io/repos/github/deweppro/go-orm/badge.svg?branch=master)](https://coveralls.io/github/deweppro/go-orm?branch=master)
[![Release](https://img.shields.io/github/release/deweppro/go-orm.svg?style=flat-square)](https://github.com/deweppro/go-orm/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/go-orm)](https://goreportcard.com/report/github.com/deweppro/go-orm)
[![Build Status](https://travis-ci.com/deweppro/go-orm.svg?branch=master)](https://travis-ci.com/deweppro/go-orm)

## Introduction

The library provides a nice and simple ActiveRecord implementation for working with your database. Each database table has a corresponding "Model" that is used to interact with this table. Models allow you to query data in tables, as well as insert new records in the table.

# DEMO

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

```