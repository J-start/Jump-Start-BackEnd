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

func (repo *NewsRepository) FetchNumberNewsToday() (int,error){
	var newsToday int
	err := repo.db.QueryRow(`SELECT COUNT(id) AS contagem FROM tb_news WHERE dateNews = DATE_FORMAT(NOW(),"%Y,%m,%d")`).Scan(&newsToday)
	if err != nil {
		return 0,err
	}

	return newsToday,nil
}

func (repo *NewsRepository) FindAllNews(dateQuery string) ([]entities.NewsStructure, error) {
	query := fmt.Sprintf("SELECT id,news,datePublished FROM tb_news WHERE dateNews = DATE_FORMAT('%s', '%%Y,%%m,%%d') AND isApproved = 1 LIMIT 12", dateQuery)
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