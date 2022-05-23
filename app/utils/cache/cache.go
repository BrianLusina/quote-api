package cache

import (
	"context"
	"net"
	"quote/api/app/config"
	"strconv"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	address  string
	username string
	password string
	db       int
	client   *redis.Client
	ctx      context.Context
	cache    *cache.Cache
}

func New(ctx context.Context, config config.CacheConfig) *Cache {
	db, _ := strconv.Atoi(config.Db)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       db,
	})

	cache := cache.New(&cache.Options{
		Redis: redisClient,

		// Cache 10k keys for 1 minute.
		LocalCache: cache.NewTinyLFU(10000, time.Minute),
	})

	return &Cache{
		address:  config.Host,
		username: config.Username,
		password: config.Password,
		db:       db,
		ctx:      ctx,
		client:   redisClient,
		cache:    cache,
	}
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func (c *Cache) Set(key string, value any) error {
	err := c.cache.Set(&cache.Item{
		Key:   key,
		Value: value,
		Do: func(*cache.Item) (interface{}, error) {
			return &value, nil
		},
	})
	return err
}
