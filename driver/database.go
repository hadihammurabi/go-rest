package driver

import (
	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type DB struct {
	*bun.DB
}

func Database(conf []config.Database) {
	var pgdb *database.PostgreSQL

	for _, dbConf := range conf {
		var err error
		pgdb, err = database.NewPostgresql(dbConf)

		if err != nil {
			panic(err)
		}

		if pgdb != nil {
			db := DB{bun.NewDB(pgdb.DB, pgdialect.New())}
			ioc.Set(func() DB { return db })
		}
	}

}