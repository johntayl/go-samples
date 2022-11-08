package main

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	redisClient *redis.Client
	mux         sync.Mutex
}

func NewRedisCache(url, port, password string) *RedisCache {
	return &RedisCache{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     url + ":" + port,
			Password: password,
			DB:       0, // use the default DB
		}),
	}
}

// Ping
func (r *RedisCache) Ping() error {
	_, err := r.redisClient.Ping().Result()
	return err
}

// Get
func (r *RedisCache) Get(key string) (string, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.redisClient.Get(key).Result()
}

// Set
func (r *RedisCache) Set(key, value string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.redisClient.Set(key, value, 0).Err()
}

// Delete
func (r *RedisCache) Delete(key string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.redisClient.Del(key).Err()
}

// Increment
func (r *RedisCache) Increment(key string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.redisClient.Incr(key).Err()
}

// Decrement
func (r *RedisCache) Decrement(key string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.redisClient.Decr(key).Err()
}

// Flushdb
func (r *RedisCache) Flushdb() error {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.redisClient.FlushDB().Err()
}

func main() {
	redisCache := NewRedisCache("localhost", "6379", "")

	// Ping
	err := redisCache.Ping()
	if err != nil {
		panic(err)
	}

	// Set
	fmt.Println("Setting key value in Redis")
	err = redisCache.Set("key", "my value")
	if err != nil {
		panic(err)
	}

	// Get
	value, err := redisCache.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Got key from Redis: " + value)

	// Delete
	err = redisCache.Delete("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted key from Redis")
}
