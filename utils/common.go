/*
 * Package utils
 * File: common.go
 * Project: SXYGo
 * File Created: 2020-03-29 17:11:50
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-29 17:33:22
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package utils

import (
	"math/rand"
	"time"
)

// RandStringRunes 生成随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
