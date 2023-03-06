package rediscache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"test3/service"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Get(ctx context.Context, key string, result interface{}) error {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("you must pass a pointer for result data")
	}

	cacheData, err := r.client.Get(ctx, key).Bytes()

	if err != nil && errors.Is(err, redis.Nil) {
		return service.CacheKeyNotExist
	} else if err != nil {
		logrus.WithError(err).Errorf("Failed to get cache with key: %s", key)
		return err
	}

	err = json.Unmarshal(cacheData, result)

	if err != nil {
		return fmt.Errorf("failed to unmarshal data from key: %s with cacheData: %v", key, string(cacheData))
	}

	return nil
}

func (r *RedisCache) Put(ctx context.Context, key string, data interface{}, ttl time.Duration) error {
	value, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) DeleteKeys(ctx context.Context, keys []string) error {
	return r.client.Del(ctx, keys...).Err()
}
