package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var (
	l1Cache = cache.New(1*time.Minute, 2*time.Minute)
)

func getData(key string) (string, error) {
	// Check L1 Cache
	if val, ok := l1Cache.Get(key); ok {
		return val.(string), nil
	}

	// Simulate DB fetch
	time.Sleep(3 * time.Second)
	val := "some data"

	l1Cache.Set(key, val, cache.DefaultExpiration)

	return val, nil
}

func main() {
	router := gin.Default()

	router.GET("/data/:key", func(c *gin.Context) {
		key := c.Param("key")

		result, err := getData(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    err,
			})

			return
		}

		fmt.Println("result:", result)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    result,
		})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
