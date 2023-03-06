package rediscache

import (
	"context"
	"errors"
	"testing"

	"test3/service"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/require"
)

func TestRedisCache_Get(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenNonPointerGiven", func(t *testing.T) {
		t.Parallel()
		mockedClient, _ := redismock.NewClientMock()
		svc := NewRedisCache(mockedClient)
		err := svc.Get(context.Background(), "key-1", "")
		require.Error(t, err)
	})
	t.Run("ShouldReturnCacheKeyNotExistErr_WhenKeyWasNotFound", func(t *testing.T) {
		t.Parallel()
		// ARRANGE
		mockedClient, redisMock := redismock.NewClientMock()
		// -- WhenKeyWasNotFound
		redisMock.ExpectGet("key-1").SetErr(redis.Nil)

		// ACTION
		svc := NewRedisCache(mockedClient)
		result := struct {
			Key string `json:"key,omitempty"`
		}{}
		err := svc.Get(context.Background(), "key-1", &result)

		// ASSERT
		// -- ShouldReturnNoErr
		require.ErrorIs(t, err, service.CacheKeyNotExist)
	})
	t.Run("ShouldReturnErr_WhenRedisReturnOtherError", func(t *testing.T) {
		t.Parallel()
		// ARRANGE
		mockedClient, redisMock := redismock.NewClientMock()
		// -- WhenRedisReturnOtherError
		redisMock.ExpectGet("key-1").SetErr(errors.New("other than redis: nil"))

		// ACTION
		svc := NewRedisCache(mockedClient)
		err := svc.Get(context.Background(), "key-1", "")

		// ASSERT
		// -- ShouldReturnErr
		require.Error(t, err)
	})
	t.Run("ShouldSetResult_WhenKeyWasFound", func(t *testing.T) {
		t.Parallel()
		// ARRANGE
		mockedClient, redisMock := redismock.NewClientMock()
		// -- WhenKeyWasFound
		redisMock.ExpectGet("key-1").SetVal(`{"key":"value"}`)

		// ACTION
		svc := NewRedisCache(mockedClient)
		type DummySchema struct {
			Key string `json:"key"`
		}
		result := DummySchema{}
		err := svc.Get(context.Background(), "key-1", &result)

		// ASSERT
		require.NoError(t, err)
		// -- ShouldSetResult
		require.Equal(t, DummySchema{Key: "value"}, result)
	})
	t.Run("ShouldReturnErr_WhenUnmarshalFailed", func(t *testing.T) {
		t.Parallel()
		// ARRANGE
		mockedClient, redisMock := redismock.NewClientMock()
		redisMock.ExpectGet("key-1").SetVal(`{"key":"value"}`)

		// ACTION
		svc := NewRedisCache(mockedClient)
		result := "" // WhenUnmarshalFailed: unmarshall will fail when given non struct result variable
		err := svc.Get(context.Background(), "key-1", &result)

		// ASSERT
		// -- ShouldReturnErr
		require.Error(t, err)
	})
}
