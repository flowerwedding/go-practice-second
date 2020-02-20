package main

import (
	//"database/sql"
	"fmt"
	"math/rand"
	//_ "github.com/go-sql-driver/mysql"
	//"log"
)

func main(){
	//db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/student?charset=utf8"  )
	//defer db.Close()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//insertDB(db)
	var food = [6]string{"小面 6元", "饭团 7元","香菇滑鸡 12元","小炒肉 15元","黄焖鸡 16元","冒菜 18元"}
	fmt.Println(food[rand.Intn(6)],food[rand.Intn(6)])
}

//func insertDB(db *sql.DB)  {
//	stmt,err := db.Prepare("insert into test(id,name) values (2019210624,'沈怡然')")
//	if err != nil{
//		log.Fatal(err)
//	}
//	stmt.Exec()
//	fmt.Println(stmt)
//}