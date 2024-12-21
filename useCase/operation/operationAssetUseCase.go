package operation

import (
	"database/sql"
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
)

type OperationAssetUseCase struct {
	repo *repository.OperationAssetRepository
}

func NewOperationAssetUseCase(repo *repository.OperationAssetRepository) *OperationAssetUseCase {
	return &OperationAssetUseCase{repo: repo}
}

func (uc *OperationAssetUseCase) InsertOperationAsset(datas entities.AssetInsertDataBase,repositoryService *sql.Tx) (int64, error) {
	idOperation,err := uc.repo.InsertOperationAsset(datas,repositoryService)
	if err != nil {
		return -1,err
	}
	return idOperation,nil
}

func (uc *OperationAssetUseCase) FetchAssetHistoryByInvestor(tokenUser string,offset int) ([]entities.AssetOperationHistory,error) {
	//TODO create logic to obtain ID_USER from token
	
	const ID_USER = 1
	if offset < 0 {
		return []entities.AssetOperationHistory{},errors.New("offset deve ser maior ou igual a 0")
	}
	assetHistory,err := uc.repo.FetchAssetHistoryByInvestor(ID_USER,offset)
	if err != nil {
		return nil,err
	}
	return assetHistory,nil
}




