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

func (rc *RedisClient) NextRequest(key string, maximumReq int, timeout int) State {
	redisValue, err := rc.client.Get(rc.ctx, key).Int()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("log -> key:", key)
	fmt.Println("log -> redisValue:", redisValue)

	if err != nil || redisValue >= maximumReq {
		if err == redis.Nil {
			fmt.Println("log -> Nova chave")
			rc.client.Set(rc.ctx, key, 1, time.Second).Result()
			return Allow
		}
		fmt.Println("log -> Timeout:", timeout)
		rc.client.Set(rc.ctx, key, redisValue, time.Duration(timeout)*time.Second).Result()
		return Deny
	}

	rc.client.Incr(rc.ctx, key).Result()
	return Allow
}
