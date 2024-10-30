package db

import (
	"database/sql"
	"log"
	"time"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {

	dsn := "root:teste_senha@tcp(127.0.0.1:3306)/teste_banco?parseTime=true"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if err := db.Ping(); err == nil {
			fmt.Println("Conexão ao MySQL bem-sucedida!")
			return
		}
		fmt.Println("Tentando conectar ao MySQL...")
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Não foi possível conectar ao MySQL após várias tentativas")
}

func GetDB() *sql.DB {
	return db
}