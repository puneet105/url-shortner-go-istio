package database

import (
	"context"
	"github.com/go-redis/redis"
	"os"
)

var Ctx = context.Background()

func CreateClient(dbno int) *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB: dbno,
	})

	return rdb
}

