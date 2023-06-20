package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// a custom wrapper for common actions
type Cache struct {
	client  *redis.Client   // the redis official client
	context context.Context // mongodb has its own context, redis also its own required context
}

func NewCache() *Cache {
	cache := &Cache{
		client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_CONNECTION"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
			Username: "default",
		}),
		context: context.TODO(),
	}
	fmt.Println("Connected to cache successfully")
	return cache
}

// closes the client and releases resources
func (cache *Cache) Disconnect() {
	cache.client.Close()
}

// saves the value on the redis database, returns error if something went wrong.
// all data stored will be remembered for 60 minutes
func (cache *Cache) Set(key string, value interface{}) error {
	err := cache.client.Set(cache.context, key, value, 60*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

// retrieves value required by the key, returns error when key doesn't exists or some error happened
func (cache *Cache) Get(key string) (interface{}, error) {
	val, err := cache.client.Get(cache.context, key).Result()
	if err == redis.Nil || err != nil {
		return nil, err
	}

	return val, nil
}

// deletes a single element in cache
func (cache *Cache) Delete(key string) error {
	return cache.client.Del(cache.context, key).Err()
}
