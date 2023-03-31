package driver

import (
	"database/sql"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	posgtresql "github.com/gowok/postgresql"
)

type DB struct {
	*sql.DB
}

func initDatabase() {
	conf := ioc.MustGet(gowok.Config{})
	var pgdb *posgtresql.PostgreSQL

	for _, dbConf := range conf.Databases {
		var err error
		pgdb, err = posgtresql.New(dbConf)

		if err != nil {
			panic(err)
		}

		if err := pgdb.Ping(); err != nil {
			panic(err)
		}

		if pgdb != nil {
			db := DB{pgdb.DB}
			ioc.Set(func() DB { return db })
		}
	}

}
