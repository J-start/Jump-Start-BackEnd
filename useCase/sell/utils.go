package sell

import (
	"fmt"
	"time"

	"jumpStart-backEnd/entities"
)

func isActionTradable(date time.Time) bool {
	OPEN := 10
	CLOSE := 17

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o fuso horário:", err)
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

func buildDatasToInsert(assetOperation entities.AssetOperation, value float64, idInvestor int) entities.AssetInsertDataBase {
	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")

	return entities.AssetInsertDataBase{
		AssetName:          assetOperation.AssetName,
		AssetType:          assetOperation.AssetType,
		AssetAmount:        assetOperation.AssetAmount,
		AssetValue:         value,
		OperationType:      assetOperation.OperationType,
		OperationDate:      formattedDate,
		IdInvestor:         idInvestor,
		IsProcessedAlready: false,
	}
}