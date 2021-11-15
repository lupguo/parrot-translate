package rdshand

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=./cache.go -package rdshand -destination cache_mock.go

// IRedisHand redis interface
type IRedisHand interface {
	GetClient() *redis.Client
	PipeWrite(ctx context.Context, objs []interface{}, kfn KeyFunc, ttl time.Duration) error
	PipeRead(ctx context.Context, keys []string) (map[string]string, error)
	GetString(ctx context.Context, keyFormat string, injects ...interface{}) (string, error)
	DelKeys(ctx context.Context, keys ...string) error
}

// RdsHand RedisCache
type RdsHand struct {
	Rds *redis.Client
}

// New create redis obj
func New(client *redis.Client) *RdsHand {
	if client == nil {
		panic(client)
	}
	return &RdsHand{Rds: client}
}

// GetClient get client from RdsHand obj
func (h *RdsHand) GetClient() *redis.Client {
	return h.Rds
}
