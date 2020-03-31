/*
 * Package sxygo
 * File: main.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:33:50
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-31 18:15:58
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package main

import (
	"lwgo/config"
	"lwgo/server"
	"lwgo/utils"
)

func main() {
	// 从配置文件读取配置
	config.Init()
	utils.Log().WithField("component", "rest").Error("hh")
	utils.Log().Info("HEllo")

	// 装载路由
	r := server.NewRouter()
	r.Run("0.0.0.0:3000")
}
