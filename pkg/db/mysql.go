package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var mysqldb *gorm.DB

func MysqlDB() *gorm.DB {
	if mysqldb == nil {
		panic("db 模块没有初始化")
	}

	return mysqldb
}

type MysqlOpts struct {
	Endpoint string
	Username string
	Password string
	Database string
}

func NewMysqlClient(mysqlOpts *MysqlOpts) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&autocommit=true",
		mysqlOpts.Username,
		mysqlOpts.Password,
		mysqlOpts.Endpoint,
		mysqlOpts.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("Could not connect to the database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	mysqldb = db
	return db, nil
}
