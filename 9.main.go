package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main()  {
	// 路由组

	engine := gin.Default()

	api := engine.Group("/api")
	{
		api.GET("/get_user", func(c *gin.Context) {
			c.String(http.StatusOK, "api get_user")
		})

		api.GET("/get_info", func(c *gin.Context) {
			c.String(http.StatusOK, "api get_info")
		})
	}

	admin := engine.Group("/admin")
	{
		admin.GET("/get_user", func(c *gin.Context) {
			c.String(http.StatusOK, "admin get_user")
		})

		admin.GET("/get_info", func(c *gin.Context) {
			c.String(http.StatusOK, "admin get_info")
		})
	}


	engine.Run()
}