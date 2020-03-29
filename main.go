package main

import (
	"lwgo/config"
	"lwgo/server"
)

func main() {
	// 从配置文件读取配置
	config.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
