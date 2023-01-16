package driver

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLiteDriver 是 SQLite 的 starter-gorm 驱动
type SQLiteDriver struct {
	markup.Component `class:"starter-gorm-driver-registry"`
}

func (inst *SQLiteDriver) _Impl() (datasource.Driver, datasource.DriverRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *SQLiteDriver) GetRegistration() *datasource.DriverRegistration {
	return &datasource.DriverRegistration{
		Driver: inst,
	}
}

// Accept 用于确定是否支持给定的配置
func (inst *SQLiteDriver) Accept(cfg *datasource.Configuration) bool {
	name := cfg.Driver
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name == "sqlite"
}

// Open 打开数据源
func (inst *SQLiteDriver) Open(cfg *datasource.Configuration) (datasource.Source, error) {

	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if cfg == nil {
		return nil, errors.New("config==nil")
	}

	path := cfg.Database

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &sqliteDataSource{db: db}, nil
}

func (inst *SQLiteDriver) prepareForDefaultPort(cfg *datasource.Configuration) {
	// const defport = 1433
	// port := cfg.Port
	// if port < 1 {
	// 	port = defport
	// }
	// cfg.Port = port
}

////////////////////////////////////////////////////////////////////////////////

type sqliteDataSource struct {
	db *gorm.DB
}

func (inst *sqliteDataSource) _Impl() datasource.Source {
	return inst
}

func (inst *sqliteDataSource) DB() (*gorm.DB, error) {
	return inst.db, nil
}

func (inst *sqliteDataSource) Close() error {
	return nil
}
