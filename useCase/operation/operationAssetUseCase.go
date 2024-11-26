package operation

import (
	"database/sql"
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




