package db

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/carepollo/librecode/utils"
	"github.com/go-redis/redis/v8"
)

var (
	cache    *redis.Client
	redisCtx context.Context = context.TODO()
)

// create instance of redis and ping it
func openCache() {
	cache = redis.NewClient(&redis.Options{
		Addr:     utils.GlobalEnv.Storage.Cache.Connection,
		Password: utils.GlobalEnv.Storage.Cache.Password,
	})

	_, err := cache.Ping(redisCtx).Result()
	if err != nil {
		panic("could not connect to cache: " + err.Error())
	}
	log.Println("Successfully connected to Redis instance")
}

// close connection to redis instance, make sure to run only at the end of the program
func closeCache() {
	if err := cache.Close(); err != nil {
		log.Println(err)
	}
}

// cache generic function as a shorthand to SET command.
// saves the value on the redis database, returns error if something went wrong.
// stored data will be remembered by the given time.
// value param must be encodable in json
func remember(key string, value interface{}, expireAt time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return cache.Set(redisCtx, key, data, expireAt).Err()
}

// cache generic function as a shorthand to GET command
// retrieves value by the key and returns error when key doesn't exists
func retrieve(key string) ([]byte, error) {
	val, err := cache.Get(redisCtx, key).Bytes()
	if err == redis.Nil || err != nil {
		return nil, err
	}

	return val, nil
}

// cache type unsafe function as a shorthand to DEL command, must cast type on concrete implementation.
// deletes a single element from cache
func forget(key string) error {
	return cache.Del(redisCtx, key).Err()
}
