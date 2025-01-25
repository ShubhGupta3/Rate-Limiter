package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var bucket int

func main() {
	fmt.Println("Namaste!")

	bucket = 5

	client := RedisInit()
	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("foo: ", val)

	go RefillBucket()
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	rateLimited := router.Group("/")

	rateLimited.Use(RateLimiter())
	router.GET("/unlimited", unlimited)

	rateLimited.GET("/limited", limited)

	router.Run("localhost:8080")

}

func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	return client
}

func unlimited(c *gin.Context) {
	fmt.Println("Inside unlimited...")
	c.JSON(200, "Unlimited! Let's GO!")
}

func limited(c *gin.Context) {
	fmt.Println("Inside limited...")
	c.JSON(200, "Limited, do not overuse me!")
}

func RateLimiter() gin.HandlerFunc {
	fmt.Println("Hello ji, middleware entered")
	return TokenBucket
}

func RefillBucket() {
	for {
		time.Sleep(60 * time.Second)
		bucket = 5
	}
}

// STEP 1: DONE -> Create a simple API with 2 endpoints /unlimited and /limited
