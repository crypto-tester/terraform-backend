package redis

import (
	"context"
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	redis "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"

	"github.com/crypto-tester/terraform-backend/internal"
)

var ctx = context.Background()

func NewPool() *redigo.Pool {
	viper.SetDefault("redis_addr", "localhost:6379")

	return &redigo.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", viper.GetString("redis_addr"))
			if err != nil {
				if c != nil {
					c.Close()
				}

				return nil, err
			}

			pass, err := internal.SecretEnvOrFile("redis_password", "redis_password_file")
			if err != nil {
				return nil, fmt.Errorf("getting redis password: %w", err)
			}

			if pass != "" {
				if _, err := c.Do("AUTH"); err != nil {
					c.Close()

					return nil, err
				}
			}

			return c, nil
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")

			return err
		},
	}
}

func NewClient() (*redis.Client, error) {
	viper.SetDefault("redis_addr", "localhost:6379")
	viper.SetDefault("redis_password", "")
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis_addr"),
		Password: viper.GetString("redis_password"),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("failed to connect to redis server: %s\n", err.Error())
		return nil, err
	}

	return rdb, nil
}
