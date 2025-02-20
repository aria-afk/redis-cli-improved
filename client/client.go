// Redis client interface from the go-redis package
package client

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	DB  *redis.Client
	ctx context.Context
}

func NewRedis(context context.Context, address string, password string, db int) (*Redis, error) {
	rdb := &Redis{
		ctx: context,
	}

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	err := client.Set(rdb.ctx, "health-check", "", time.Second*1).Err()
	if err != nil {
		return rdb, err
	}
	rdb.DB = client

	return rdb, nil
}
