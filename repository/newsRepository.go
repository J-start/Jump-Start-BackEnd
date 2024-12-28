package repository

import (
	"database/sql"
	"fmt"
	"jumpStart-backEnd/entities"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type NewsRepository struct {
	db *sql.DB
}


func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{db: db}
}


func (repo *NewsRepository) FindAllNews(offset int) ([]entities.NewsStructure, error) {
	offset *= 10
	query := fmt.Sprintf("SELECT id,news,datePublished FROM tb_news ORDER BY dateNews DESC LIMIT 10 OFFSET %d", offset)
	rows, err := repo.db.Query(query)
	if err != nil {
		return []entities.NewsStructure{}, err
	}
	defer rows.Close()

	var newsList []entities.NewsStructure

	for rows.Next() {
		var news entities.NewsStructure
		var dateNews time.Time
		err := rows.Scan(&news.Id,&news.News, &dateNews)
		if err != nil {
			return []entities.NewsStructure{}, err
		}
		news.DateNews = dateNews.Format("02-01-2006")
		newsList = append(newsList, news)
	}

	return newsList, nil
}


func (repo *NewsRepository) DeleteNews(idNews int) error {
	query := fmt.Sprintf(`DELETE FROM tb_news WHERE id = %d `, idNews)
	_, err := repo.db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}


func (repo *NewsRepository) GetDateLastNews() (string,error) {
	var date time.Time
	err := repo.db.QueryRow(`SELECT dateNews FROM tb_news ORDER BY dateNews DESC LIMIT 1`).Scan(&date)
	if 	err != nil {
		return "",err
	}

	return date.Format("2006-01-02"),nil

}