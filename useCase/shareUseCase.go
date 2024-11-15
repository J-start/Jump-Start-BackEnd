package usecase

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"log"
	"strings"
)

type ShareUseCase struct {
	repo *repository.ShareRepository
}

func NewShareUseCase(repo *repository.ShareRepository) *ShareUseCase {
	return &ShareUseCase{repo: repo}
}

func (uc *ShareUseCase) FindAllShares() ([]entities.Share, error) {

	shares, err := uc.repo.FindAllShares()
	
	if err != nil {
		log.Fatal(err)
	}	
	return shares, nil
}

func (uc *ShareUseCase) ListSharesBasedOffSet(offset int) ([]entities.Share, error) {

	shares, err := uc.repo.ListSharesBasedOffSet(offset)
	
	if err != nil {
		log.Fatal(err)
	}	
	return shares, nil
}

func (uc *ShareUseCase) FindShareById(shareName string) (entities.Share, error) {

	if shareName == "" || len(strings.Split(shareName, " ")) ==0 || len(shareName) == 0 {
		return entities.Share{}, nil
	}

		isContainName, err := uc.ShareNameList(shareName)

		if err != nil {
			return entities.Share{}, err
		}
		if !isContainName {
			return entities.Share{}, errors.New("ação não encontrada")
		}

		share, err := uc.repo.FindShareById(shareName)

		if	err != nil {
			return entities.Share{}, err
		}

		return share, nil
	
}

func (uc *ShareUseCase) ShareList(shareName string) ([]entities.Share, error) {
	
	isContainName, err := uc.ShareNameList(shareName)

	if err != nil {
		return []entities.Share{}, err
	}
	if !isContainName {
		return []entities.Share{}, errors.New("ação não encontrada")
	}

	shareList, err := uc.repo.ShareList(shareName)

	if err != nil {
		return []entities.Share{}, err
	}

	return shareList,nil
}


func (uc *ShareUseCase) ShareNameList(shareName string) (bool, error) {
	
	shareNames,err := uc.repo.ShareNameList()
	if err != nil {
		return false,err
	}

	for _, name := range shareNames {
		if name.NameShare == shareName {
			return true, nil	
			
		}
	}

	return false,nil
}
