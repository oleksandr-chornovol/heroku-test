package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(5).Seconds().Do(func() {
		fmt.Println("Hello from scheduler!")
	})
	if err != nil {
		log.Println(err)
	}

	s.StartAsync()

	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	err = engine.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
	}
}
