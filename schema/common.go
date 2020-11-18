/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package schema

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrPoolNotFound = errors.New("pool not found")
)

const (
	MySQLDialect  = "mysql"
	SQLiteDialect = "sqlite"
)

type (
	ConfigInterface interface {
		List() []ItemInterface
	}
	ItemInterface interface {
		GetName() string
		GetDSN() string
		Setup(SetupInterface)
	}
	SetupInterface interface {
		SetMaxIdleConns(int)
		SetMaxOpenConns(int)
		SetConnMaxLifetime(time.Duration)
	}
	Connector interface {
		Dialect() string
		Pool(string) (*sql.DB, error)
		Reconnect() error
		Close() error
	}
)
