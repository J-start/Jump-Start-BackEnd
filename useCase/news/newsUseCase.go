package news

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
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
	newsList, err := uc.repo.FindAllNews(offset)
	if err != nil {
		return []entities.NewsStructure{}, errors.New("Erro ao buscar noticias")
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