package entities


type AssetOperation struct {
	AssetName     string
	AssetCode     string
	AssetType     string
	AssetAmount   float64
	OperationType string
	CodeInvestor  string
}

type AssetInsertDataBase struct {
	IdAsset            int
	AssetName          string
	AssetType          string
	AssetCode		   string
	AssetAmount        float64
	AssetValue         float64
	OperationType      string
	OperationDate      string
	IdInvestor         int
	IsProcessedAlready bool
}
