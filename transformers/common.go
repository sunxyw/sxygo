/*
 * Package transformers
 * File: common.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:33:19
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-30 18:8:47
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package transformers

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 错误代码，约定：
// - 以 1 开头即为系统级错误，2 开头为服务级错误
// - 第 2-3 位数字为服务模块代码
// - 第 4-5 位数字为具体错误代码
// 以下定义了常见的系统级错误代码
// 错误代码
const (
	// CodeSystemError 系统错误
	CodeSystemError = 10001
	// CodeServiceNotAvailableError 服务不可用
	CodeServiceNotAvailableError = 10002
	// CodeRemoteServiceError 远程服务错误
	CodeRemoteServiceError = 10003
	// CodeIPLimitDeniedError IP 限制不能请求
	CodeIPLimitDeniedError = 10004
	// CodeAccessDeniedError 权限不足
	CodeAccessDeniedError = 10005
	// CodeInvalidArgumentError 传入参数错误
	CodeInvalidArgumentError = 10006
	// CodeSystemBusyError 系统繁忙
	CodeSystemBusyError = 10007
	// CodeResponseTimeoutError 响应超时
	CodeResponseTimeoutError = 10008
	// CodeIllegalRequestError 非法请求
	CodeIllegalRequestError = 10009
	// CodeMissingArgumentError 缺少参数
	CodeMissingArgumentError = 10010
	// CodeEndpointNotExistsError 接口不存在
	CodeEndpointNotExistsError = 10011
	// CodeBadRequestMethodError 请求方法不支持
	CodeBadRequestMethodError = 10012
	// CodeTooManyRequestError 请求次数过多
	CodeTooManyRequestError = 10013
	// CodeMaintenanceModeError 服务维护中
	CodeMaintenanceModeError = 10014
	// CodeDatabaseOperationFailError 数据库操作失败
	CodeDatabaseOperationFailError = 10015
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code:    CodeUnauthorizedError,
		Message: "未通过认证",
	}
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code:    errCode,
		Message: msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDatabaseOperationFailError, msg, err)
}

func InvalidArgumentErr(msg string, err error) Response {
	return Err(CodeInvalidArgumentError, msg, err)
}
