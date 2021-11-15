package rdshand

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestSimpleSet(t *testing.T) {
	hand := New(getRds())
	rds := hand.GetClient()
	key := RdKey(`simple:%d`, time.Now().Unix())
	_, err := rds.Get(ctx, key).Bool()
	assert.Equal(t, err, redis.Nil)

	// set and get
	rds.Set(ctx, key, true, 30*time.Minute)
	val, err := rds.Get(ctx, key).Bool()
	assert.NotEqual(t, err, redis.Nil)
	assert.Equal(t, val, true)

	// set again
	rds.Set(ctx, key, false, 30*time.Minute)
	val, err = rds.Get(ctx, key).Bool()
	assert.NotEqual(t, err, redis.Nil)
	assert.Equal(t, val, false)
}

func TestHashSet(t *testing.T) {
	hand := New(getRds())
	rds := hand.GetClient()
	key := RdKey(`hashexp:%d`, time.Now().Unix())

	// hgetall
	var us UserList
	ma := rds.HGetAll(ctx, key).Val()
	for k, v := range ma {
		t.Logf("k=>%v, v=>%+v", k, v)
		uv := &User{}
		err := json.Unmarshal([]byte(v), uv)
		assert.Nil(t, err)
		us = append(us, uv)
	}

	// hgetall
	var usm UserList
	ma = rds.HGetAll(ctx, key).Val()
	for k, v := range ma {
		t.Logf("k=>%v, v=>%+v", k, v)
		uv := &User{}
		err := json.Unmarshal([]byte(v), uv)
		assert.Nil(t, err)
		usm = append(usm, uv)
	}
	t.Logf("usm:+%+v", usm)
}

type User struct {
	ID   uint64 `redis:"id"`
	Name string `redis:"name"`
}

type UserList []*User

func (u *UserList) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}
