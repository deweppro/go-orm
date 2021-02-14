package orm

import (
	"github.com/deweppro/go-orm/plugins"
	"github.com/deweppro/go-orm/schema"
)

type (
	//DB connection storage
	DB struct {
		conn schema.Connector
		plug *Plugins
	}
	//Plugins storage
	Plugins struct {
		Logger  plugins.Logger
		Metrics plugins.Metrics
	}
)

//NewDB init database connections
func NewDB(c schema.Connector, p *Plugins) *DB {
	plug := &Plugins{
		Logger:  plugins.StdOutLog,
		Metrics: plugins.StdOutMetric,
	}
	if p != nil {
		if p.Metrics != nil {
			plug.Metrics = p.Metrics
		}
		if p.Logger != nil {
			plug.Logger = p.Logger
		}
	}
	return &DB{
		conn: c,
		plug: plug,
	}
}

//Pool getting pool connections by name
func (d *DB) Pool(name string) StmtInterface {
	return newStmt(name, d.conn, d.plug)
}
