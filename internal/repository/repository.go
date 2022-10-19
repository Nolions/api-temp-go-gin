package repository

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type Repository struct {
	idx         int8
	db          *xorm.EngineGroup
	redis       *redis.Client
	cache       *cache.Cache
	cachePrefix string
}

func New(db *xorm.EngineGroup, redis *redis.Client, cachePrefix string) Repository {
	return Repository{
		db:    db,
		redis: redis,
		cache: cache.New(&cache.Options{
			Redis: redis,
		}),
		cachePrefix: cachePrefix,
	}
}
