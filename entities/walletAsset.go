package entities


type WalletAsset struct {
	Id int `json:"id"`
	AssetName string `json:"assetName"`
	AssetType string `json:"assetType"`
	AssetQuantity float64 `json:"assetQuantity"`
	IdWallet int `json:"idWallet"`

}
	