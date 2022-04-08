package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Init() *sql.DB {
	db, err := sql.Open("mysql", "root:Lmf950122.@tcp(127.0.0.1:3306)/takeHomeAssignment")
	checkErr(err)

	err = db.Ping()
	checkErr(err)
	fmt.Println("Connect to database successfully")

	return db
}
