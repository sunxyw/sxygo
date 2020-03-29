package services

import (
	"lwgo/models"
	"lwgo/transformers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginservices 管理用户登录的服务
type UserLoginservices struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (services *UserLoginservices) setSession(c *gin.Context, user models.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (services *UserLoginservices) Login(c *gin.Context) transformers.Response {
	var user models.User

	if err := models.DB.Where("user_name = ?", services.UserName).First(&user).Error; err != nil {
		return transformers.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(services.Password) == false {
		return transformers.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	services.setSession(c, user)

	return transformers.BuildUserResponse(user)
}
