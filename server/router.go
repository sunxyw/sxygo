/*
 * Package server
 * File: router.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:33:32
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-31 18:22:18
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package server

import (
	"lwgo/controllers"
	"lwgo/middleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	// 设置 GIN 运行模式
	isDebug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if isDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// 中间件, 顺序不能改
	r.Use(middleware.RequestID())
	r.Use(middleware.LogRequest())
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	r.GET("status", controllers.Status)

	// 路由
	v1 := r.Group("/controllers/v1")
	{

		// 用户登录
		v1.POST("user/register", controllers.UserRegister)

		// 用户登录
		v1.POST("user/login", controllers.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", controllers.UserMe)
			auth.DELETE("user/logout", controllers.UserLogout)
		}
	}
	return r
}
