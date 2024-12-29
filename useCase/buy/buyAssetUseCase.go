package buy

import (
	"database/sql"
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/service/investor_service"
	"jumpStart-backEnd/serviceRepository"
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

type BuyAssetUseCase struct {
	repo                     *repository.ShareRepository
	shareUseCase             *usecase.ShareUseCase
	walletUseCase            *wallet.WalletUseCase
	operationAssetUseCase    *operation.OperationAssetUseCase
	assetWalletUseCase       *assetwallet.AssetWalletUseCase
	repositoryService 	     *servicerepository.ServiceRepository
	investorService          *investor_service.InvestorService
}

func NewBuyAssetsUseCase(repo *repository.ShareRepository, shareUseCase *usecase.ShareUseCase, walletUseCase *wallet.WalletUseCase, operationAssetUseCase *operation.OperationAssetUseCase,assetWalletUseCase *assetwallet.AssetWalletUseCase, repositoryService  *servicerepository.ServiceRepository,investorService *investor_service.InvestorService) *BuyAssetUseCase {
	return &BuyAssetUseCase{repo: repo, 
							shareUseCase: shareUseCase, 
							walletUseCase: walletUseCase, 
							operationAssetUseCase: operationAssetUseCase,
							assetWalletUseCase: assetWalletUseCase, 
							repositoryService: repositoryService,
							investorService: investorService}

}

func (uc *BuyAssetUseCase) BuyAsset(assetOperation entities.AssetOperation) (int, string) {

	if err := uc.validateBuyAssetInput(assetOperation); err != nil {
		return http.StatusNotAcceptable, err.Error()
	}

	repositoryService,err := uc.repositoryService.StartTransaction()
	if err != nil {
		return http.StatusInternalServerError, errors.New("erro ao processar requisição, tente novamente").Error()
	}

	value, err := uc.getAssetValue(assetOperation)
	if err != nil {
		return http.StatusNotAcceptable, err.Error()
	}
	idInvestor := 2
	//  idInvestor,err := uc.investorService.GetIdByToken(assetOperation.CodeInvestor)
	//  if err != nil {
	//  	repositoryService.Rollback()
	//  	return http.StatusInternalServerError, errors.New("erro ao processar requisição, tente novamente").Error()
	//  }

	if err := uc.executeOperation(assetOperation,idInvestor,value,repositoryService); err != nil {
		repositoryService.Rollback()
		return http.StatusInternalServerError, err.Error()
	}

	amount := utils.BuildDatasToInsert(assetOperation, value, idInvestor).AssetAmount

	if err := uc.updateWallet(assetOperation,idInvestor,amount,repositoryService); err != nil {
		repositoryService.Rollback()
		return http.StatusInternalServerError, err.Error()
	}

	errService := repositoryService.Commit()
	if errService != nil {
		repositoryService.Rollback()
		return http.StatusInternalServerError, errors.New("erro ao processar requisição, tente novamente").Error()
	}

	return http.StatusOK, "operação realizada com sucesso"
}

func (uc *BuyAssetUseCase) updateWallet(assetOperation entities.AssetOperation,idInvestor int, assetAmount float64,repositoryService *sql.Tx) error {
	idWallet, err := uc.walletUseCase.FindIdWallet(idInvestor)
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
			if err := uc.assetWalletUseCase.InsertAssetIntoWallet(walletAsset,repositoryService); err != nil {
				fmt.Println(err)
				return errors.New("erro ao inserir ativo")
			}
			return nil
		}
		return errors.New("erro ao buscar ativo")
	}

	assetWallet.AssetQuantity += assetAmount
	
	if err := uc.assetWalletUseCase.UpdateAssetIntoWallet(assetWallet.AssetQuantity, assetWallet.Id,repositoryService); err != nil {
		return errors.New("erro ao atualizar ativo")
	}

	return nil
}


func (uc *BuyAssetUseCase) executeOperation(assetOperation entities.AssetOperation,idInvestor int, value float64,repositoryService *sql.Tx) error {
	
	

	datasToInsert := utils.BuildDatasToInsert(assetOperation, value, idInvestor)
	valueOperation := datasToInsert.AssetAmount * datasToInsert.AssetValue

	if err := uc.walletUseCase.VerifyIfInvestorCanOperate(idInvestor, valueOperation); err != nil {
		return errors.New("saldo insuficiente")
	}

	idOperation, err := uc.operationAssetUseCase.InsertOperationAsset(datasToInsert,repositoryService)
	if err != nil {
		return errors.New("erro ao concluir operação")
	}

	if err := uc.walletUseCase.InsertValueBalance(idInvestor, -valueOperation, idOperation,repositoryService); err != nil {
		return errors.New("erro ao atualizar saldo")
	}
	return nil
}


func (uc *BuyAssetUseCase) getAssetValue(assetOperation entities.AssetOperation) (float64, error) {
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

func (uc *BuyAssetUseCase) validateBuyAssetInput(assetOperation entities.AssetOperation) error {
	if assetOperation.AssetType == "SHARE" {
		if !utils.IsActionTradable(time.Now()) {
			return errors.New("mercado fechado")
		}
	}
	if err := utils.ValidateFields(assetOperation); err != nil {
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


