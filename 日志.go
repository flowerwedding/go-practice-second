package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
	"runtime"
)

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin-rizhi.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s\t|\t%s|\t%s|\t%s|\t%s|\t%d|\t%s|\t%s|\t%s \n",
			param.ClientIP,//客户端ip
			param.TimeStamp.Format("2006-01-02 15:04:05"),//时间格式
			param.Method,//http请求 get post
			param.Path,//客户端请求的路径
			param.Request.Proto,//http请求协议版本
			param.StatusCode,//http请求状态码
			param.Latency,//耗时
			param.Request.UserAgent(),//http请求代理头
			param.ErrorMessage,//处理请求错误时设置错误消息
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.POST("/denglu",denglu)

	_ = router.Run(":8080")
}

func denglu(c *gin.Context) {
	user_id := c.PostForm("user_id")
	password := c.PostForm("password")
	if userEmpty(user_id) {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "请输入手机号或邮箱"})
	} else if userSignin(user_id, password) {
		_, _ = fmt.Fprintln(gin.DefaultWriter, user_id,"登录")
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "账号或密码错误"})
	}
}

func userEmpty(username string)bool {
	if username == "" {return true}
	return false
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dome7?charset=utf8mb4")
	if err != nil {
		_, _ = fmt.Fprintln(gin.DefaultWriter, "err；",err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	err = db.Ping()
	if err != nil {
		_, _ = fmt.Fprintln(gin.DefaultWriter, "err；",err)
	}
}

func Conn() *sql.DB {
	return db
}


func userSignin(user_id string,password string) bool {
	db := Conn()
	var passwd string
	err := db.QueryRow(`select password from test where user_id = ?`,user_id).Scan(&passwd)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		if _, file, line, ok := runtime.Caller(0); ok {
			_, _ = fmt.Fprintln(gin.DefaultWriter, "err；",err,"file：", file,"line：", line)
		}
	}
	if  passwd==password{ return true}
	return false
}