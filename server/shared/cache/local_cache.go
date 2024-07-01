package cache

import (
	"errors"
	"github.com/allegro/bigcache"
	"github.com/bytedance/sonic"
	"github.com/coocood/freecache"
	"github.com/patrickmn/go-cache"
	"reflect"
	"time"
)

type LocalCache interface {
	Set(key string, value interface{}) error
	Get(key string, value interface{}) error
	Delete(key string) error
	Clear() error
}

type LocalCacheConfig struct {
	Engine string
	Opts   []Option
}

type Config struct {
	ExpireTime int
	ClearTime  int
	MaxSizeMB  int
	Shards     int
}

func NewLocalCache(conf *LocalCacheConfig) (LocalCache, error) {
	cfg := &Config{}
	for _, opt := range conf.Opts {
		opt(cfg)
	}

	if f := GetConstructor(conf.Engine); f != nil {
		return f(cfg)
	}
	return nil, ErrNoSuchEngine
}

type LocalCacheConstructor func(config *Config) (LocalCache, error)

func GetConstructor(engine string) LocalCacheConstructor {
	switch engine {
	case EngineBigCache:
		return NewBigCache
	case EngineFreeCache:
		return NewFreeCache
	case EngineGoCache:
		return NewGoCache
	default:
		return nil
	}
}

func NewBigCache(config *Config) (LocalCache, error) {
	bigCache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(config.ExpireTime) * time.Second))
	if err != nil {
		return nil, err
	}
	return &BigCache{
		cache: bigCache,
	}, nil
}

type BigCache struct {
	cache *bigcache.BigCache
}

func (b *BigCache) Set(key string, value interface{}) error {
	val, err := sonic.Marshal(value)
	if err != nil {
		return errors.New("marshal value error")
	}
	return b.cache.Set(key, []byte(val))
}

func (b *BigCache) Get(key string, value interface{}) error {
	val, err := b.cache.Get(key)
	if err != nil {
		return err
	}
	err = sonic.Unmarshal(val, value)
	if err != nil {
		return err
	}
	return nil
}

func (b *BigCache) Delete(key string) error {
	return b.cache.Delete(key)
}

func (b *BigCache) Clear() error {
	return b.cache.Reset()
}

func NewFreeCache(config *Config) (LocalCache, error) {
	return &FreeCache{
		cache: freecache.NewCache(config.MaxSizeMB * 1024 * 1024),
	}, nil
}

type FreeCache struct {
	cache      *freecache.Cache
	expireTime int
}

func (f *FreeCache) Set(key string, value interface{}) error {
	val, err := sonic.Marshal(value)
	if err != nil {
		return errors.New("marshal value error")
	}
	return f.cache.Set([]byte(key), []byte(val), f.expireTime)
}

func (f *FreeCache) Get(key string, value interface{}) error {
	val, err := f.cache.Get([]byte(key))
	if err != nil {
		return err
	}
	err = sonic.Unmarshal(val, value)
	if err != nil {
		return err
	}
	return nil
}

func (f *FreeCache) Delete(key string) error {
	affected := f.cache.Del([]byte(key))
	if !affected {
		return errors.New("key not found")
	}
	return nil
}

func (f *FreeCache) Clear() error {
	f.cache.Clear()
	return nil
}

// NewGoCache creates a new GoCache instance
func NewGoCache(config *Config) (LocalCache, error) {
	return &GoCache{
		cache:      cache.New(time.Duration(config.ExpireTime)*time.Second, time.Duration(config.ClearTime)*time.Second),
		expireTime: time.Duration(config.ExpireTime) * time.Second,
	}, nil
}

type GoCache struct {
	cache      *cache.Cache
	expireTime time.Duration
}

func (g *GoCache) Set(key string, value interface{}) error {
	g.cache.Set(key, value, g.expireTime)
	return nil
}

func (g *GoCache) Get(key string, value interface{}) error {
	val, found := g.cache.Get(key)
	if !found {
		return errors.New("key not found")
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("value must be a non-nil pointer")
	}
	if v.Elem().Kind() == reflect.String {
		strVal, ok := val.(string)
		if !ok {
			return errors.New("type assertion to string failed, cache value is not a string")
		}
		v.Elem().SetString(strVal)
	} else {
		v.Elem().Set(reflect.ValueOf(val).Elem())
	}
	return nil
}

func (g *GoCache) Delete(key string) error {
	g.cache.Delete(key)
	return nil
}

func (g *GoCache) Clear() error {
	g.cache.Flush()
	return nil
}

type Option func(*Config)

func WithExpireTime(expireTime int) Option {
	return func(config *Config) {
		config.ExpireTime = expireTime
	}
}

func WithClearTime(clearTime int) Option {
	return func(config *Config) {
		config.ClearTime = clearTime
	}
}

func WithMaxSizeMB(maxSizeMB int) Option {
	return func(config *Config) {
		config.MaxSizeMB = maxSizeMB * 1024 * 1024
	}
}

func WithShards(shards int) Option {
	return func(config *Config) {
		config.Shards = shards
	}
}

// engine category
const (
	EngineBigCache  = "bigcache"
	EngineFreeCache = "freecache"
	EngineGoCache   = "gocache"
)

var ErrNoSuchEngine = errors.New("no such engine")
