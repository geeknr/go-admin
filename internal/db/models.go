package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql driver
	"github.com/spf13/viper"
)

var (
	orm *gorm.DB
)

func Init() (err error) {
	if err = SetEngin(); err != nil {
		return err
	}
	return nil
}

func SetEngin() (err error) {
	orm, err = getEngine()
	if err != nil {
		return fmt.Errorf("connect to database: %v", err)
	}

	orm.DB().SetMaxIdleConns(viper.GetInt("mysql.pool.min"))
	orm.DB().SetMaxOpenConns(viper.GetInt("mysql.pool.max"))

	if gin.Mode() != gin.ReleaseMode {
		orm.LogMode(true)
	}
	return nil
}

func getEngine() (*gorm.DB, error) {
	var connStr = viper.GetString("mysql.dsn")
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return &gorm.DB{}, err
	}
	return db, err
}
