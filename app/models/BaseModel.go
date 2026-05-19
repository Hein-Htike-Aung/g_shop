/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/soft_delete"
	"yixiang.co/go-mall/pkg/global"

	//"gorm.io/plugin/soft_delete"
	"log"
	"os"
	"time"
	"yixiang.co/go-mall/pkg/casbin"
)

var db *gorm.DB

type BaseModel struct {
	Id         int64                 `gorm:"primary_key" json:"id"`
	UpdateTime time.Time             `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime time.Time             `json:"createTime" gorm:"autoCreateTime"`
	IsDel      soft_delete.DeletedAt `json:"isDel" gorm:"softDelete:flag"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	var connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.YSHOP_CONFIG.Database.User,
		global.YSHOP_CONFIG.Database.Password,
		global.YSHOP_CONFIG.Database.Host,
		global.YSHOP_CONFIG.Database.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer for log output
		logger.Config{
			SlowThreshold:             time.Second,  // slow SQL threshold
			LogLevel:                  logger.Error, // log level
			IgnoreRecordNotFoundError: true,         // ignore ErrRecordNotFound
			Colorful:                  true,         // enable colorized output
		},
	)

	db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Printf("[info] gorm %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("[info] gorm %s", err)
	}

	// SetMaxIdleConns: max idle connections
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns: max open connections
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime: max connection lifetime
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.YSHOP_DB = db

	casbin.InitCasbin(db)

}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
