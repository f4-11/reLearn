package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/portmapping/go-reuse"
	"time"
)

func main() {
	go func() {
		l, err := reuse.Listen("tcp", "0.0.0.0:8080")
		if err != nil {
			panic(err)
		}
		r := gin.Default()
		r.GET("/echo", func(context *gin.Context) {
			content := context.Query("content")
			context.JSON(200, gin.H{
				"msg": fmt.Printf("response%s", content),
			})
		})
		if err := r.RunListener(l); err != nil {
			panic(err)
		}
	}()
	c, err := reuse.Dial("tcp", "0.0.0.0:8080", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	for {
		_, err := c.Write([]byte("Ping"))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
