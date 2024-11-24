package operation

import (
	"fmt"
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

func (uc *OperationAssetUseCase) ChangeStateOperation(idOperation int) error {
	err := uc.repo.ChangeStateOperation(idOperation)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


