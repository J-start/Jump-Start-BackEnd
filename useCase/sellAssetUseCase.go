package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PairData struct {
	AssetCode  string `json:"assetCode"`
	AssetValue string `json:"assetValue"`
}
type SellAssetsUseCase struct {
	repo *repository.ShareRepository
}

func NewSellAssetsUseCase(repo *repository.ShareRepository) *SellAssetsUseCase {
	return &SellAssetsUseCase{repo: repo}
}

func MakeRequestAsset(assetType, assetCode string) (string, error) {

	url, err := buildUrl(assetType, assetCode)
	if err != nil && url == "" {
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func buildUrl(typeAsset string, codeAsset string) (string, error) {
	switch typeAsset {
	case "COIN":
		return fmt.Sprintf(`https://economia.awesomeapi.com.br/json/last/%s`, codeAsset), nil
	case "CRYPTO":
		return fmt.Sprintf(`https://api.mercadobitcoin.net/api/v4/tickers?symbols=%s`, codeAsset), nil
	}

	return "", fmt.Errorf("tipo de ativo não encontrado")
}

func (uc *SellAssetsUseCase) ManipulationAsset(assetOperation entities.AssetOperation) {
	//TODO - Implemnetar lógica para verificar se a ação pode ser comprada ou vendida
	var value float64
	if assetOperation.AssetType != "SHARE" {
		response, err := MakeRequestAsset(assetOperation.AssetType, assetOperation.AssetCode)

		if err != nil {
			fmt.Println(err)
		}

		if assetOperation.AssetType == "COIN" {
			valueReturn, err := getValueFromCoin(response, assetOperation.AssetCode)
			if err != nil {
				fmt.Println(err)
			}
			value = valueReturn
		} else if assetOperation.AssetType == "CRYPTO" {
			valueReturn, err := getValueFromCrypto(response)
			if err != nil {
				fmt.Println(err)
			}
			value = valueReturn
		}

	} else {
		valueReturn, err := uc.getValueFromShare(assetOperation.AssetCode)
		if err != nil {
			fmt.Println(err)
		}
		value = valueReturn
	}

	fmt.Println(buildDatasToInsert(assetOperation, value, 1))

}

func getValueFromCoin(response string, code string) (float64, error) {
	var data map[string]entities.Coin

	err := json.Unmarshal([]byte(response), &data)

	if err != nil {
		fmt.Println(err)
	}

	code = strings.ReplaceAll(code, "-", "")
	bidFloat, err := strconv.ParseFloat(data[code].Bid, 64)

	if err != nil {
		fmt.Println(err)
	}

	return bidFloat, nil
}

func getValueFromCrypto(response string) (float64, error) {

	var data []entities.Crypto

	err := json.Unmarshal([]byte(response), &data)

	if err != nil {
		fmt.Println(err)
	}
	bidFloat, err := strconv.ParseFloat(data[0].Last, 64)

	if err != nil {
		fmt.Println(err)
	}

	return bidFloat, nil
}

func (uc *SellAssetsUseCase) getValueFromShare(code string) (float64, error) {
	//TODO - Implementar lógica JWT para obter o id do usuário
	//TODO - Implemnetar lógica para verificar se a ação existe
	if code == "" || len(strings.Split(code, " ")) == 0 || len(code) == 0 {
		return 0, nil
	}

	share, err := uc.repo.FindShareById(code)

	if err != nil {
		return 0, err
	}

	return share.CloseShare, nil
}

func buildDatasToInsert(assetOperation entities.AssetOperation, value float64, idInvestor int) entities.AssetInsertDataBase {
	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")

	datasReturn := entities.AssetInsertDataBase{
		AssetName:          assetOperation.AssetName,
		AssetType:          assetOperation.AssetType,
		AssetAmount:        assetOperation.AssetAmount,
		AssetValue:         value,
		OperationType:      assetOperation.OperationType,
		OperationDate:      formattedDate,
		IdInvestor:         idInvestor,
		IsProcessedAlready: false,
	}

	return datasReturn
}
