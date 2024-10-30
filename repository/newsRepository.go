package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type NewsRepository struct {
	db *sql.DB
}


type news struct {
	Id      int
	News    string
	DateNews time.Time
	IsApproved bool
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (repo *NewsRepository) FindAllNews() {
	rows, err := repo.db.Query("SELECT * FROM tb_news")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var news news
		err := rows.Scan(&news.Id, &news.News, &news.DateNews, &news.IsApproved)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(news)
	}
}