package db

import (
	"github.com/go-redis/redis/v8"
	"log"
)

func InitRedisClient() *redis.Client {
	host := "localhost:6379"
	password := ""

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0, // use default DB
	})
	log.Println("Redis Client Initialized...")
	return client
}
