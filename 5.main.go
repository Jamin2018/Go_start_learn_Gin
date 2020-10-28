package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	// 请求带参数

	engine := gin.Default()
	
	engine.GET("/get", func(context *gin.Context) {
		fmt.Println(context.Query("name"))
		fmt.Println(context.Query("age"))
		
		context.String(http.StatusOK, "get")
	})
	
	engine.POST("/post", func(context *gin.Context) {
		fmt.Println(context.Query("name"))
		fmt.Println(context.Query("age"))

		context.String(http.StatusOK, "post")
	})

	engine.GET("/get_user/:id", func(context *gin.Context) {
		fmt.Println(context.Param("id"))

		context.String(http.StatusOK, "url Param")
	})
	
	engine.Run()
}