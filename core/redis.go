package core

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/server-gin/config"
)

func ConnectRedisServer(rdc *config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     rdc.Addr,
		Password: rdc.Password, // no password set
		DB:       rdc.Db,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		return client, err
	}

	return client, nil
}
