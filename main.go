package main

import (
	"database/sql"
	"ferryapp/cmd"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn, err := cmd.Setup()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado a la base de datos")

}
