package utils

import (
	"strings"
	"testing"
	"time"

	"jumpStart-backEnd/entities"

	"github.com/stretchr/testify/require"
)

func TestIsActionTradable(t *testing.T) {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil{
		t.Error("Erro ao obter hora de são paulo")
	}

	tests := []struct {
		name     string
		date     time.Time
		expected bool
	}{
		{"Weekday within trading hours", time.Date(2025, 02, 14, 12, 30, 0, 0, location), true},
		{"Weekday before trading hours", time.Date(2025, 02, 15, 10, 0, 0, 0, location), false},
		{"Weekday after trading hours", time.Date(2025, 02, 13, 19, 0, 0, 0, location), false},
		{"Saturday", time.Date(2025, 02, 16, 12, 0, 0, 0, location), false},
		{"Sunday", time.Date(2025, 02, 15, 12, 0, 0, 0, location), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsActionTradable(tt.date)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestBuildDatasToInsert(t *testing.T) {
	assetOperation := entities.AssetOperation{
		AssetName:     "Test Asset",
		AssetType:     "SHARE",
		AssetAmount:   10,
		AssetCode:     "TST",
		OperationType: "BUY",
	}
	value := 100.0
	idInvestor := 1

	result := BuildDatasToInsert(assetOperation, value, idInvestor)
	
	require.Equal(t, assetOperation.AssetName, result.AssetName)
	require.Equal(t, assetOperation.AssetType, result.AssetType)
	require.Equal(t, assetOperation.AssetAmount, result.AssetAmount)
	require.Equal(t, assetOperation.AssetCode, result.AssetCode)
	require.Equal(t, value, result.AssetValue)
	require.Equal(t, assetOperation.OperationType, result.OperationType)
	require.Equal(t, idInvestor, result.IdInvestor)
	require.False(t, result.IsProcessedAlready)
}

 func TestValidateFields(t *testing.T) {
 	tests := []struct {
 		name          string
 		assetOperation entities.AssetOperation
 		expectedError string
 	}{
 		{"Valid fields", entities.AssetOperation{AssetName: "Test Asset",AssetCode: "TEST-BRL", AssetType: "SHARE", AssetAmount: 10, OperationType: "BUY", CodeInvestor: "CODE_INVESTOR"}, ""},
 		{"Invalid asset amount", entities.AssetOperation{AssetName: "Test Asset",AssetCode: "TEST-BRL", AssetType: "SHARE", AssetAmount: 0, OperationType: "BUY", CodeInvestor: "CODE_INVESTOR"}, "quantidade de ativos inválida"},
 		{"Invalid asset type", entities.AssetOperation{AssetName: "Test Asset",AssetCode: "TEST-BRL", AssetType: "INVALID", AssetAmount: 10, OperationType: "SELL", CodeInvestor: "CODE_INVESTOR"}, "tipo de ativo inválido"},
 		{"Invalid operation type", entities.AssetOperation{AssetName: "Test Asset",AssetCode: "TEST-BRL", AssetType: "SHARE", AssetAmount: 10, OperationType: "INVALID", CodeInvestor: "CODE_INVESTOR"}, "tipo de operação inválida"},
 		{"Empty asset code", entities.AssetOperation{AssetName: "Test Asset",AssetCode: "" ,AssetType: "SHARE", AssetAmount: 10, OperationType: "SELL", CodeInvestor: "CODE_INVESTOR"}, "código de ativo inválido"},
 		{"Empty asset name", entities.AssetOperation{AssetName: "",AssetCode: "TEST-BRL", AssetType: "SHARE", AssetAmount: 10, OperationType: "SELL", CodeInvestor: "CODE_INVESTOR"}, "nome de ativo inválido"},
 		{"Asset name too long", entities.AssetOperation{AssetName: strings.Repeat("a", 300), AssetCode: "SHARE-BRL", AssetType:"SHARE", AssetAmount: 10, OperationType: "BUY",CodeInvestor: "CODE_INVESTOR"}, "nome de ativo inválido"},
 	}

 	for _, tt := range tests {
 		t.Run(tt.name, func(t *testing.T) {
 			err := ValidateFields(tt.assetOperation)
 			if tt.expectedError == "" {
 				require.Nil(t, err)
 			} else {
 				require.EqualError(t, err, tt.expectedError)
 			}
 		})
 	}
 }