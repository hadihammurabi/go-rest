package driver

import (
	"github.com/gowok/gowok/config"
	"github.com/gowok/ioc"
	posgtresql "github.com/gowok/postgresql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type DB struct {
	*bun.DB
}

func Database(conf []config.Database) {
	var pgdb *posgtresql.PostgreSQL

	for _, dbConf := range conf {
		var err error
		pgdb, err = posgtresql.New(dbConf)

		if err != nil {
			panic(err)
		}

		if err := pgdb.Ping(); err != nil {
			panic(err)
		}

		if pgdb != nil {
			db := DB{bun.NewDB(pgdb.DB, pgdialect.New())}
			ioc.Set(func() DB { return db })
		}
	}

}
