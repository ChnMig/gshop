package db

import (
	"github.com/chnmig/gshop/tools"
	"github.com/go-redis/redis/v8"
)

// RDB Redis client
var RDB *redis.Client

// ConnRedis connect to redis database
func ConnRedis(address string) error {
	// https://redis.uptrace.dev/#connecting-to-redis-server
	opt, err := redis.ParseURL(address)
	if err != nil {
		tools.Log.Panic("There was a problem parsing the redis address")
		return err
	}
	RDB = redis.NewClient(opt)
	return nil
}
