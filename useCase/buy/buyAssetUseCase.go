package buy

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/useCase"
	"jumpStart-backEnd/useCase/assetwallet"
	"jumpStart-backEnd/useCase/operation"
	"jumpStart-backEnd/useCase/wallet"
	"net/http"
	"strings"
	"time"
)

type BuyAssetUseCase struct {
	repo                     *repository.ShareRepository
	shareUseCase             *usecase.ShareUseCase
	walletUseCase            *wallet.WalletUseCase
	operationAssetUseCase    *operation.OperationAssetUseCase
	assetWalletUseCase       *assetwallet.AssetWalletUseCase
}

func NewBuyAssetsUseCase(repo *repository.ShareRepository, shareUseCase *usecase.ShareUseCase, walletUseCase *wallet.WalletUseCase, operationAssetUseCase *operation.OperationAssetUseCase,assetWalletUseCase *assetwallet.AssetWalletUseCase) *BuyAssetUseCase {
	return &BuyAssetUseCase{repo: repo, shareUseCase: shareUseCase, walletUseCase: walletUseCase, operationAssetUseCase: operationAssetUseCase,assetWalletUseCase: assetWalletUseCase}
}

func (uc *BuyAssetUseCase) BuyAsset(assetOperation entities.AssetOperation) (int, string) {

	if err := uc.validateBuyAssetInput(assetOperation); err != nil {
		return http.StatusNotAcceptable, err.Error()
	}

	value, err := uc.getAssetValue(assetOperation)
	if err != nil {
		return http.StatusNotAcceptable, err.Error()
	}

	if err := uc.executeOperation(assetOperation, value); err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	if err := uc.updateWallet(assetOperation, buildDatasToInsert(assetOperation, value, 1).AssetAmount); err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, "operação realizada com sucesso"
}

func (uc *BuyAssetUseCase) updateWallet(assetOperation entities.AssetOperation, assetAmount float64) error {
	idWallet, err := uc.walletUseCase.FindIdWallet(1)
	if err != nil {
		return errors.New("erro ao buscar carteira")
	}

	assetWallet, err := uc.assetWalletUseCase.FindAssetWallet(assetOperation.AssetCode, idWallet)
	if err != nil {
		if err.Error() == "ativo não existe na carteira" {
			walletAsset := entities.WalletAsset{
				AssetName:     assetOperation.AssetCode,
				AssetType:     assetOperation.AssetType,
				AssetQuantity: assetAmount,
				IdWallet:      idWallet,
			}
			if err := uc.assetWalletUseCase.InsertAssetIntoWallet(walletAsset); err != nil {
				return errors.New("erro ao inserir ativo")
			}
			return nil
		}
		return errors.New("erro ao buscar ativo")
	}

	assetWallet.AssetQuantity += assetAmount
	if err := uc.assetWalletUseCase.UpdateAssetIntoWallet(assetWallet.AssetQuantity, idWallet); err != nil {
		return errors.New("erro ao atualizar ativo")
	}

	return nil
}


func (uc *BuyAssetUseCase) executeOperation(assetOperation entities.AssetOperation, value float64) error {
	
	datasToInsert := buildDatasToInsert(assetOperation, value, 1)
	valueOperation := datasToInsert.AssetAmount * datasToInsert.AssetValue

	if err := uc.walletUseCase.VerifyIfInvestorCanOperate(1, valueOperation); err != nil {
		return errors.New("saldo insuficiente")
	}

	idOperation, err := uc.operationAssetUseCase.InsertOperationAsset(datasToInsert)
	if err != nil {
		return errors.New("erro ao concluir operação")
	}

	if err := uc.walletUseCase.InsertValueBalance(1, -valueOperation, idOperation); err != nil {
		return errors.New("erro ao atualizar saldo")
	}
	return nil
}


func (uc *BuyAssetUseCase) getAssetValue(assetOperation entities.AssetOperation) (float64, error) {
	if assetOperation.AssetType == "SHARE" {
		if !isActionTradable(time.Now()) {
			return 0, errors.New("mercado fechado")
		}
		if err := uc.isAssetValid(assetOperation.AssetCode); err != nil {
			return 0, errors.New("ação inválida")
		}
		return uc.getValueFromShare(assetOperation.AssetCode)
	}

	response, err := MakeRequestAsset(assetOperation.AssetType, assetOperation.AssetCode)
	if err != nil {
		return 0, errors.New("erro ao buscar ativo")
	}

	if assetOperation.AssetType == "COIN" {
		return getValueFromCoin(response, assetOperation.AssetCode)
	} else if assetOperation.AssetType == "CRYPTO" {
		return getValueFromCrypto(response)
	}

	return 0, errors.New("tipo de ativo inválido")
}

func (uc *BuyAssetUseCase) validateBuyAssetInput(assetOperation entities.AssetOperation) error {
	if err := ValidateFields(assetOperation); err != nil {
		return err
	}
	if assetOperation.OperationType != "BUY" {
		return errors.New("operação inválida. Somente operações de compra são permitidas")
	}
	return nil
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


