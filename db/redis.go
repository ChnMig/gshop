package db

import (
	"gshop/tool"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// RDB Redis client
var RDB *redis.Client

// ConnRedis connect to redis database
func ConnRedis(address string) error {
	// https://redis.uptrace.dev/#connecting-to-redis-server
	opt, err := redis.ParseURL(address)
	if err != nil {
		tools.Log.Panic("There was a problem parsing the redis address", zap.Error(err))
		return err
	}
	RDB = redis.NewClient(opt)
	// ping
	_, err = RDB.Ping(RDB.Context()).Result()
	if err != nil {
		tools.Log.Panic("Redis Connect Error", zap.Error(err))
	}
	return nil
}
