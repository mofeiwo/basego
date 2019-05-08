package exdt

import (
	"sync"
	"github.com/go-redis/redis"
	"strconv"
	"site/base"
)

type RedisClient struct {
}

var instance *redis.Client
var once sync.Once

func (r *RedisClient) getInstance() *redis.Client {
	once.Do(func() {
		redisConf := base.RedisConfig
		db, _ := strconv.Atoi(redisConf["db"])
		instance = redis.NewClient(&redis.Options{
			Addr:     redisConf["addr"],
			Password: redisConf["password"],
			DB:       db,
		})
	})
	return instance
}