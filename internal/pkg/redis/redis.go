package redis

import (
	"fmt"

	"github.com/asyauqi1511/go-workshop-service/internal/entity"
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(config entity.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Hostname, config.Port),
		Password: config.Password,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
