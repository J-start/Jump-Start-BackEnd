package operation

import (
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
)

type OperationAssetUseCase struct {
	repo *repository.OperationAssetRepository
}

func NewOperationAssetUseCase(repo *repository.OperationAssetRepository) *OperationAssetUseCase {
	return &OperationAssetUseCase{repo: repo}
}

func (uc *OperationAssetUseCase) InsertOperationAsset(datas entities.AssetInsertDataBase) (int64, error) {
	idOperation,err := uc.repo.InsertOperationAsset(datas)
	if err != nil {
		return -1,err
	}
	return idOperation,nil
}




