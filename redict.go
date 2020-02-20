package main

//others


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/redict/zhihu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.zhihu.com")
	})
	router.Run(":8080")
}

//进入知乎
/*
重定向
c.Redirect(http.StatusMovedPermanently, "https://www.zhihu.com")
把这句加到c.JSON(登录成功)下面，再把c.JSON(登录成功)注释掉
否则
Headers were already written. Wanted to override status code 200 with 301
*/