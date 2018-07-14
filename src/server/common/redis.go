package common

import (
	"github.com/go-redis/redis"
	"fmt"
)

var RedisClient *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ConfigurationContext.Redis.Host, ConfigurationContext.Redis.Port),
		Password: ConfigurationContext.Redis.Password,
		DB:       ConfigurationContext.Redis.DB,
	})
	if pong, err := client.Ping().Result(); err != nil {
		panic(err)
	} else {
		fmt.Println("redis ping=>", pong)
	}
	RedisClient = client
}
