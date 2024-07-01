package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type CacheManager interface {
	GetObj(context.Context, string, interface{}) error
	GetObjOpt(context.Context, string, interface{}, bool) error
	SetObj(context.Context, string, interface{}) error
	SetObjOpt(context.Context, string, interface{}, bool) error
	Get(context.Context, string) (string, error)
	GetOpt(context.Context, string, bool) (string, error)
	Set(context.Context, string, string) error
	SetOpt(context.Context, string, string, bool) error
	Client() redis.UniversalClient
}

type CacheItem interface {
	GetID() string
	IsDirty() bool
}
