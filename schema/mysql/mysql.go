/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/deweppro/go-orm/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	defaultTimeout     = time.Second * 5
	defaultTimeoutConn = time.Second * 60
)

var (
	_ schema.Connector       = (*pool)(nil)
	_ schema.ConfigInterface = (*Config)(nil)
)

type (
	Config struct {
		Pool []Item `yaml:"mysql"`
	}

	Item struct {
		Name              string        `yaml:"name"`
		Host              string        `yaml:"host"`
		Port              int           `yaml:"port"`
		Schema            string        `yaml:"schema"`
		User              string        `yaml:"user"`
		Password          string        `yaml:"password"`
		MaxIdleConn       int           `yaml:"maxidleconn"`
		MaxOpenConn       int           `yaml:"maxopenconn"`
		MaxConnTTL        time.Duration `yaml:"maxconnttl"`
		InterpolateParams bool          `yaml:"interpolateparams"`
		Timezone          string        `yaml:"timezone"`
		TxIsolationLevel  string        `yaml:"txisolevel"`
		Charset           string        `yaml:"charset"`
		Timeout           time.Duration `yaml:"timeout"`
		ReadTimeout       time.Duration `yaml:"readtimeout"`
		WriteTimeout      time.Duration `yaml:"writetimeout"`
	}

	pool struct {
		conf schema.ConfigInterface
		db   map[string]*sql.DB
		l    sync.RWMutex
	}
)

func (c *Config) List() (list []schema.ItemInterface) {
	for _, item := range c.Pool {
		list = append(list, item)
	}
	return
}

func (i Item) GetName() string {
	return i.Name
}

func (i Item) Setup(s schema.SetupInterface) {
	s.SetMaxIdleConns(i.MaxIdleConn)
	s.SetMaxOpenConns(i.MaxOpenConn)
	s.SetConnMaxLifetime(i.MaxConnTTL)
}

func (i Item) GetDSN() string {
	if len(i.Charset) == 0 {
		i.Charset = "utf8mb4,utf8"
	}
	if i.Timeout == 0 {
		i.Timeout = defaultTimeoutConn
	}
	if i.ReadTimeout == 0 {
		i.ReadTimeout = defaultTimeout
	}
	if i.WriteTimeout == 0 {
		i.WriteTimeout = defaultTimeout
	}
	if len(i.TxIsolationLevel) == 0 {
		i.TxIsolationLevel = "READ-COMMITTED"
	}
	if len(i.Timezone) == 0 {
		i.Timezone = "UTC"
	}
	base := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		i.User, i.Password, i.Host, i.Port, i.Schema)
	params := fmt.Sprintf(
		"charset=%s&autocommit=false&transaction_isolation=%s"+
			"&timeout=%s&readTimeout=%s&writeTimeout=%s&loc=%s&interpolateParams=%t",
		i.Charset, i.TxIsolationLevel, i.Timeout, i.ReadTimeout,
		i.WriteTimeout, i.Timezone, i.InterpolateParams,
	)

	return base + "?" + params
}

func New(conf schema.ConfigInterface) (schema.Connector, error) {
	c := &pool{
		conf: conf,
		db:   make(map[string]*sql.DB),
	}

	return c, c.Reconnect()
}

func (p *pool) Dialect() string {
	return schema.MySQLDialect
}

func (p *pool) Reconnect() error {
	if err := p.Close(); err != nil {
		return err
	}

	p.l.Lock()
	defer p.l.Unlock()

	for _, item := range p.conf.List() {
		db, err := sql.Open("mysql", item.GetDSN())
		if err != nil {
			if er := p.Close(); er != nil {
				return errors.Wrap(err, er.Error())
			}
			return err
		}
		item.Setup(db)
		p.db[item.GetName()] = db
	}
	return nil
}

func (p *pool) Close() error {
	p.l.Lock()
	defer p.l.Unlock()

	if len(p.db) > 0 {
		for name, db := range p.db {
			if err := db.Close(); err != nil {
				return err
			}
			delete(p.db, name)
		}
	}
	return nil
}

func (p *pool) Pool(name string) (*sql.DB, error) {
	p.l.RLock()
	defer p.l.RUnlock()

	db, ok := p.db[name]
	if !ok {
		return nil, schema.ErrPoolNotFound
	}
	return db, db.Ping()
}
