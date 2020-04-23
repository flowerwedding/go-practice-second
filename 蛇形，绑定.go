package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Loginll
		if err := c.ShouldBindJSON(&json); err != nil {
			//fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Return User":json.User,"Return Password":json.Password})
	})

	_ = router.Run(":8080")
}

type Loginll struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

//后面的form:user表示在form中这个字段是user,不是User, 同样json:user也是
//注意:binding:"required"要求这个字段在client端发送的时候必须存在,否则报错!