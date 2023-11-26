package driver

import (
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/exception"
	"gorm.io/gorm"
)

var sql gowok.SQL

func GetSQL(name ...string) *gorm.DB {
	if sql != nil {
		return sql.Get(name...).OrPanic(exception.ErrNoDatabaseFound)
	}

	sql = gowok.Must(gowok.NewSQL(GetConfig().Databases))
	return sql.Get(name...).OrPanic(exception.ErrNoDatabaseFound)
}
