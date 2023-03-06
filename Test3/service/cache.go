package service

import (
	"context"
	"time"
)

type CacheError string

func (c CacheError) Error() string {
	return string(c)
}

const CacheKeyNotExist = CacheError("key not exist")

type Cache interface {
	Get(ctx context.Context, key string, result interface{}) error
	Put(ctx context.Context, key string, data interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	DeleteKeys(ctx context.Context, keys []string) error
}
