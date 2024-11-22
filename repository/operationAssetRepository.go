package repository

import (
	"database/sql"
	"fmt"
	"jumpStart-backEnd/entities"

	_ "github.com/go-sql-driver/mysql"
)


type OperationAssetRepository struct {
	db *sql.DB
}

func NewOperationAssetRepository(db *sql.DB) *OperationAssetRepository {
	return &OperationAssetRepository{db: db}
}

func (oar *OperationAssetRepository) InsertOperationAsset(datas entities.AssetInsertDataBase) (error) {
	
	query := fmt.Sprintf(`INSERT INTO tb_operationAsset(assetName,assetType,assetQuantity,assetValue,operationType,operationDate,idInvestor,isProcessedAlready) VALUES ('%s','%s',%f,%f,'%s','%s',%d,%t)`, datas.AssetName, datas.AssetType, datas.AssetAmount, datas.AssetValue, datas.OperationType, datas.OperationDate, datas.IdInvestor, datas.IsProcessedAlready)

	_, err := oar.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}