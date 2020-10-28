package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main()  {
	// 表单上传多个文件

	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*")

	engine.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "mul_upload.html", nil)
	})

	engine.POST("/upload", func(context *gin.Context) {
		form, _ := context.MultipartForm()

		files := form.File["file"]

		for _, file := range files{
			context.SaveUploadedFile(file, fmt.Sprintf("uploads_file/%s", file.Filename))
		}

		context.String(http.StatusOK, "ok")

	})

	engine.Run()
}