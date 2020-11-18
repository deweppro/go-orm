/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package sqlite

import (
	"database/sql"
	"sync"

	"github.com/deweppro/go-orm/schema"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var (
	_ schema.Connector       = (*pool)(nil)
	_ schema.ConfigInterface = (*Config)(nil)
)

type (
	Config struct {
		Pool []Item `yaml:"sqlite"`
	}

	Item struct {
		Name string `yaml:"name"`
		File string `yaml:"file"`
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

func (i Item) GetName() string               { return i.Name }
func (i Item) GetDSN() string                { return i.File }
func (i Item) Setup(_ schema.SetupInterface) {}

func New(conf schema.ConfigInterface) (schema.Connector, error) {
	c := &pool{
		conf: conf,
		db:   make(map[string]*sql.DB),
	}

	return c, c.Reconnect()
}

func (p *pool) Dialect() string {
	return schema.SQLiteDialect
}

func (p *pool) Reconnect() error {
	if err := p.Close(); err != nil {
		return err
	}

	p.l.Lock()
	defer p.l.Unlock()

	for _, item := range p.conf.List() {
		db, err := sql.Open("sqlite3", item.GetDSN())
		if err != nil {
			if er := p.Close(); er != nil {
				return errors.Wrap(err, er.Error())
			}
			return err
		}
		p.db[item.GetName()] = db
	}
	return nil
}

func (p *pool) Close() error {
	p.l.Lock()
	defer p.l.Unlock()

	if len(p.db) > 0 {
		for _, db := range p.db {
			if err := db.Close(); err != nil {
				return err
			}
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
