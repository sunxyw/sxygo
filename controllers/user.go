/*
 * Package controllers
 * File: user.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:30
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-30 18:10:13
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package controllers

import (
	"lwgo/services"
	"lwgo/transformers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service services.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service services.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := transformers.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, transformers.Response{
		Code:    0,
		Message: "登出成功",
	})
}
