package infrastructure

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	ctx    context.Context
	client *redis.Client
}

func ConnectRedis() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Clear all data in Redis
	if err := client.FlushAll(context.Background()).Err(); err != nil {
		return nil, err
	}

	redisClient := &RedisClient{
		ctx:    context.Background(),
		client: client,
	}

	return redisClient, nil
}

func (rc *RedisClient) Set(key string, val interface{}) error {
	return rc.client.Set(rc.ctx, key, val, 60*time.Minute).Err()
}
func (rc *RedisClient) Get(key string) (map[string]string, error) {
	data := make(map[string]string)

	jsonData, err := rc.client.Get(rc.ctx, key).Result()
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return data, err
	}

	return data, nil
}

func (rc *RedisClient) KeyExist(key string) (int64, error) {
	return rc.client.Exists(rc.ctx, key).Result()
}
func (rc *RedisClient) ClearKeys() (int64, error) {
	return rc.client.Del(rc.ctx, "message").Result()
}
