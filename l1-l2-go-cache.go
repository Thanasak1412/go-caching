package main

import (
	"fmt"
	"net/http"
	"time"

	"context"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

var (
	ctx     = context.Background()
	l1Cache = cache.New(1*time.Minute, 2*time.Minute)
	l2Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
)

func getDataFromCache(key string) (string, error) {
	//	L1
	if val, ok := l1Cache.Get(key); ok {
		return val.(string), nil
	}

	//	L2
	val, err := l2Cache.Get(ctx, key).Result()
	if err == nil {
		l1Cache.Set(key, val, cache.DefaultExpiration)

		return val, nil
	}

	//	Cache Miss (Simulate DB/API)
	time.Sleep(3 * time.Second)
	val = fmt.Sprintf("fresh value at %v", time.Now())

	// Store in both
	l1Cache.Set(key, val, cache.DefaultExpiration)
	l2Cache.Set(ctx, key, val, 5*time.Minute)

	return val, nil
}

func main() {
	router := gin.Default()

	router.GET("/data/:key", func(c *gin.Context) {
		key := c.Param("key")

		data, err := getDataFromCache(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    data + "",
		})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
