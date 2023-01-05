package common

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewCacheConnection(cacheUrl string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cacheUrl,
		Password: password,
		DB:       0,
	})

	// create context
	ctx := context.Background()

	// ping
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("redis connection error: ", err)
	}

	return rdb
}
