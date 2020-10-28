package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `form:"name" binding:"required,len=6"`
	Age int `form:"age" binding:"numeric,min=0,max=100"`
}

func main()  {
	// 参数绑定

	engine := gin.Default()

	engine.GET("/get", func(context *gin.Context) {
		var user User

		err := context.ShouldBind(&user)	// 将url参数绑定到User结构体中，不满足则报错err

		if err != nil {
			context.String(http.StatusOK, err.Error())
		}else{
			context.String(http.StatusOK, "name -> %s age=%d", user.Name, user.Age)
		}

	})

	engine.Run()
}