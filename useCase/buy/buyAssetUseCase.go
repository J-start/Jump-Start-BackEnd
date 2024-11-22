package buy

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/useCase"
	"strings"
	"time"
	"net/http"
)

type BuyAssetUseCase struct {
	repo                     *repository.ShareRepository
	shareUseCase             *usecase.ShareUseCase
	walletRepository         *repository.WalletRepository
	operationAssetRepository *repository.OperationAssetRepository
}

func NewBuyAssetsUseCase(repo *repository.ShareRepository, shareUseCase *usecase.ShareUseCase, walletRepository *repository.WalletRepository, operationAssetRepository *repository.OperationAssetRepository) *BuyAssetUseCase {
	return &BuyAssetUseCase{repo: repo, shareUseCase: shareUseCase, walletRepository: walletRepository, operationAssetRepository: operationAssetRepository}
}

func (uc *BuyAssetUseCase) BuyAsset(assetOperation entities.AssetOperation) (int, string) {

	err := ValidateFields(assetOperation)
	if err != nil {
		return http.StatusNotAcceptable, strings.ToUpper(string(err.Error()[0])) + err.Error()[1:]
	}

	var value float64

	if assetOperation.AssetType != "SHARE" {

		response, err := MakeRequestAsset(assetOperation.AssetType, assetOperation.AssetCode)
		if err != nil {
			return http.StatusInternalServerError, "Algum campo está inválido"
		}

		if assetOperation.AssetType == "COIN" {
			valueReturn, err := getValueFromCoin(response, assetOperation.AssetCode)
			if err != nil {
				return http.StatusInternalServerError, "Ocorreu um erro ao buscar o valor da moeda, tente novamente"
			}
			value = valueReturn
		} else if assetOperation.AssetType == "CRYPTO" {
			valueReturn, err := getValueFromCrypto(response)
			if err != nil {
				return http.StatusInternalServerError, "Ocorreu um erro ao buscar o valor da cryptomoeda, tente novamente"
			}
			value = valueReturn
		}
	} else {
		if !isActionTradable(time.Now()) {
			return http.StatusNotAcceptable, "O mercado está fechado.Não é possível comprar ou vender ações"
		}

		err := uc.isAssetValid(assetOperation.AssetCode)
		if err != nil {
			return http.StatusNotAcceptable, "Ação inválida"
		}

		valueReturn, err := uc.getValueFromShare(assetOperation.AssetCode)
		if err != nil {
			return http.StatusNotAcceptable, "Problema ao consultar o valor da ação, tente novamente"
		}
		value = valueReturn
	}
	datasToInsert := buildDatasToInsert(assetOperation, value, 1)
	valueOperation := datasToInsert.AssetAmount * datasToInsert.AssetValue
	
	errBuy := uc.verifyIfInvestorCanBuy(1, valueOperation)
	
	datasToInsert.AssetValue = valueOperation

	if errBuy != nil {
		return http.StatusNotAcceptable, "Saldo insuficiente"
	}

	err = uc.operationAssetRepository.InsertOperationAsset(datasToInsert)

	if err != nil {
		return http.StatusInternalServerError, "Ocorreu algum erro quando tentamos concluir a operação, tente novamente"
	}

	return 200,"Operação realizada com sucesso"
}

func (uc *BuyAssetUseCase) getValueFromShare(code string) (float64, error) {
	share, err := uc.repo.FindShareById(code)
	if err != nil {
		return 0, err
	}
	return share.CloseShare, nil
}

func (uc *BuyAssetUseCase) isAssetValid(code string) error {
	if code == "" || len(strings.Split(code, " ")) == 0 || len(code) == 0 {
		return errors.New("código de ativo inválido")
	}

	isContainName, err := uc.shareUseCase.ShareNameList(code)
	if err != nil {
		return err
	}
	if !isContainName {
		return errors.New("nome de ação inválida")
	}

	return nil
}

func (uc *BuyAssetUseCase) verifyIfInvestorCanBuy(id int, value float64) error {

	balance, err := uc.walletRepository.FindBalanceInvestor(id)

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
