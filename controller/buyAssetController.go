package controller

import (
	"encoding/json"

	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/useCase/buy"
	"net/http"
	"jumpStart-backEnd/handleError"
	//"strconv"
)

type BuyAssetController struct {
	useCase *buy.BuyAssetUseCase
}

func NewBuyAssetController(useCase *buy.BuyAssetUseCase) *BuyAssetController {
	return &BuyAssetController{useCase: useCase}
}

func (bac *BuyAssetController) BuyAsset(w http.ResponseWriter, r *http.Request) {
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

	code,message := bac.useCase.BuyAsset(asset)
	handleError.WriteHTTPStatus(w, code, message)
		
}

