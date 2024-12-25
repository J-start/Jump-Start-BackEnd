package news

import (
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"time"
)

type NewsUseCase struct {
	repo *repository.NewsRepository
}

func NewNewsUseCase(repo *repository.NewsRepository) *NewsUseCase {
	return &NewsUseCase{repo: repo}
}

func (uc *NewsUseCase) FindAllNews(offset int) ([]entities.NewsStructure, error) {
	if offset < 0 {
		return []entities.NewsStructure{}, errors.New("offset deve ser maior ou igual a 0")
	}

	numberNewsAvailable, errNumberNews := uc.repo.FetchNumberNewsToday()
	if errNumberNews != nil {
		return []entities.NewsStructure{}, errNumberNews
	}
	dateQuery := createDateQuery(0)
	if numberNewsAvailable == 0 {
		dateQuery = createDateQuery(-24)
	}

	if numberNewsAvailable == 0 && offset > 0 {
		dateQuery = createDateQuery(-24 * ( offset + 1 ))
	}else if numberNewsAvailable > 0 && offset > 0 {
		dateQuery = createDateQuery(-24 * offset)
	}

	newsList, err := uc.repo.FindAllNews(dateQuery)
	if err != nil {
		fmt.Println(err)
		return []entities.NewsStructure{}, errors.New("erro ao buscar noticias")
	}
	return newsList, nil
}

func (uc *NewsUseCase) DeleteNews(idNews int,tokenInvestor string) error {

	//TODO CREATE LOGIC TO VERIFY IF USER IS ADMIN TO DELETE NEWS

	if idNews < 0 {
		return errors.New("idNews deve ser maior ou igual a 0")
	}
	err := uc.repo.DeleteNews(idNews)
	if err != nil {
		return errors.New("erro ao deletar noticia")
	}
	return nil
}

func createDateQuery(rangeNews int) string {
	currentDate := time.Now()	
	yesterday := currentDate.Add(time.Duration(rangeNews) * time.Hour)
	return yesterday.Format("2006-01-02")
}