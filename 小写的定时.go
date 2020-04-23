package main

import (
	"github.com/robfig/cron"
	"log"
)

/*
func main() {
	go func() {
		crontab := cron.New()//不支持秒
		_, _ = crontab.AddFunc("0 20 8 * * *",myfunc) // 每天 8:20:00 定时执行 myfunc 函数
		crontab.Start()
	}()
}
*/

func main() {
	i := 0
	c := cron.NewWithSeconds()
	spec := "*/5 * * * * ?"
	_, _ = c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	select{}
}