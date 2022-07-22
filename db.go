package orm

import (
	"github.com/deweppro/go-logger"
	"github.com/deweppro/go-orm/plugins"
	"github.com/deweppro/go-orm/schema"
)

type (
	//DB connection storage
	DB struct {
		conn schema.Connector
		plug Plugins
	}
	//Plugins storage
	Plugins struct {
		Logger  logger.Logger
		Metrics plugins.MetricGetter
	}
)

//NewDB init database connections
func NewDB(c schema.Connector, plug Plugins) *DB {
	if plug.Logger == nil {
		plug.Logger = plugins.DevNullLog
	}
	if plug.Metrics == nil {
		plug.Metrics = plugins.DevNullMetric
	}
	return &DB{
		conn: c,
		plug: plug,
	}
}

//Pool getting pool connections by name
func (d *DB) Pool(name string) *Stmt {
	return newStmt(name, d.conn, d.plug)
}
