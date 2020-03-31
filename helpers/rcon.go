/*
 * Package helpers
 * File: rcon.go
 * Project: SXYGo
 * File Created: 2020-03-30 18:15:59
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-30 21:3:48
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package helpers

import (
	"lwgo/utils"

	gorcon "github.com/seeruk/minecraft-rcon/rcon"
)

// RCON RCON 连接助手
type RCON struct {
	Connection *gorcon.Client
}

// Connect 建立 RCON 连接
func (rcon *RCON) Connect(host string, port int, password string) (_ *RCON, err error) {
	client, err := gorcon.NewClient(host, port, password)
	if err != nil {
		utils.Log().Errorf("无法连接至 RCON %v", err)
		return nil, err
	}
	rcon.Connection = client

	return rcon, nil
}

// Close 关闭 RCON 连接
func (rcon *RCON) Close() {
	//
}

// Send 发送指令
func (rcon *RCON) Send(command string) (string, error) {
	resp, err := rcon.Connection.SendCommand(command)
	if err != nil {
		utils.Log().Errorf("RCON 指令发送失败 %v", err)
		return "", err
	}
	return resp, nil
}
