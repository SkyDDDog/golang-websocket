package cache

import (
	"demo04/config"
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client

func InitCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDbNumber,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Redis链接失败", err)
		panic(err)
	}
	RedisClient = client

}
