package server

import (
	"os"
	"lwgo/controllers"
	"lwgo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/controllers/v1")
	{
		v1.POST("ping", controllers.Ping)

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
