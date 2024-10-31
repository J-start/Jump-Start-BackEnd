package usecase

import (
	"jumpStart-backEnd/dto"
	"jumpStart-backEnd/repository"
	"log"
)

type ShareUseCase struct {
	repo *repository.ShareRepository
}

func NewShareUseCase(repo *repository.ShareRepository) *ShareUseCase {
	return &ShareUseCase{repo: repo}
}

func (uc *ShareUseCase) FindAllShares() ([]dto.ShareDTO, error) {

	shares, err := uc.repo.FindAllShares()
	
	if err != nil {
		log.Fatal(err)
	}	
	return shares, nil
}
