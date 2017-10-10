package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kwri/go-workflow/modules/setting"
)

var Engine *gorm.DB

func Initialize() error {
	engine, err := NewEngine()
	if err != nil {
		return err
	}
	Engine = engine
	return nil
}

func NewEngine() (*gorm.DB, error) {
	config := setting.Db
	engine, err := gorm.Open(config.Driver, config.GetDataSourceName())
	engine.LogMode(config.Debug)
	return engine, err
}
