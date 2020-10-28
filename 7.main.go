package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func main()  {
	// 表单上传单文件

	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*")

	engine.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})

	engine.POST("/upload", func(context *gin.Context) {
		f, err := context.FormFile("file")
		if err != nil{
			log.Println(err)
		}else{
			err := context.SaveUploadedFile(f, fmt.Sprintf("uploads_file/%s", f.Filename))
			if err != nil{
				context.String(http.StatusOK, "上传文件失败：%v", err.Error())

			}else {
				context.String(http.StatusOK, "上传成功")
			}
		}
	})

	engine.Run()
}