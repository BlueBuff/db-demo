package cache

import (
	"time"
	"github.com/go-redis/redis"
)

type Serialize interface {
	Serialization() ([]byte, error)
	UnSerialization(str []byte) error
	ToString() string
	GetStringKey() string
}

type Cache interface {
	Set(key string, ser Serialize, time time.Duration) error
	Get(key string, ser Serialize) error
}

type UserRedisCache struct {
	redisClient *redis.Client
}

func NewUserRedisCache(redisClient *redis.Client) Cache {
	userRedisCache := new(UserRedisCache)
	userRedisCache.redisClient = redisClient
	return userRedisCache
}

func (cache *UserRedisCache) Set(key string, ser Serialize, time time.Duration) error {
	if key == "" {
		key = ser.GetStringKey()
	}
	err := cache.redisClient.Set(key, ser.ToString(), time).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cache *UserRedisCache) Get(key string, ser Serialize) error {
	if key == "" {
		key = ser.GetStringKey()
	}
	str, err := cache.redisClient.Get(key).Result()
	if err != nil {
		return err
	}
	return ser.UnSerialization([]byte(str))
}
