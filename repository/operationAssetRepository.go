package repository

import (
	"database/sql"
	"jumpStart-backEnd/entities"

	_ "github.com/go-sql-driver/mysql"
)


type OperationAssetRepository struct {
	db *sql.DB
}

func NewOperationAssetRepository(db *sql.DB) *OperationAssetRepository {
	return &OperationAssetRepository{db: db}
}
func (oar *OperationAssetRepository) InsertOperationAsset(datas entities.AssetInsertDataBase) error {
	
	tx, err := oar.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO tb_operationAsset(assetName, assetType,assetCode, assetQuantity, assetValue, operationType, operationDate, idInvestor, isProcessedAlready) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback() 
		return err
	}
	defer stmt.Close()

	
	_, err = stmt.Exec(datas.AssetName, datas.AssetType,datas.AssetCode, datas.AssetAmount, datas.AssetValue, datas.OperationType, datas.OperationDate, datas.IdInvestor, datas.IsProcessedAlready)
	if err != nil {
		tx.Rollback() 
		return err
	}

	
	err = tx.Commit()
	if err != nil {
		tx.Rollback() 
		return err
	}

	return nil
}
