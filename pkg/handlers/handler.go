package handlers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
)

type handler struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func New(db *gorm.DB, redis *redis.Client) handler {
	return handler{db, redis}
}

func (h handler) getDataFromRedis(key string) (redisResult string, err error) {
	redisResult, err = h.Redis.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		log.Printf("unable to GET data from Redis. error: %v", err)
		return
	}
	return
}

func (h handler) setDataToRedis(key string, data string) (err error) {
	err = h.Redis.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		log.Printf("unable to SET data to Redis. error: %v\n", err)
		return
	}
	return
}
