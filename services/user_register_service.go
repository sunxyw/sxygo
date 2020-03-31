/*
 * Package services
 * File: user_register_service.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:30
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-30 18:10:40
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package services

import (
	"lwgo/models"
	"lwgo/transformers"
)

// UserRegisterservices 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *transformers.Response {
	if service.PasswordConfirm != service.Password {
		return &transformers.Response{
			Code:    transformers.CodeInvalidArgumentError,
			Message: "两次输入的密码不相同",
		}
	}

	count := 0
	models.DB.Model(&models.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &transformers.Response{
			Code:    transformers.CodeInvalidArgumentError,
			Message: "昵称被占用",
		}
	}

	count = 0
	models.DB.Model(&models.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &transformers.Response{
			Code:    transformers.CodeInvalidArgumentError,
			Message: "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() transformers.Response {
	user := models.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   models.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return transformers.Err(
			transformers.CodeSystemError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := models.DB.Create(&user).Error; err != nil {
		return transformers.DBErr("注册失败", err)
	}

	return transformers.BuildUserResponse(user)
}
