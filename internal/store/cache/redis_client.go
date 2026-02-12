package cache

import (
    "context"
    "github.com/redis/go-redis/v9"
    "time"
)

var ctx = context.Background()

type RedisClient struct {
    Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password, 
        DB:       db,      
    })

    return &RedisClient{Client: rdb}
}

func (r *RedisClient) Set(key, value string, expiration time.Duration) error {
    return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
    return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) Delete(key string) error {
    return r.Client.Del(ctx, key).Err()
}