package db

import (
	"log"

	// "github.com/Vingurzhou/wechat-robot/query"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	GormDB *gorm.DB
}

func NewDatabase() *Database {
	newLogger := logger.New(log.Default(), logger.Config{})

	gormDB, err := gorm.Open(sqlite.Open("wechat_robot_db.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		GormDB: gormDB,
	}
}
