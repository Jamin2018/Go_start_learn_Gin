package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	// 	hello world

	engine := gin.Default()

	engine.GET("/helloworld", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	engine.Run()
}