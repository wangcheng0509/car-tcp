package core

import (
	"car.open.service/conf"
	"github.com/redis/go-redis/v9"
)

// NewRedisCache redis cache
func NewRedisCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Redis.Addr,
		Password: conf.Conf.Redis.Password,
		DB:       conf.Conf.Redis.DB,
	})
	return rdb
}
