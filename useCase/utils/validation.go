package utils

import (
	"errors"
	"strings"
	"time"

	"jumpStart-backEnd/entities"
	
)

func IsActionTradable(date time.Time) bool {
	OPEN := 10
	CLOSE := 17

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return false
	}

	brazilTime := date.In(location)

	if brazilTime.Weekday() == time.Saturday || brazilTime.Weekday() == time.Sunday {
		return false
	}

	if brazilTime.Hour() < OPEN || brazilTime.Hour() > CLOSE {
		return false
	}

	return true
}

func BuildDatasToInsert(assetOperation entities.AssetOperation, value float64, idInvestor int) entities.AssetInsertDataBase {
	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")

	return entities.AssetInsertDataBase{
		AssetName:          assetOperation.AssetName,
		AssetType:          assetOperation.AssetType,
		AssetAmount:        assetOperation.AssetAmount,
		AssetCode:          assetOperation.AssetCode,
		AssetValue:         value,
		OperationType:      assetOperation.OperationType,
		OperationDate:      formattedDate,
		IdInvestor:         idInvestor,
		IsProcessedAlready: false,
	}
}

func ValidateFields(assetOperation entities.AssetOperation) error{

	if assetOperation.AssetAmount <= 0 {
		return errors.New("quantidade de ativos inválida")
	}

	if assetOperation.AssetType != "CRYPTO" && assetOperation.AssetType != "COIN" && assetOperation.AssetType != "SHARE" {
		return errors.New("tipo de ativo inválido")
	}

	if assetOperation.OperationType != "BUY" && assetOperation.OperationType != "SELL" {
		return errors.New("tipo de operação inválida")
	}

	if assetOperation.AssetCode == "" {
		return errors.New("código de ativo inválido")
	}

	if assetOperation.AssetName == "" || len(strings.Split(assetOperation.AssetName, " ")) == 0 || len(assetOperation.AssetName) == 0 || len(assetOperation.AssetName) > 255 {
		return errors.New("nome de ativo inválido")
	}
	
	return nil

}