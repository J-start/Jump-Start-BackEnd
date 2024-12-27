package news

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/service/investor_service"
	"time"
)

type NewsUseCase struct {
	repo            *repository.NewsRepository
	investorService *investor_service.InvestorService
}

func NewNewsUseCase(repo *repository.NewsRepository, investorService *investor_service.InvestorService) *NewsUseCase {
	return &NewsUseCase{repo: repo, investorService: investorService}
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
		dateQuery = createDateQuery(-24 * (offset + 1))
	} else if numberNewsAvailable > 0 && offset > 0 {
		dateQuery = createDateQuery(-24 * offset)
	}

	newsList, err := uc.repo.FindAllNews(dateQuery)
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

func createDateQuery(rangeNews int) string {
	currentDate := time.Now()
	yesterday := currentDate.Add(time.Duration(rangeNews) * time.Hour)
	return yesterday.Format("2006-01-02")
}