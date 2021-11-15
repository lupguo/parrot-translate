package rdshand

import (
	"context"

	"github.com/pkg/errors"
)

// ICurrencyLocker 分布式并发锁接口，支持通过业务操作唯一Key进行短时间锁定
type ICurrencyLocker interface {
	DistrbLock(ctx context.Context, lockKey string, ttl string) error
	DistrbUnLock(ctx context.Context, busKey string) error
}

// DistrbLock 通过Redis增加分布式记录锁设计，防止业务操作短时并发
func (h *RdsHand) DistrbLock(ctx context.Context, lockKey string, ttl string) error {
	expire := MustParseKeyDuration(ttl)
	if !h.Rds.SetNX(ctx, lockKey, true, expire).Val() {
		return errors.New("concurrency lock crash")
	}
	return nil
}

// DistrbUnLock 分布式解锁操作
func (h *RdsHand) DistrbUnLock(ctx context.Context, busKey string) error {
	return h.Rds.Del(ctx, busKey).Err()
}
