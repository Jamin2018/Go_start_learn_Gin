package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	// 请求返回的数据类型

	engine := gin.Default()

	engine.GET("/helloworld", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	engine.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "admin",
			"age": 24,
		})
	})

	engine.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"name": "admin",
			"age": 24,
		})
	})

	engine.Run()
}