package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main(){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/dome7?charset=utf8"  )
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	insertDB(db)
	//updateDB(db)
	//selectDB(db)
	//deleteDB(db)
}
func insertDB(db *sql.DB)  {
	stmt,err := db.Prepare("insert into test(id,name) values (2019210624,'沈怡然')")
	if err != nil{
		log.Fatal(err)
	}
	_, _ = stmt.Exec("12222")
	fmt.Println(stmt)
}
func updateDB(db *sql.DB)  {
	stmt, err := db.Prepare("UPDATE test SET name = '花嫁' WHERE id = 2019210624")
	if err != nil{
		log.Fatal(err)
	}
	_, _ = stmt.Exec();
	fmt.Println(stmt)
}
func selectDB(db *sql.DB)  {
	stmt, err := db.Query("select * from test;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next(){
		var id int
		var name string
		//var name sql.NullString
		err := stmt.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name)
	}
fmt.Println(stmt)
}
func deleteDB(db *sql.DB)  {
	stmt, err := db.Prepare("delete from test where id =20192210624")
	if err != nil{
		log.Fatal(err);
	}
	_, _ = stmt.Exec();
    fmt.Println(stmt)
}
