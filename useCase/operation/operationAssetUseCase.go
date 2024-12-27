package operation

import (
	"database/sql"
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/service/investor_service"
)

type OperationAssetUseCase struct {
	repo *repository.OperationAssetRepository
	investorService *investor_service.InvestorService
}

func NewOperationAssetUseCase(repo *repository.OperationAssetRepository,investorService *investor_service.InvestorService) *OperationAssetUseCase {
	return &OperationAssetUseCase{repo: repo, investorService: investorService}
}

func (uc *OperationAssetUseCase) InsertOperationAsset(datas entities.AssetInsertDataBase,repositoryService *sql.Tx) (int64, error) {
	idOperation,err := uc.repo.InsertOperationAsset(datas,repositoryService)
	if err != nil {
		return -1,err
	}
	return idOperation,nil
}

func (uc *OperationAssetUseCase) FetchAssetHistoryByInvestor(tokenInvestor string,offset int) ([]entities.AssetOperationHistory,error) {

    //  idInvestor,err := uc.investorService.GetIdByToken(tokenInvestor)
	//  if err != nil {
	//  	return []entities.AssetOperationHistory{}, errors.New("token inválido, realize o login novamente")
	//  }
	const ID_USER = 2
	if offset < 0 {
		return []entities.AssetOperationHistory{},errors.New("offset deve ser maior ou igual a 0")
	}
	assetHistory,err := uc.repo.FetchAssetHistoryByInvestor(ID_USER,offset)
	if err != nil {
		return nil,err
	}
	return assetHistory,nil
}




