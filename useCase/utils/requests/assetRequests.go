package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"jumpStart-backEnd/entities"
)

func MakeRequestAsset(assetType, assetCode string) (string, error) {
	url, err := buildUrl(assetType, assetCode)
	if err != nil || url == "" {
		return "", errors.New("erro ao construir a URL")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code diferente de 200: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(body) == 0 || string(body) == "[]" {
		return "", errors.New("resposta vazia ou inválida")
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

func GetValueFromCoin(response, code string) (float64, error) {
	if len(response) == 0{
		return 0,errors.New("moeda inválida")
	}
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

	if bidFloat == 0 {
		return 0, errors.New("valor do ativo é zero")
	}

	return bidFloat, nil
}

func GetValueFromCrypto(response string) (float64, error) {
	if len(response) == 0{
		return 0,errors.New("crypto inválida")
	}
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