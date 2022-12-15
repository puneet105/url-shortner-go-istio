package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()

func CreateClient(dbno int) *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB: dbno,
	})

	result, err := rdb.Ping(Ctx).Result()
	fmt.Println(result)

	if err != nil {
		fmt.Println(rdb, err)
	}

	return rdb
}

