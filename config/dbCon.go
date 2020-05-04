package config

import (
	"database/sql"
	"fmt"
	_ "github.com/sudesh35139/prx/go-sql-driver/mysql"
)

var DB *sql.DB

func init()  {

	var err error
	DB,err = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/proximity?charset=utf8")
	if err != nil{
		panic(err)
	}
	if err = DB.Ping();err!=nil{
		panic(err)
	}
	fmt.Println("localhost db connected")
}