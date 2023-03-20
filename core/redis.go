package core

import (
	"context"
	"errors"
	"time"

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

type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string) error

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string

	//Verify captcha's answer directly
	Verify(id, answer string, clear bool) bool
}

type RedisStore struct {
	client  *redis.Client
	timeout time.Duration
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client, 15}
}

func (R *RedisStore) Set(id string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), R.timeout*time.Minute)
	defer cancel()
	rdb := R.client
	return rdb.Set(ctx, id, value, time.Hour).Err()
}

func (R *RedisStore) Get(id string, clear bool) string {

	ctx, cancel := context.WithTimeout(context.Background(), R.timeout*time.Minute)
	defer cancel()

	value := R.client.Get(ctx, id).Val()
	if clear {
		R.client.Del(ctx, id)
	}
	return value
}

func (R *RedisStore) Verify(id, answer string, clear bool) bool {
	match := R.Get(id, clear) == answer
	return match
}
