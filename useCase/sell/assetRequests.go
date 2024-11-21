package sell

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"jumpStart-backEnd/entities"
)

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

func buildUrl(typeAsset, codeAsset string) (string, error) {
	switch typeAsset {
	case "COIN":
		return fmt.Sprintf(`https://economia.awesomeapi.com.br/json/last/%s`, codeAsset), nil
	case "CRYPTO":
		return fmt.Sprintf(`https://api.mercadobitcoin.net/api/v4/tickers?symbols=%s`, codeAsset), nil
	}

	return "", fmt.Errorf("tipo de ativo não encontrado")
}

func getValueFromCoin(response, code string) (float64, error) {
	var data map[string]entities.Coin
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		return 0, err
	}

	code = strings.ReplaceAll(code, "-", "")
	bidFloat, err := strconv.ParseFloat(data[code].Bid, 64)
	if err != nil {
		return 0, err
	}

	return bidFloat, nil
}

func getValueFromCrypto(response string) (float64, error) {
	var data []entities.Crypto
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		return 0, err
	}

	bidFloat, err := strconv.ParseFloat(data[0].Last, 64)
	if err != nil {
		return 0, err
	}

	return bidFloat, nil
}