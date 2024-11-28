package db

import (
	"database/sql"
	"log"
	"time"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {

	err2 := godotenv.Load()
    if err2 != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
    }

	PASSWORD := os.Getenv("MYSQL_ROOT_PASSWORD")
	DATABASE := os.Getenv("MYSQL_DATABASE")
	USER := os.Getenv("MYSQL_USER")
	HOST := os.Getenv("MYSQL_HOST")
	PORT := os.Getenv("MYSQL_PORT")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", USER, PASSWORD,HOST,PORT,DATABASE)
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