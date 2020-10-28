package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	// templates模版
	engine := gin.Default()
	
	engine.LoadHTMLGlob("templates/*")
	
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"name": "jamin",
			"age": 24,
			"users": []string{"1", "2", "3"},
		})
	})


	engine.Run()
}