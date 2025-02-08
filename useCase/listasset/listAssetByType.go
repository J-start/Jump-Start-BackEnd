package listasset

import (
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/service/investor_service"
	"strings"
)

type ListAssetUseCase struct {
	repo *repository.ListAssetRepository
	investorService   *investor_service.InvestorService
}

func NewListAssetUseCase(repo *repository.ListAssetRepository,investorService *investor_service.InvestorService) *ListAssetUseCase {
	return &ListAssetUseCase{repo: repo,investorService: investorService}
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

func (lauc *ListAssetUseCase) GetListAssets(token string) ([]entities.ListAsset, error) {
    isAdm,err := lauc.investorService.IsAdm(token)
	if err != nil {
		return []entities.ListAsset{}, errors.New("ocorreu um erro, tente novamente")
	}
	if !isAdm {
		return []entities.ListAsset{}, errors.New("você não tem permissão para acessar essa funcionalidade")
	}
	listAssets,err := lauc.repo.FetchListAssetsAdm()
	if err != nil {
		return []entities.ListAsset{}, errors.New("ocorreu um erro, tente novamente")
	}
	return listAssets, nil
}

func (lauc *ListAssetUseCase) UpdateUrlImage(token string,datas entities.UpdateUrlImage) error{
	if datas.IdAsset <= 0 || datas.NewUrl == "" {
		return errors.New("algum dado é inválido")
	}
    isAdm,err := lauc.investorService.IsAdm(token)
	if err != nil {
		return errors.New("ocorreu um erro, tente novamente")
	}
	if !isAdm {
		return errors.New("você não tem permissão para acessar essa funcionalidade")
	}
	errUpdate := lauc.repo.UpdateAssetImageUrlById(datas.NewUrl,datas.IdAsset)
	if errUpdate != nil {
		return errors.New("ocorreu um erro, tente novamente")
	}
	return nil
}