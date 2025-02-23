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

type RedisOptions struct {
	Host   string // h   REDISCLI_HOST
	Auth   string // a   REDISCLI_AUTH
	User   string // usr REDISCLI_USER
	Uri    string // u   REDISCLI_URI
	Number int    // n   REDISCLI_NUM NOTE: This is just the base conn
	Port   int    // p   REDISCLI_PORT
	// TODO: dial timeout
	// TODO: TLS options
	// TODO: pooling/cluster?
}

func NewRedis(context context.Context, cliOpts RedisOptions, envOpts RedisOptions, defaultOpts RedisOptions) (*Redis, error) {
	rdb := &Redis{
		ctx: context,
	}

	// TODO: actually use the opts.
	// build URI string from opts
	// redis://user:[password]@host:port/dbnum
	// redis://host:port/dbnum if no auth
	// redis://user@host:port/dbnum?dial_timeout=5s
	// cli opts > env opts > defaultOpts

	opts, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		return rdb, err
	}
	client := redis.NewClient(opts)

	err = client.Set(rdb.ctx, "health-check", "", time.Second*1).Err()
	if err != nil {
		return rdb, err
	}
	rdb.DB = client

	return rdb, nil
}
