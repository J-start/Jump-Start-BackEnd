package controller

import (
	"encoding/json"

	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/useCase/buy"
	"net/http"
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

	var asset entities.AssetOperation

	err := json.NewDecoder(r.Body).Decode(&asset)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = bac.useCase.BuyAsset(asset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

