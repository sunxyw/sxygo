/*
 * Package models
 * File: init.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:14
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-29 22:51:34
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package models

import (
	"lwgo/utils"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string, isDebug bool) {
	db, err := gorm.Open("mysql", connString)
	// 启用日志记录
	db.LogMode(isDebug)

	if err != nil {
		utils.Log().Fatalf("无法连接至数据库 %v", err)
	}

	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(50)
	// 打开
	db.DB().SetMaxOpenConns(100)
	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	// 运行数据库迁移
	migration()
}
