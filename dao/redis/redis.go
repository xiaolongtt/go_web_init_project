package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func InitRedis() (err error) {
	addr := fmt.Sprintf("%s:%d", viper.GetString("datasource.redis.host"), viper.GetInt("datasource.redis.port"))
	fmt.Println("redis addr:", addr)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		//Password: viper.GetString("datasource.redis.password"),
		DB:       viper.GetInt("datasource.redis.database"),
		PoolSize: viper.GetInt("datasource.redis.pool_size"),
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
