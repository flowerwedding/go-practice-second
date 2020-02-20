package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router:=gin.Default()
	router.Use(cors.Default())

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("SESSION",store))

	router.GET("/login",loginHandler)
	router.GET("/currentUser",currentUserHandler)
	router.GET("/logout",logoutHandler)

	_ = router.Run(":8080")
}

func loginHandler(c *gin.Context){
	username:=c.Query("username")
	password:=c.Query("password")

	if password != ""{
		session:=sessions.Default(c)
		session.Set("username",username)

		_ = session.Save()

		c.String(200,"登录成功")
	}else{
		c.String(401,"密码错误")
	}
}

func currentUserHandler(c *gin.Context){
	session :=sessions.Default(c)
	c.JSON(200, gin.H{"hello": session.Get("username")})
}

func logoutHandler(c *gin.Context){
	session:=sessions.Default(c)
	session.Clear()
	c.String(401,"登出")
}