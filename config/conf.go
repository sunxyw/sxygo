package config

import (
	"os"
	"lwgo/cache"
	"lwgo/models"
	"lwgo/utils"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	utils.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		utils.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	models.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
