package driver

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	driver_sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDriver struct {
	markup.Component
}

func (inst *SQLiteDriver) _Impl() datasource.Driver {
	return inst
}

func (inst *SQLiteDriver) Accept(cfg *datasource.Configuration) bool {
	name := cfg.Driver
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return (name == "sqlite")
}

func (inst *SQLiteDriver) Open(cfg *datasource.Configuration) (datasource.Source, error) {

	//	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	dsn := cfg.Database

	dialector := driver_sqlite.Open(dsn)
	if dialector == nil {
		return nil, errors.New("driver_sqlserver.Open() return nil")
	}

	gc := &gorm.Config{}

	builder := &datasource.GormDataSourceBuilder{}
	builder.Config1 = *cfg
	builder.Config2 = *gc
	builder.Dialector = dialector
	return builder.Open()
}
