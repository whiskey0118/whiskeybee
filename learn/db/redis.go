package main

import (
	"github.com/go-redis/redis"
)

var rc *redis.Client

func init() {
	opt, err := redis.ParseURL("redis://192.168.3.115:6379/0")
	if err != nil {
		panic(err)
	}
	rc = redis.NewClient(opt)
}
