/*
 * Package transformers
 * File: user.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:14
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-30 18:2:47
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package transformers

import "lwgo/models"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

const (
	// CodeAuthorizationFailError 认证失败
	CodeAuthorizationFailError = 20001
	// CodeInvalidCredentialsError 用户名或密码错误
	CodeInvalidCredentialsError = 20002
	// CodeBadTokenError Token 不合法
	CodeBadTokenError = 20003
	// CodeTokenExpiredError Token 已过期
	CodeTokenExpiredError = 20004
	// CodeUnauthorizedError 未通过认证
	CodeUnauthorizedError = 20005
)

// BuildUser 序列化用户
func BuildUser(user models.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user models.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
