package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


type OperationAssetRepository struct {
	db *sql.DB
}

func NewOperationAssetRepository(db *sql.DB) *OperationAssetRepository {
	return &OperationAssetRepository{db: db}
}
func (oar *OperationAssetRepository) InsertOperationAsset(datas entities.AssetInsertDataBase,repositoryService *sql.Tx) (int64,error) {
	
	tx := repositoryService
	query := `INSERT INTO tb_operationAsset(assetName, assetType,assetCode, assetQuantity, assetValue, operationType, operationDate, idInvestor, isProcessedAlready) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1,err
	}
	defer stmt.Close()

	
	result, err := stmt.Exec(datas.AssetName, datas.AssetType,datas.AssetCode, datas.AssetAmount, datas.AssetValue, datas.OperationType, datas.OperationDate, datas.IdInvestor, datas.IsProcessedAlready)
	if err != nil {
		return -1,err
	}

	idOperation, err := result.LastInsertId()
	if err != nil {
		return -1,err
	}
	
	return idOperation,nil
}

func (oar *OperationAssetRepository) FetchAssetHistoryByInvestor(idInvestor int,offset int) ([]entities.AssetOperationHistory,error) {

	/*
		offset start with value 0, and by then it will be multiplied by 10, returning 10 firts results. 
		if offset is 1, it will be multiplied by 10, returning 10 results from 10 to 20, and so on.
		
		To use this endpoint, simply sum the value of the offset by 1 to use pagginated results.
		OBS: offset must initate with value 0
	*/
	offset *= 20
	query := fmt.Sprintf(`SELECT assetName,assetType,assetQuantity,assetValue,operationType,operationDate FROM tb_operationAsset WHERE idInvestor = %d LIMIT 20 OFFSET %d;`, idInvestor,offset)

	rows, err := oar.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return []entities.AssetOperationHistory{},errors.New("erro ao buscar histórico de ativos")
	}

	assetHistory,errB := buildAssetHistory(rows)
	if errB != nil {
		fmt.Println(errB)
		return []entities.AssetOperationHistory{},errors.New("erro ao buscar histórico de ativos")
	}

	return assetHistory,nil

}

func buildAssetHistory(rows *sql.Rows) ([]entities.AssetOperationHistory, error) {
	var assetsHistory []entities.AssetOperationHistory

	for rows.Next() {
		var dateShare time.Time
		var assetHistory entities.AssetOperationHistory
		err := rows.Scan(&assetHistory.AssetName, &assetHistory.AssetType, &assetHistory.AssetQuantity, &assetHistory.AssetValue,&assetHistory.OperationType, &dateShare)
		if err != nil {
			return nil, err
		}
		assetHistory.OperationDate = dateShare.Format("2006-01-02")

		assetsHistory = append(assetsHistory, assetHistory)

	}

	return assetsHistory, nil

}



