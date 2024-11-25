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
func (oar *OperationAssetRepository) InsertOperationAsset(datas entities.AssetInsertDataBase) (int64,error) {
	
	tx, err := oar.db.Begin()
	if err != nil {
		return -1,err
	}

	query := `INSERT INTO tb_operationAsset(assetName, assetType,assetCode, assetQuantity, assetValue, operationType, operationDate, idInvestor, isProcessedAlready) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback() 
		return -1,err
	}
	defer stmt.Close()

	
	result, err := stmt.Exec(datas.AssetName, datas.AssetType,datas.AssetCode, datas.AssetAmount, datas.AssetValue, datas.OperationType, datas.OperationDate, datas.IdInvestor, datas.IsProcessedAlready)
	if err != nil {
		tx.Rollback() 
		return -1,err
	}

	idOperation, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return -1,err
	}
	
	err = tx.Commit()
	if err != nil {
		tx.Rollback() 
		return -1,err
	}
	return idOperation,nil
}



