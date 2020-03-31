/*
 * Package cache
 * File: main.go
 * Project: SXYGo
 * File Created: 2020-03-29 15:34:02
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-29 22:51:04
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package cache

import (
	"lwgo/utils"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		utils.Log().Fatalf("连接Redis不成功 %v", err)
	}

	RedisClient = client
}
