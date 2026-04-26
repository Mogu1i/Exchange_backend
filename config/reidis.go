package config

import (
	"exchangeapp/global"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {
	addr := Appconfig.Redis.Addr
	if addr == "" {
		addr = "localhost:6379"
	}

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       Appconfig.Redis.DB, //表示默认数据库
		Password: Appconfig.Redis.Password,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis,got error:%v", err)
	}
	global.RedisDB = RedisClient
}
