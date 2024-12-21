package entities


type WalletAsset struct {
	Id int `json:"id"`
	AssetName string `json:"assetName"`
	AssetType string `json:"assetType"`
	AssetQuantity float64 `json:"assetQuantity"`
	IdWallet int `json:"idWallet"`

}

type Asset struct {
	AssetName string
	AssetType string
	AssetQuantity float64
}

type WalletDatas struct {
	InvestorBalance float64 
	Assets []Asset
}

type WalletRequest struct {
	TokenInvestor string `json:"tokenInvestor"`
}
type WalletOperation struct {
	OperationType  string
	OperationValue float64
	DperationDate  string
}

	