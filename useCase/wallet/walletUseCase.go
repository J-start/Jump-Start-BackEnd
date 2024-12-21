package wallet

import (
	"database/sql"
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/useCase/operation"
)

type WalletUseCase struct {
	repo                  *repository.WalletRepository
	operationAssetUseCase *operation.OperationAssetUseCase
}

func NewWalletUseCase(repo *repository.WalletRepository, operationAssetUseCase *operation.OperationAssetUseCase) *WalletUseCase {
	return &WalletUseCase{repo: repo, operationAssetUseCase: operationAssetUseCase}
}

func (uc *WalletUseCase) InsertValueBalance(idInvestor int, value float64, idOperation int64, repositoryService *sql.Tx) error {
	balance, err := uc.isInvestorValid(idInvestor)
	if err != nil {
		return errors.New("erro ao verificar saldo do usuário")
	}
	balance += value

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

