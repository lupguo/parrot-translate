package cache

import (
	"github.com/go-redis/redis/v8"
	"parrot-translate/pkg/rdshand"
)

// RedisInfra 通过redis实现，Redis的基础设施接口
type RedisInfra struct {
	rdsHand rdshand.IRedisHand
}

// NewRedisInfra 创建一个redis基础设施实现
func NewRedisInfra() *RedisInfra {
	rdsOpts := &redis.Options{}
	// todo 从yaml配置读取redis配置

	rdsClient := redis.NewClient(rdsOpts)
	return &RedisInfra{
		rdsHand: rdshand.New(rdsClient),
	}
}

func (r RedisInfra) GetTransContentByHashID(hashID string) (text string, err error) {
	panic("implement me")
}

func (r RedisInfra) SetTransContent(hashID string, content string) error {
	panic("implement me")
}
