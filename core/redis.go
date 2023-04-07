package core

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/zhaogongchengsi/starter-gin/config"
)

var ErrRedisNot = errors.New("error:redis service is not configured, missing address")

func ConnectRedisServer(conf *config.Config) (*redis.Client, error) {
	rdc := conf.Redis

	if len(rdc.Addr) < 1 {
		return nil, ErrRedisNot
	}

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
