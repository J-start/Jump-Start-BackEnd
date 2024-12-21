package listasset

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"strings"
)

type ListAssetUseCase struct {
	repo *repository.ListAssetRepository
}

func NewListAssetUseCase(repo *repository.ListAssetRepository) *ListAssetUseCase {
	return &ListAssetUseCase{repo: repo}
}

func (lauc *ListAssetUseCase) ListAssetByType(asset string) ([]entities.ListAsset, error) {
	if asset == "" || (asset != "CRYPTO" && asset != "SHARE" && asset != "COIN") {
		return []entities.ListAsset{}, errors.New("tipo de ativo inválido")	
	}
	listAsset,err := lauc.repo.ListAsset(asset)
	if err != nil {
		return []entities.ListAsset{}, errors.New("ocorreu um erro, tente novamente")
	}
	return listAsset,nil
}

func (lauc *ListAssetUseCase) ListAssetRequest(asset string) (string, error) {
	if asset == "" || (asset != "CRYPTO" && asset != "SHARE" && asset != "COIN") {
		return "", errors.New("tipo de ativo inválido")	
	}
	listAssetRequest,err := lauc.repo.ListAssetRequest(asset)
	if err != nil {
		return "", errors.New("ocorreu um erro, tente novamente")
	}
	listAssetRequestString := strings.Join(listAssetRequest, ",")
	return listAssetRequestString, nil
}