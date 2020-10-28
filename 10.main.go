package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test()  gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("当前URL:", c.Request.URL)
	}
}


func main()  {
	// 中间件
	//engine := gin.Default()
	engine := gin.New()

	engine.Use(gin.Logger())	// gin自己的logger中间件
	//engine.Use(Test())			// 全局使用中间件


	engine.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "index")
	})


	// 指定某个url使用中间件
	engine.GET("/handlers", Test(), func(c *gin.Context) {
		c.String(http.StatusOK, "handlers")
	})

	// 指定某个路由组使用中间件
	api := engine.Group("/api", Test())
	{
		api.GET("get_user", func(c *gin.Context) {
			c.String(http.StatusOK, "get_user")
		})
	}

	engine.Run()
}