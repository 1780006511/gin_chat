package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Rdb *redis.Client
)

const (
	PublishKey = "websocket"
)

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Rdb.Ping().Result err:", err)
		return
	}
	fmt.Println("初始化redis成功，pong", pong)
}

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) (err error) {
	err = Rdb.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(" Rdb.Publish err:", err)
		return
	}
	return
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Rdb.Subscribe(ctx, channel)
	fmt.Println("Subscribe....ctx:", ctx)

	msg, err := sub.ReceiveMessage(ctx)

	if err != nil {
		fmt.Println("sub.ReceiveMessage err:", err)
		return "", err
	}
	fmt.Println("Subscribe....", msg.Payload)
	return msg.Payload, err
}
