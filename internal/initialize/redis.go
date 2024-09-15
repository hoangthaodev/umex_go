package initialize

import (
	"github.com/redis/go-redis/v9"
	"umex.com/global"
)

func InitRedis() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
