package assetwallet

import (
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

func (uc *AssetWalletUseCase) InsertAssetIntoWallet(walletAsset entities.WalletAsset) error {
	err := uc.repo.InsertAssetIntoWallet(walletAsset)
	if err != nil {
		return err
	}
	return nil
}

func (uc *AssetWalletUseCase) UpdateAssetIntoWallet(newQuantity float64, idWallet int) error {
	err := uc.repo.UpdateAssetIntoWallet(newQuantity, idWallet)
	if err != nil {
		return err
	}
	return nil
}
