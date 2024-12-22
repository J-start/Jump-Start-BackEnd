package wallet

import (
	"database/sql"
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/serviceRepository"
	"jumpStart-backEnd/useCase/operation"
	"net/http"
	"time"
)

type WalletUseCase struct {
	repo                  *repository.WalletRepository
	operationAssetUseCase *operation.OperationAssetUseCase
	repositoryService 	  *servicerepository.ServiceRepository
}

func NewWalletUseCase(repo *repository.WalletRepository, operationAssetUseCase *operation.OperationAssetUseCase,repositoryService *servicerepository.ServiceRepository) *WalletUseCase {
	return &WalletUseCase{repo: repo, operationAssetUseCase: operationAssetUseCase,repositoryService:repositoryService}
}

func (uc *WalletUseCase) InsertValueBalance(idInvestor int, value float64, idOperation int64, repositoryService *sql.Tx) error {
	balance, err := uc.isInvestorValid(idInvestor)
	if err != nil {
		return errors.New("erro ao verificar saldo do usuário")
	}
	balance += value

	if balance < 0 {
		return errors.New("saldo invalido")
	}

	errInsert := uc.repo.UpdateBalanceInvestor(idInvestor, balance, idOperation, repositoryService)

	if errInsert != nil {
		return errors.New("erro ao atualizar o saldo")
	}

	return nil

}

func (uc *WalletUseCase) isInvestorValid(idInvestor int) (float64, error) {
	balance, err := uc.repo.IsBalanceExists(idInvestor)
	if err != nil {
		if err.Error() == "nenhum dado encontrado" {
			err := uc.repo.CreateBalanceUser(idInvestor)
			if err != nil {
				return 0, errors.New("erro ao criar saldo para o usuário")
			}
		} else {
			return 0, errors.New("ocorreu um erro ao verificar o saldo do usuário")
		}
		return 0, err
	}
	return balance, nil
}

func (uc *WalletUseCase) VerifyIfInvestorCanOperate(id int, value float64) error {

	balance, err := uc.repo.FindBalanceInvestor(id)

	if err != nil {
		return err
	}

	if balance == 0 {

		return errors.New("saldo invalido")
	}

	if balance < value {

		return errors.New("saldo invalido")
	}

	return nil
}

func (uc *WalletUseCase) FindIdWallet(id int) (int, error) {
	wallet, err := uc.repo.FindIdWallet(id)
	if err != nil {
		return 0, err
	}
	return wallet, nil
}

func (uc *WalletUseCase) FetchDatasWalletInvestor(tokenInvestor string) (entities.WalletDatas, error) {
	// TODO: CREATE LOGIC TO VALIDATE AND RECOVER ID INVESTOR
	const ID_INVESTOR = 1
	balanceChan := make(chan struct {
		result float64
		err    error
	})
	assetsChan := make(chan struct {
		result []entities.Asset
		err    error
	})

	go func() {
		balance, err := uc.repo.FindBalanceInvestor(ID_INVESTOR)
		balanceChan <- struct {
			result float64
			err    error
		}{result: balance, err: err}
	}()

	go func() {
		assets, err := uc.repo.SearchDatasWallet(ID_INVESTOR)
		assetsChan <- struct {
			result []entities.Asset
			err    error
		}{result: assets, err: err}
	}()

	balanceResult := <-balanceChan
	assetsResult := <-assetsChan

	close(balanceChan)
	close(assetsChan)

	if balanceResult.err != nil {
		return entities.WalletDatas{}, balanceResult.err
	}

	if assetsResult.err != nil {
		return entities.WalletDatas{}, assetsResult.err
	}

	WalletDatas := entities.WalletDatas{
		InvestorBalance: balanceResult.result,
		Assets:          assetsResult.result,
	}

	return WalletDatas, nil
}

func (uc *WalletUseCase) FetchOperationsWallet(tokenInvestor string, offset int) ([]entities.WalletOperation, error) {

	//TODO CREATE LOGIC TO OBTAIN ID_USER FROM TOKEN
	const ID_USER = 1

	if offset < 0 {
		return []entities.WalletOperation{}, errors.New("offset deve ser maior ou igual a 0")
	}

	operations, err := uc.repo.FetchDatasWalletOperation(ID_USER, offset)
	if err != nil {
		return nil, err
	}
	return operations, nil

}

func (uc *WalletUseCase) WithDraw(operation entities.WalletOperationRequest) (int, string) {
	// TODO CREATE LOGIC TO VALIDATE AND RECOVER ID INVESTOR
	const ID_INVESTOR = 1
	balance , err := uc.isInvestorValid(ID_INVESTOR)
	if err != nil {
		return http.StatusBadRequest, "dados do usuário inválidos"
	}

	repositoryService,errRepository := uc.repositoryService.StartTransaction()

	if errRepository != nil {
		return http.StatusBadRequest, "erro ao processar requisição, tente novamente"
	}

	if balance < operation.Value {
		return http.StatusNotAcceptable, "saldo insuficiente"
	}
	balance -= operation.Value

	errUpdateBalance := uc.repo.UpdateBalanceFromOperation(ID_INVESTOR, balance,repositoryService)
	if errUpdateBalance != nil {
		repositoryService.Rollback()
		return http.StatusBadRequest, "erro ao atualizar saldo"
	}

	dateMysql := convertDateToMysql()
	errInsert := uc.repo.InsertOperationWallet("WITHDRAW", operation.Value,dateMysql, ID_INVESTOR, repositoryService)
	if errInsert != nil {
		repositoryService.Rollback()
		return http.StatusBadRequest, "erro ao concluir operação, tente novamente"
	}

	errService := repositoryService.Commit()
	if errService != nil {
		repositoryService.Rollback()
		return http.StatusBadRequest, "erro ao processar requisição, tente novamente"
	}


	return http.StatusOK, "operação realizada com sucesso"
}

func (uc *WalletUseCase) Deposit(operation entities.WalletOperationRequest) (int, string) {
	// TODO CREATE LOGIC TO VALIDATE AND RECOVER ID INVESTOR
	const ID_INVESTOR = 1
	balance , err := uc.isInvestorValid(ID_INVESTOR)
	if err != nil {
		return http.StatusBadRequest, "dados do usuário inválidos"
	}

	balanceDay , errBalanceDay := uc.repo.FetchDayDeposits(ID_INVESTOR)
	if errBalanceDay != nil {
		return http.StatusBadRequest, "erro ao processar requisição, tente novamente"
	}
	if  (balanceDay + operation.Value) > 1000 {
		return http.StatusNotAcceptable, "limite diário de depósito atingido"
	}

	repositoryService,errRepository := uc.repositoryService.StartTransaction()

	if errRepository != nil {
		return http.StatusBadRequest, "erro ao processar requisição, tente novamente"
	}

	if operation.Value <= 0 {
		return http.StatusNotAcceptable, "saldo insuficiente"
	}
	balance += operation.Value

	errUpdateBalance := uc.repo.UpdateBalanceFromOperation(ID_INVESTOR, balance,repositoryService)
	if errUpdateBalance != nil {
		repositoryService.Rollback()
		return http.StatusBadRequest, "erro ao atualizar saldo"
	}

	dateMysql := convertDateToMysql()
	errInsert := uc.repo.InsertOperationWallet("DEPOSIT", operation.Value,dateMysql, ID_INVESTOR, repositoryService)
	if errInsert != nil {
		repositoryService.Rollback()
		return http.StatusBadRequest, "erro ao concluir operação, tente novamente"
	}

	errService := repositoryService.Commit()
	if errService != nil {
		repositoryService.Rollback()
		return http.StatusBadRequest, "erro ao processar requisição, tente novamente"
	}


	return http.StatusOK, "operação realizada com sucesso"
}

func convertDateToMysql() string{
	currentDate := time.Now()
	dateMysql := currentDate.Format("2006-01-02")
	return dateMysql
}
