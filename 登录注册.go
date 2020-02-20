package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {//主体部分
	router := gin.Default()
	router.GET("/LoginRegistration", Login)
	router.POST("/LoginRegistration", Registration)
	_ = router.Run(":8080")
}

func Login (c *gin.Context) {//注册
	name := c.DefaultQuery("name", "Guest")
	password := c.Query("password")
	Bool := IsExist(name)
	if Bool {
		State["text"] = "此用户已存在！"
	} else {
		AddStruct(name, password)
		State["text"] = "注册成功！"
		cookie := &http.Cookie{
			Name:     name,
			Value:password,
			Path:      "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
	}
	c.String(http.StatusOK, "%v", State["text"])
}

func Registration (c *gin.Context) {//登录
	name := c.DefaultQuery("name", "Guest")
	password := c.Query("password")
	Bool:=IsExist(name)
	i:=0
	if Bool {
		Bool_Pwd :=IsRight(name, password)
		if Bool_Pwd {
			State["text"] = "登录成功！\n"
			if cookie, err := c.Request.Cookie(name); err == nil {
				value := cookie.Value
				if value == password {
					c.Next()
					i=1
				}
			}
		} else {
			State["text"] = "登录失败！密码错误！\n"
		}
	} else {
		State["text"] = "登录失败！此用户尚未注册！\n"
	}
	c.String(http.StatusOK, "%v", State["text"])
	cookiereremamber(i,name,c)
}

type User struct {
	Name   string
	Password string
}

var Slice []User

var State = make(map[string]interface{})

func AddStruct(name string, password string) {
	var user User
	user.Name = name
	user.Password = password
	Slice = append(Slice, user)
}

func IsExist(name string) bool {//判断用户是否存在，之前有无注册过
	if len(Slice) == 0 {
		return false
	}else {
		for _, v := range Slice {
			if v.Name == name {
				return true
			}
		}
	}
	return false
}

func IsRight(name string, password string) bool {//判断密码是否正确
	for _, v := range Slice {
		if v.Name == name {
			return v.Password == password
		}
	}
	return false
}

func cookiereremamber(i int,name string,c *gin.Context){
	var named string
	if i==1 {
		named=name
	}else {
		named="guest"
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":200,
		"message": "hello"+named,
	})
}