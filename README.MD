# gin框架快速入门

### 前言

B站视频：[gin框架快速入门](https://www.bilibili.com/video/BV1RJ411a7iL?p=10,'gin框架快速入门')

其中 `1.main.go`对应P1



#### Goland IDE 项目创建

创建New Project

<img src="/Users/jamin/Documents/Blog/Typora-img/image-20201028104212794.png" alt="image-20201028104212794" style="zoom: 50%;" />

> 注意其中的Proxy 填入 https://goproxy.io 进行go模块的代理

​	

### 执行文件main.go创建

```go
package main	// 命名为main

// 包导入
import (
	"fmt"
)

// go run main.go 会执行main函数
func main()  {
	fmt.Println("fmt输出")
}
```

### Gin框架使用

#### 1.hello word

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main()  {
   //     hello world

   engine := gin.Default()

   engine.GET("/helloworld", func(c *gin.Context) {
      c.String(http.StatusOK, "hello world")
   })

   engine.Run()
}
```

#### 2.返回不同的数据类型

```go
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
```

#### 3.使用templates模版

```go
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
```

对应的`templates/index.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <h1>index</h1>
    <h3>name:{{.name}}</h3>
    <h3>age:{{.age}}</h3>
    <h3>users:{{.users}}</h3>
    {{range.users}}
        <p>{{.}}</p>
    {{end}}

    <img src="/static/image/10.jpg">
</body>
</html>
```



#### 4.静态文件

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main()  {
   // 静态文件
   engine := gin.Default()
   
   engine.LoadHTMLGlob("templates/*")
   engine.Static("/static", "./static")
   
   engine.GET("/index", func(context *gin.Context) {
      context.HTML(http.StatusOK, "index.html", gin.H{
         "name": "jamin",
         "age": 24,
         "users": []string{"1", "2", "3"},
      })
   })


   engine.Run()
}
```

对应的`templates/index.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="/static/css/index.css">
</head>
<body>
    <h1>static</h1>
    <img src="/static/image/10.jpg">
</body>
</html>
```

对应的`static/css/index.css`

```css
body{
    background-color: skyblue;
}
```

对应的`static/image/10.jpg`为一张图片



#### 5.请求带参数

```go
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
```





#### 6.请求参数绑定（参数结构体的序列化）

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

// 对应参数name --> Name,参数属性binding:必须，长度为6
// 对应参数age  -->  Age,参数属性binding:数字，最小为0，最大为100
type User struct {
   Name string `form:"name" binding:"required,len=6"`
   Age int `form:"age" binding:"numeric,min=0,max=100"`
}

func main()  {
   // 参数绑定

   engine := gin.Default()

   engine.GET("/get", func(context *gin.Context) {
      var user User

      err := context.ShouldBind(&user)   // 将url参数绑定到User结构体中，不满足则报错err

      if err != nil {
         context.String(http.StatusOK, err.Error())
      }else{
         context.String(http.StatusOK, "name -> %s age=%d", user.Name, user.Age)
      }

   })

   engine.Run()
}
```

#### 

#### 7.表单上传单文件

```go
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
```

`templates/upload.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="file">
        <input type="submit">
    </form>
</body>
</html>
```

#### 8.表单上传多个文件

```go
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
```

`templates/mul_upload.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="file" multiple="multiple">
        <input type="submit">
    </form>
</body>
</html>
```

#### 9.路由组

```go
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
```

> 对应URL：/api/get_user

#### 10.中间件

```go
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

   engine.Use(gin.Logger())   // gin自己的logger中间件
   //engine.Use(Test())         // 全局使用中间件


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
```

