package redis

import (
	"context"
	"fmt"
	"time"
	"web_app_go/settings"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(cfg *settings.RedisConfig) (err error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("init redis failed, err:%w", err))
	}
	return
}

func Close() {
	_ = RedisClient.Close()
}
