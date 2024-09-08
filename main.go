package main

import (
	"database/sql"
	"ferryapp/cmd"
	"fmt"
	"log"

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
	defer db.Close()

	fmt.Println("Conectado a la base de datos")

	re, err := db.Exec("CALL InsertarCategoria(?, ?)", "categoria 7", true)
	if err != nil {
		log.Printf("error insertando: %v", err)
		return
	}
	fmt.Println(re)
}
