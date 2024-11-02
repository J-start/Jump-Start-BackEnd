package usecase

import (
	"jumpStart-backEnd/repository"
	"log"
)

type ShareUseCase struct {
	repo *repository.ShareRepository
}

func NewShareUseCase(repo *repository.ShareRepository) *ShareUseCase {
	return &ShareUseCase{repo: repo}
}

func (uc *ShareUseCase) FindAllShares() ([]repository.Share, error) {

	shares, err := uc.repo.FindAllShares()
	
	if err != nil {
		log.Fatal(err)
	}	
	return shares, nil
}

func (uc *ShareUseCase) ListSharesBasedOffSet(offset int) ([]repository.Share, error) {

	shares, err := uc.repo.ListSharesBasedOffSet(offset)
	
	if err != nil {
		log.Fatal(err)
	}	
	return shares, nil
}
