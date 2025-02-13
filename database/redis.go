// database/redis.go
package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "52.5.28.74:6379", // rEDIS DIRECCTION
		Password: "",                // PASSWORD
		DB:       0,
	})

	// Verificar la conexión
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Successfully connected to Redis")
	return client
}
