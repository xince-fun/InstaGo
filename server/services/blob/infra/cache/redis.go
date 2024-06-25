package cache

import (
	"context"
	"github.com/go-redis/cache/v9"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/xince-fun/InstaGo/server/services/blob/conf"
	"time"
)

var CacheManagerSet = wire.NewSet(
	NewCache,
	NewRedisManager,
)

type CacheItem interface {
	GetID() string
	IsDirty() bool
}

// RedisManager 二级缓存
type RedisManager struct {
	cache *cache.Cache
}

func NewRedisManager(cache *cache.Cache) *RedisManager {
	return &RedisManager{
		cache: cache,
	}
}

func NewCache() *cache.Cache {
	c := conf.GlobalServerConf.RedisConfig
	addr := make(map[string]string)
	for _, sc := range c.RedisServerConfig {
		addr[sc.Name] = sc.Addr
	}

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: addr,
	})

	return cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(c.LocalCacheTime, time.Minute),
	})
}

func (r *RedisManager) Get(ctx context.Context, key string, value interface{}) error {
	if err := r.cache.Get(ctx, key, value); err != nil {
		return err
	}
	return nil
}

func (r *RedisManager) Set(ctx context.Context, key string, value CacheItem) error {
	return r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	})
}

type BlobItem struct {
	BlobID     string
	ObjectName string
}

func (b *BlobItem) GetID() string {
	return b.BlobID
}

func (b *BlobItem) IsDirty() bool {
	return b.ObjectName != ""
}
