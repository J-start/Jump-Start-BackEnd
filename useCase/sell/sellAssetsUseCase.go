package sell

import (
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/useCase"
	"strings"
	"time"
)

type SellAssetsUseCase struct {
	repo         *repository.ShareRepository
	shareUseCase *usecase.ShareUseCase
}

func NewSellAssetsUseCase(repo *repository.ShareRepository, shareUseCase *usecase.ShareUseCase) *SellAssetsUseCase {
	return &SellAssetsUseCase{repo: repo, shareUseCase: shareUseCase}
}

func (uc *SellAssetsUseCase) ManipulationAsset(assetOperation entities.AssetOperation) {
	var value float64
	if assetOperation.AssetType != "SHARE" {
		response, err := MakeRequestAsset(assetOperation.AssetType, assetOperation.AssetCode)
		if err != nil {
			fmt.Println(err)
		}

		if assetOperation.AssetType == "COIN" {
			valueReturn, err := getValueFromCoin(response, assetOperation.AssetCode)
			if err != nil {
				fmt.Println(err)
			}
			value = valueReturn
		} else if assetOperation.AssetType == "CRYPTO" {
			valueReturn, err := getValueFromCrypto(response)
			if err != nil {
				fmt.Println(err)
			}
			value = valueReturn
		}
	} else {
		if !isActionTradable(time.Now()) {
			fmt.Println("Ação não pode ser comprada ou vendida")
			return
		}

		err := uc.isAssetValid(assetOperation.AssetCode)
		if err != nil {
			fmt.Println("Ação inválida")
			return
		}

		valueReturn, err := uc.getValueFromShare(assetOperation.AssetCode)
		if err != nil {
			fmt.Println(err)
		}
		value = valueReturn
	}

	fmt.Println(buildDatasToInsert(assetOperation, value, 1))
}

func (uc *SellAssetsUseCase) getValueFromShare(code string) (float64, error) {
	share, err := uc.repo.FindShareById(code)
	if err != nil {
		return 0, err
	}
	return share.CloseShare, nil
}

func (uc *SellAssetsUseCase) isAssetValid(code string) error {
	if code == "" || len(strings.Split(code, " ")) == 0 || len(code) == 0 {
		return errors.New("código de ativo inválido")
	}

	isContainName, err := uc.shareUseCase.ShareNameList(code)
	if err != nil {
		return err
	}
	if !isContainName {
		return errors.New("ação não encontrada")
	}

	return nil
}