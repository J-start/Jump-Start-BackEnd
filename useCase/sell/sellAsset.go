package sell

import (
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/useCase"
	"jumpStart-backEnd/useCase/assetwallet"
	"jumpStart-backEnd/useCase/operation"
	"jumpStart-backEnd/useCase/utils"
	"jumpStart-backEnd/useCase/utils/requests"
	"jumpStart-backEnd/useCase/wallet"
	"net/http"
	"strings"
	"time"
)

type SellAssetUseCase struct {
	repo                  *repository.ShareRepository
	shareUseCase          *usecase.ShareUseCase
	walletUseCase         *wallet.WalletUseCase
	operationAssetUseCase *operation.OperationAssetUseCase
	assetWalletUseCase    *assetwallet.AssetWalletUseCase
}

func NewSellAssetsUseCase(repo *repository.ShareRepository, shareUseCase *usecase.ShareUseCase, walletUseCase *wallet.WalletUseCase, operationAssetUseCase *operation.OperationAssetUseCase, assetWalletUseCase *assetwallet.AssetWalletUseCase) *SellAssetUseCase {
	return &SellAssetUseCase{repo: repo, shareUseCase: shareUseCase, walletUseCase: walletUseCase, operationAssetUseCase: operationAssetUseCase, assetWalletUseCase: assetWalletUseCase}
}

func (uc *SellAssetUseCase) SellAsset(assetOperation entities.AssetOperation) (int, string) {

	if err := uc.validateSellAssetInput(assetOperation); err != nil {
		return http.StatusNotAcceptable,err.Error()
	}
	if err := uc.VerifyIfInvestorCanSell(assetOperation); err != nil {
		return http.StatusNotAcceptable,err.Error()
	}

	valueAsset, err := uc.getAssetValue(assetOperation)
	if err != nil {
		return http.StatusNotAcceptable, err.Error()
	}

	idOperation, err := uc.ExecuteOperationRegisterAsset(assetOperation,valueAsset)
	if err != nil {
		return http.StatusInternalServerError,err.Error()
	}

	valueAsset *= assetOperation.AssetAmount

	err = uc.UpdateWallet(assetOperation,valueAsset,idOperation)

	if err != nil {
		return http.StatusInternalServerError,err.Error()
	}

	return 200,"Venda realizada com sucesso"
}

func (uc *SellAssetUseCase) updateWallet(assetOperation entities.AssetOperation, assetAmount float64) error {
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
}
		
	assetWallet.AssetQuantity -= assetAmount

	if assetWallet.AssetQuantity == 0 {
		if err := uc.assetWalletUseCase.DeleteAssetWallet(assetWallet.Id); err != nil {
			return errors.New("erro ao deletar ativo")
		}	
		return nil
	}
	if err := uc.assetWalletUseCase.UpdateAssetIntoWallet(assetWallet.AssetQuantity, assetWallet.Id); err != nil {
		return errors.New("erro ao atualizar ativo")
	}

	return nil
}

func (uc *SellAssetUseCase) VerifyIfInvestorCanSell(assetOperation entities.AssetOperation) error {

	walletAsset, err := uc.assetWalletUseCase.FindAssetWallet(assetOperation.AssetCode, 1)
	if err != nil {
		if err.Error() == "ativo não existe na carteira" {
			return fmt.Errorf("o ativo %s não existe em carteira", assetOperation.AssetCode)
		}
		return  err
	}
	if assetOperation.AssetAmount > walletAsset.AssetQuantity {
		return  fmt.Errorf("quantidade de ativos em carteira insuficiente")
	}
	
	return nil   
}

func (uc *SellAssetUseCase) getAssetValue(assetOperation entities.AssetOperation) (float64, error) {
	if assetOperation.AssetType == "SHARE" {
		if !utils.IsActionTradable(time.Now()) {
			return 0, errors.New("mercado fechado")
		}
		if err := uc.isAssetValid(assetOperation.AssetCode); err != nil {
			return 0, errors.New("ação inválida")
		}
		return uc.getValueFromShare(assetOperation.AssetCode)
	}

	response, err := requests.MakeRequestAsset(assetOperation.AssetType, assetOperation.AssetCode)
	if err != nil {
		return 0, errors.New("erro ao buscar ativo")
	}

	if assetOperation.AssetType == "COIN" {
		return requests.GetValueFromCoin(response, assetOperation.AssetCode)
	} else if assetOperation.AssetType == "CRYPTO" {
		return requests.GetValueFromCrypto(response)
	}

	return 0, errors.New("tipo de ativo inválido")
}

func (uc *SellAssetUseCase) getValueFromShare(code string) (float64, error) {
	share, err := uc.repo.FindShareById(code)
	if err != nil {
		return 0, err
	}
	return share.CloseShare, nil
}

func (uc *SellAssetUseCase) isAssetValid(code string) error {
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

func (uc *SellAssetUseCase) validateSellAssetInput(assetOperation entities.AssetOperation) error {
	if err := utils.ValidateFields(assetOperation); err != nil {
		return err
	}
	if assetOperation.OperationType != "SELL" {
		return errors.New("operação inválida. Somente operações de compra são permitidas")
	}
	return nil
}

func (uc *SellAssetUseCase) ExecuteOperationRegisterAsset(assetOperation entities.AssetOperation,valueAsset float64) (int64, error) {
	datasToInsert := utils.BuildDatasToInsert(assetOperation, valueAsset, 1)

	idOperation, err := uc.operationAssetUseCase.InsertOperationAsset(datasToInsert)

	if err != nil {
		return -1,errors.New("erro ao concluir operação, tente novamente")
	}
	return idOperation,nil
}

func (uc *SellAssetUseCase) UpdateWallet(assetOperation entities.AssetOperation,valueAsset float64,idOperation int64) error {
	if err := uc.walletUseCase.InsertValueBalance(1, valueAsset, idOperation); err != nil {
		return errors.New("erro ao atualizar saldo, tente realizar a operação novamente")
	}

	if err := uc.updateWallet(assetOperation, assetOperation.AssetAmount); err != nil {
		return errors.New("erro ao atualizar ativo na carteira")
	}
	return nil
}
