/*
 * Package config
 * File: conf.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:14
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-29 23:0:31
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package config

import (
	"fmt"
	"lwgo/cache"
	"lwgo/models"
	"lwgo/utils"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 判断是否开启调试模式
	isDebug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	// 设置日志级别
	logLevel := os.Getenv("LOG_LEVEL")
	if isDebug {
		logLevel = "debug"
	}
	utils.BuildLogger(logLevel)
	utils.Log().Debug("调试模式已开启")

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		utils.Log().Fatalf("本地化文件加载失败 %v", err)
	}

	// 连接数据库
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	models.Database(dsn, isDebug)
	cache.Redis()
}
