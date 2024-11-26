package repository

import (
	"database/sql"
	"jumpStart-backEnd/entities"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type NewsRepository struct {
	db *sql.DB
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
		var news entities.News
		err := rows.Scan(&news.Id, &news.News, &news.DateNews, &news.IsApproved)
		if err != nil {
			log.Fatal(err)
		}
	}
}