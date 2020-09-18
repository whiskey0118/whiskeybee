package tools

import "github.com/go-redis/redis"

var RedisConn *redis.Client

func init() {
	opt, err := redis.ParseURL("redis://192.168.3.115:6379/0")
	if err != nil {
		panic(err)
	}
	RedisConn = redis.NewClient(opt)

}
