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

func (uc *OperationAssetUseCase) FetchAssetHistoryByInvestor(idInvestor int,offset int) ([]entities.AssetOperationHistory,error) {
	//TODO create logic to validate id investor
	
	if offset < 0 {
		return []entities.AssetOperationHistory{},errors.New("offset deve ser maior que 0")
	}
	assetHistory,err := uc.repo.FetchAssetHistoryByInvestor(idInvestor,offset)
	if err != nil {
		return nil,err
	}
	return assetHistory,nil
}




