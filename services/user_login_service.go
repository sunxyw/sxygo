/*
 * Package services
 * File: user_login_service.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:30
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-30 18:10:25
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package services

import (
	"lwgo/models"
	"lwgo/transformers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginservices 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user models.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) transformers.Response {
	var user models.User

	if err := models.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return transformers.Err(transformers.CodeInvalidCredentialsError, "用户名或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return transformers.Err(transformers.CodeInvalidCredentialsError, "用户名或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	return transformers.BuildUserResponse(user)
}
