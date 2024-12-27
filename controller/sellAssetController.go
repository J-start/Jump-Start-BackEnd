package controller

import (
	"encoding/json"

	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/useCase/sell"
	"net/http"
	"jumpStart-backEnd/handleError"
	//"strconv"
)

type SellAssetController struct {
	useCase *sell.SellAssetUseCase
}

func NewSellAssetController(useCase *sell.SellAssetUseCase) *SellAssetController {
	return &SellAssetController{useCase: useCase}
}

func (bac *SellAssetController) SellAsset(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")

	code,token := getTokenJWT(r)
	if code != http.StatusOK {
		handleError.WriteHTTPStatus(w, code, token)
		return
	}

	var asset entities.AssetOperation

	err := json.NewDecoder(r.Body).Decode(&asset)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	asset.CodeInvestor = token

		code,message := bac.useCase.SellAsset(asset)
		handleError.WriteHTTPStatus(w, code, message)
	
}

