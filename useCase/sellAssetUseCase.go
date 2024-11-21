package usecase

import (
	"encoding/json"
	"errors"
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
	shareUseCase *ShareUseCase
}

func NewSellAssetsUseCase(repo *repository.ShareRepository,shareUseCase *ShareUseCase) *SellAssetsUseCase {
	return &SellAssetsUseCase{repo: repo, shareUseCase: shareUseCase}
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
	err := uc.isAssetValid(assetOperation.AssetCode)
	if err != nil {
		fmt.Println("Ação inválida")
		return 
	}

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

		if !isActionTradable(time.Now()) {
			fmt.Println("Ação não pode ser comprada ou vendida")
			return
		}

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

func isActionTradable(date time.Time ) bool {

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

func (uc *SellAssetsUseCase) isAssetValid(code string) error {
	if code == "" || len(strings.Split(code, " ")) == 0 || len(code) == 0 {
		return  errors.New("código de ativo inválido")
	}

	isContainName, err := uc.shareUseCase.ShareNameList(code)
	if err != nil {
		return  err
	}
	if !isContainName {
		return  errors.New("ação não encontrada")
	}

	return  nil
}
