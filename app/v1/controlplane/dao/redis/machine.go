package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	redisclient "github.com/ztalab/ZAManager/pkg/redis"
)

type Machine struct {
	c      *gin.Context
	ctx    context.Context
	client *redis.Client
	Prefix string
}

func NewMachine(c *gin.Context) *Machine {
	return &Machine{
		c:      c,
		ctx:    context.Background(),
		client: redisclient.Client,
		Prefix: "zta:",
	}
}

func (m *Machine) SetLoginHash(key, hash string) (err error) {
	// 先将之前的machine取出来
	result, err := m.client.Get(m.ctx, fmt.Sprintf("%s%s", m.Prefix, key)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return
	}
	if len(result) > 0 {
		// 删除鉴权的redis数据
		if err = m.client.Del(m.ctx, result).Err(); err != nil {
			return
		}
	}
	if err = m.client.Set(m.ctx,
		fmt.Sprintf("%s%s", m.Prefix, key), fmt.Sprintf("%s%s", m.Prefix, hash), time.Second*60*5).Err(); err != nil {
		return
	}
	if err = m.client.Set(m.ctx, fmt.Sprintf("%s%s", m.Prefix, hash), "", time.Second*60*5).Err(); err != nil {
		return
	}
	return
}

func (m *Machine) GetLoginHash(hash string) (exist bool, data string, err error) {
	result, err := m.client.Get(context.Background(),
		fmt.Sprintf("%s%s", m.Prefix, hash)).Result()
	if errors.Is(err, redis.Nil) {
		return false, data, nil
	} else if err != nil {
		return false, data, err
	} else {
		return true, result, nil
	}
}

func (m *Machine) SetMachineHash(key, value string) (err error) {
	return m.client.Set(m.ctx, fmt.Sprintf("%s%s", m.Prefix, key), value, time.Second*60+5).Err()
}

func (m *Machine) GetMachineHash(key, value string) (data string, err error) {
	return
}
