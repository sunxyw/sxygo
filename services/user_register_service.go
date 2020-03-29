package services

import (
	"lwgo/models"
	"lwgo/transformers"
)

// UserRegisterservices 管理用户注册服务
type UserRegisterservices struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Passwordconfigirm string `form:"password_configirm" json:"password_configirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (services *UserRegisterservices) valid() *transformers.Response {
	if services.Passwordconfigirm != services.Password {
		return &transformers.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	models.DB.models(&models.User{}).Where("nickname = ?", services.Nickname).Count(&count)
	if count > 0 {
		return &transformers.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	models.DB.models(&models.User{}).Where("user_name = ?", services.UserName).Count(&count)
	if count > 0 {
		return &transformers.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (services *UserRegisterservices) Register() transformers.Response {
	user := models.User{
		Nickname: services.Nickname,
		UserName: services.UserName,
		Status:   models.Active,
	}

	// 表单验证
	if err := services.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(services.Password); err != nil {
		return transformers.Err(
			transformers.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := models.DB.Create(&user).Error; err != nil {
		return transformers.ParamErr("注册失败", err)
	}

	return transformers.BuildUserResponse(user)
}
