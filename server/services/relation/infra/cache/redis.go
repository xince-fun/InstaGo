package cache

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/xince-fun/InstaGo/server/services/relation/conf"
	scache "github.com/xince-fun/InstaGo/server/shared/cache"
	"log"
	"time"
)

var CacheManagerSet = wire.NewSet(
	scache.NewLocalCache,
	ProvideLocalCache,
	NewClient,
	NewCacheManager,
)

// CacheManager 二级缓存
type CacheManager struct {
	cache  scache.LocalCache
	client redis.UniversalClient
}

func NewCacheManager(cache scache.LocalCache, client redis.UniversalClient) *CacheManager {
	return &CacheManager{
		cache:  cache,
		client: client,
	}
}

func NewClient() redis.UniversalClient {
	c := conf.GlobalServerConf.RedisConfig
	addrs := make([]string, 0)
	for _, sc := range c.RedisServerConfig {
		addrs = append(addrs, sc.Addr)
	}
	opts := &redis.UniversalOptions{
		Addrs: addrs,
	}
	return redis.NewUniversalClient(opts)
}

func ProvideLocalCache() *scache.LocalCacheConfig {
	c := conf.GlobalServerConf.RedisConfig
	return &scache.LocalCacheConfig{
		Engine: c.CacheEngine,
		Opts: []scache.Option{
			scache.WithExpireTime(c.ExpireTime),
			scache.WithClearTime(c.ClearTime),
			scache.WithMaxSizeMB(c.MaxSizeMB),
		},
	}
}

// Get 从LocalCache和Redis获取字符串数据
func (r *CacheManager) Get(ctx context.Context, key string) (string, error) {
	return r.GetOpt(ctx, key, true)
}

// GetOpt 选择是否从LocalCache中获取字符串数据
func (r *CacheManager) GetOpt(ctx context.Context, key string, fromLocal bool) (value string, err error) {
	if fromLocal {
		if err = r.cache.Get(key, &value); err != nil {
			return "", err
		}
		// localCache hit
		if value != "" {
			return value, nil
		}
	}

	if value, err = r.client.Get(ctx, key).Result(); err != nil {
		if err == redis.Nil {
			return "", ErrCacheNotFound
		}
		return "", err
	}

	if fromLocal {
		// redis hit, set localCache
		if err = r.cache.Set(key, value); err == nil {
			return value, nil
		}
		// 没写回localCache，不影响返回值
		return value, ErrWriteLocalErr
	}
	return value, nil
}

func (r *CacheManager) Set(ctx context.Context, key, value string) error {
	return r.SetOpt(ctx, key, value, true)
}

func (r *CacheManager) SetOpt(ctx context.Context, key, value string, toRedis bool) (err error) {
	if err = r.cache.Set(key, value); err != nil {
		return err
	}
	if toRedis {
		if err = r.client.Set(ctx, key, value, 10*time.Minute).Err(); err != nil {
			return err
		}
	}
	return nil
}

// GetObj 从LocalCache和Redis获取对象数据
func (r *CacheManager) GetObj(ctx context.Context, key string, value interface{}) error {
	return r.GetObjOpt(ctx, key, value, true)
}

// GetObjOpt 选择是否从LocalCache中获取对象数据
func (r *CacheManager) GetObjOpt(ctx context.Context, key string, value interface{}, fromLocal bool) (err error) {
	obj, ok := value.(scache.CacheItem)
	if !ok {
		return errors.New("value does not implement CacheItem")
	}
	cacheKey := key + obj.GetID()
	if fromLocal {
		if err = r.cache.Get(cacheKey, value); err == nil {
			return nil
		}
	}
	if err = r.client.HGetAll(ctx, cacheKey).Scan(value); err != nil {
		return err
	}

	obj, ok = value.(scache.CacheItem)
	if !ok {
		return errors.New("value does not implement CacheItem")
	}
	if fromLocal {
		if obj.IsDirty() {
			// redis hit, set localCache
			if err = r.cache.Set(cacheKey, value); err == nil {
				return nil
			}
		} else {
			return ErrCacheNotFound
		}
	}
	if !obj.IsDirty() {
		return ErrCacheNotFound
	}
	return nil
}

// SetObj 默认写入LocalCache和Redis对象数据
func (r *CacheManager) SetObj(ctx context.Context, key string, value interface{}) error {
	return r.SetObjOpt(ctx, key, value, true)
}

// SetObjOpt 选择是否写入Redis对象数据
func (r *CacheManager) SetObjOpt(ctx context.Context, key string, value interface{}, toRedis bool) (err error) {
	obj, ok := value.(scache.CacheItem)
	if !ok {
		return errors.New("value does not implement CacheItem")
	}
	cacheKey := key + obj.GetID()
	if err = r.cache.Set(cacheKey, obj); err != nil {
		return err
	}

	if toRedis {
		if err = r.client.HSet(ctx, cacheKey, obj).Err(); err != nil {
			log.Default().Println(err)
			return err
		}
	}
	return nil
}

func (r *CacheManager) Client() redis.UniversalClient {
	return r.client
}

type RelationItem struct {
	UserID string `redis:"user_id"`
	Count  int32  `redis:"count"`
}

func (r *RelationItem) GetID() string {
	return r.UserID
}

func (r *RelationItem) IsDirty() bool {
	return r.Count != 0
}

type IDList []string

func (u IDList) GetID() string {
	return "user_list"
}

func (u IDList) IsDirty() bool {
	return len(u) != 0
}

var (
	ErrCacheNotFound = errors.New("cache not found")
	ErrWriteLocalErr = errors.New("write local cache error")
)
