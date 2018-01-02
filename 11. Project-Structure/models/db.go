package models

import (
	"github.com/go-redis/redis"
)
var client *redis.Client

func Init() {
	//redis connection
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}