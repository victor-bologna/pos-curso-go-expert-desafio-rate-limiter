package limiter

import (
	"context"
	"fmt"
	"time"

	"github.com/pos-curso-go-expert-desafio-rate-limiter/config"
	"github.com/redis/go-redis/v9"
)

type State bool

const (
	Deny  State = false
	Allow State = true
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

// TODO: Create config file
func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddr,
		Password: config.AppConfig.RedisPass,
		DB:       0,
	})
	return &RedisClient{
		client: client,
		ctx:    context.Background(),
	}
}

func (rc *RedisClient) NextRequest(key string, maximumReq int, timeout int) (State, string) {
	redisValue, err := rc.client.Get(rc.ctx, key).Int()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("log -> key:", key)
	fmt.Println("log -> redisValue:", redisValue)

	if err != nil {
		if err == redis.Nil {
			rc.client.Set(rc.ctx, key, 1, time.Second).Result()
			return Allow, ""
		}
		return Deny, err.Error()
	}

	if redisValue > maximumReq {
		rc.client.Set(rc.ctx, key, redisValue, time.Duration(timeout)*time.Second).Result()
		return Deny, "you have reached the maximum number of requests or actions allowed within a certain time frame"
	}

	rc.client.Incr(rc.ctx, key).Result()
	return Allow, ""
}
