package database

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
)
var (
	client *redis.Client
)
func Test_CreateClient(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("Error '%s' opening database connection", err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	t.Run("success", func(t *testing.T) {

	})
}
