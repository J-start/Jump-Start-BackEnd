package news

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository/news_repository"
	"jumpStart-backEnd/service/investor_service"
)

type NewsUseCase struct {
	repo            *news_repository.NewsRepository
	investorService *investor_service.InvestorService
}

func NewNewsUseCase(repo *news_repository.NewsRepository, investorService *investor_service.InvestorService) *NewsUseCase {
	return &NewsUseCase{repo: repo, investorService: investorService}
}

func (uc *NewsUseCase) FindAllNews(offset int) ([]entities.NewsStructure, error) {
	if offset < 0 {
		return []entities.NewsStructure{}, errors.New("offset deve ser maior ou igual a 0")
	}


	newsList, err := uc.repo.FindAllNews(offset)
	if err != nil {
		return []entities.NewsStructure{}, errors.New("erro ao buscar noticias")
	}
	return newsList, nil
}

func (uc *NewsUseCase) DeleteNews(idNews int, tokenInvestor string) error {

	if idNews < 0 {
		return errors.New("idNews deve ser maior ou igual a 0")
	}

	isAdm, errAdm := uc.investorService.IsAdm(tokenInvestor)
	if errAdm != nil {
		return errors.New("erro ao verificar se o usuário é administrador")
	}

	if !isAdm {
		return errors.New("usuário não é administrador")
	}

	err := uc.repo.DeleteNews(idNews)
	if err != nil {
		return errors.New("erro ao deletar noticia")
	}
	return nil
}


