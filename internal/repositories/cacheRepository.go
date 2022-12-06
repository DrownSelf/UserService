package repositories

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"

	configs "github.com/DrownSelf/UserService/internal/config"
)

type ICacheRepository interface {
	PutToken(ctx context.Context, token string) error
	DoesInvalidTokenExist(ctx context.Context, token string) bool
}

type CacheRepository struct {
	db      *redis.Client
	myCache *cache.Cache
}

func NewCacheRepo(config configs.Config) *CacheRepository {
	db := redis.NewClient(&redis.Options{
		Password: config.RedisPassword,
		Addr:     config.RedisHost,
	})
	cache := cache.New(&cache.Options{Redis: db})
	return &CacheRepository{db, cache}
}

func (r *CacheRepository) DestroyRepo(ctx context.Context) error {
	err := make(chan error)
	go func(out chan error) {
		if err := r.db.Close(); err != nil {
			out <- err
		}
	}(err)
	var dbError error
	select {
	case <-ctx.Done():
		return nil
	case dbError = <-err:
		return dbError
	}
}

func (r *CacheRepository) PutToken(ctx context.Context, token string) error {
	err := r.myCache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   token,
		Value: &struct{}{},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository) DoesInvalidTokenExist(ctx context.Context, token string) bool {
	return r.myCache.Exists(ctx, token)
}
