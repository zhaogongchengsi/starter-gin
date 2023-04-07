package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

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
