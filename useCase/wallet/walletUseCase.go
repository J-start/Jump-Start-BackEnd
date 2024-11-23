package wallet

import (
	"errors"
	"jumpStart-backEnd/repository"
)

type WalletUseCase struct {
	repo *repository.WalletRepository
}

func NewWalletUseCase(repo *repository.WalletRepository) *WalletUseCase {
	return &WalletUseCase{repo: repo}
}

func (uc *WalletUseCase) InsertValueBalance(idInvestor int,value float64) error {
	balance,err := uc.isInvestorValid(idInvestor)
	if err != nil {
		return errors.New("erro ao verificar saldo do usuário")
	}
	balance += value

	errInsert := uc.repo.UpdateBalanceInvestor(idInvestor, balance)

	if errInsert != nil {
		return errors.New("erro ao atualizar o saldo")
	}

	return nil
}

func (uc *WalletUseCase) isInvestorValid(idInvestor int) (float64,error) {
	balance,err := uc.repo.IsBalanceExists(idInvestor)
	if err != nil {
		if err.Error() == "nenhum dado encontrado" {
			err := uc.repo.CreateBalanceUser(idInvestor)
			if err != nil {
				return 0,errors.New("erro ao criar saldo para o usuário")
			}
		} else {
			return 0,errors.New("ocorreu um erro ao verificar o saldo do usuário")
		}
		return 0,err
	}
	return balance,nil
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

