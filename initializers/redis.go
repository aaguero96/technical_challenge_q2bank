package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v7"
)

var Client *redis.Client

func ConnectRedisClient() {
	var err error
	clientConfig := redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	Client = redis.NewClient(&clientConfig)

	fmt.Println(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"))

	_, err = Client.Ping().Result()
	if err != nil {
		log.Fatal("Failed to connect to redis")
	}
}
