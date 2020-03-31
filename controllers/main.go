/*
 * Package controllers
 * File: main.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:14
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-30 21:1:45
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package controllers

import (
	"encoding/json"
	"fmt"
	"lwgo/config"
	"lwgo/helpers"
	"lwgo/models"
	"lwgo/transformers"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Status 获取服务状态
func Status(c *gin.Context) {
	// 检查接口状态（自检）
	apiStatus := "online"

	// 检查数据库状态
	dbErr := models.DB.DB().Ping()
	dbStatus := "online"
	if dbErr != nil {
		dbStatus = dbErr.Error()
	}

	// 检查服务器状态
	mcStatus := "online"
	var rcon helpers.RCON
	xrcon, err := rcon.Connect("127.0.0.1", 25575, "itpassword")
	if err == nil {
		_, err := xrcon.Send("help")
		if err != nil {
			mcStatus = "offline"
		}
	} else {
		mcStatus = "offline"
	}

	c.JSON(200, transformers.Response{
		Code:    0,
		Message: "Checked",
		Data: struct {
			API       string `json:"api"`
			Database  string `json:"database"`
			Minecraft string `json:"minecraft"`
		}{apiStatus, dbStatus, mcStatus},
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *models.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*models.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) transformers.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return transformers.InvalidArgumentErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return transformers.InvalidArgumentErr("JSON类型不匹配", err)
	}

	return transformers.InvalidArgumentErr("参数错误", err)
}
