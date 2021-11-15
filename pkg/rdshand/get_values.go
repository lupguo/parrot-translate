package rdshand

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// GetString 获取缓存的字符串值，若没有找返回空，若redis失败返回redis操作错误，其他情况正常返回
func (h *RdsHand) GetString(ctx context.Context, format string, injects ...interface{}) (string, error) {
	key := RdKey(format, injects...)
	cmd := h.Rds.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return cmd.Val(), nil
}
