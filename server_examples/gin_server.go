package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"time"
)

func EventHandler(c *gin.Context) {
	chanStream := make(chan int, 11)
	go func() {
		defer close(chanStream)
		for i := 0; i < 11; i++ {
			chanStream <- i
			time.Sleep(time.Second * 1)
		}
	}()
	c.Header("Access-Control-Allow-Origin", "*")
	c.Stream(func(w io.Writer) bool {
		if progress, ok := <-chanStream; ok {
			c.SSEvent("sse-message", map[string]int{"progress": progress})
			return true
		}
		return false
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/api/event", EventHandler)

	server := &http.Server{
		Addr:         ":9999",
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: 0,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server listen: %s\n", err)
	}
}
