package assetwallet

import (
	"database/sql"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
)

type AssetWalletUseCase struct {
	repo *repository.WalletAssetRepository
}

func NewAssetWalletUseCase(repo *repository.WalletAssetRepository) *AssetWalletUseCase {
	return &AssetWalletUseCase{repo: repo}
}

func (uc *AssetWalletUseCase) FindAssetWallet(assetName string,idWallet int) (entities.WalletAsset, error) {
	walletAsset, err := uc.repo.FindAssetWallet(assetName,idWallet)
	if err != nil {
		return entities.WalletAsset{}, err
	}
	return walletAsset, nil
}

func (uc *AssetWalletUseCase) InsertAssetIntoWallet(walletAsset entities.WalletAsset,repositoryService *sql.Tx) error {
	err := uc.repo.InsertAssetIntoWallet(walletAsset,repositoryService)
	if err != nil {
		return err
	}
	return nil
}

func (uc *AssetWalletUseCase) UpdateAssetIntoWallet(newQuantity float64, idWallet int,repositoryService *sql.Tx) error {
	err := uc.repo.UpdateAssetIntoWallet(newQuantity, idWallet,repositoryService)
	if err != nil {
		return err
	}
	return nil
}
 func (uc *AssetWalletUseCase) DeleteAssetWallet(idAsset int) error {
	err := uc.repo.DeleteAssetWallet(idAsset)
	if err != nil {
		return err
	}
	return nil
}