package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

var (
	ctx     = context.Background()
	l2Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
)

func getDataFromL2Cache(key string) (string, error) {
	//	Check cache from Redis
	val, err := l2Cache.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}
	if errors.Is(err, redis.Nil) {
		//	Cache miss
		// Simulate DB fetch
		time.Sleep(3 * time.Second)
		val := "fresh value"
		l2Cache.Set(ctx, key, val, 5*time.Minute)

		return val, nil
	}

	return "", err
}

func main() {
	router := gin.Default()

	router.GET("/data/:key", func(c *gin.Context) {
		key := c.Param("key")
		data, err := getDataFromL2Cache(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    data,
		})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
