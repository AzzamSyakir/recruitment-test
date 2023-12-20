package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Kosongkan jika tidak ada password
		DB:       0,  // Gunakan default DB
	})
	if client == nil {
		fmt.Println("Tidak dapat membuat klien Redis")

		return nil
	}

	return client
}
