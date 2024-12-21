package controller

import (
	"encoding/json"
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/handleError"
	"jumpStart-backEnd/useCase/wallet"
	"net/http"
)

type WalletController struct {
	useCase *wallet.WalletUseCase
}

func NewWalletController(useCase *wallet.WalletUseCase) *WalletController {
	return &WalletController{useCase: useCase}
}


func (wc *WalletController) FetchDatasWallet(w http.ResponseWriter, r *http.Request) {
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

	var tokenInvestor entities.WalletRequest

	err := json.NewDecoder(r.Body).Decode(&tokenInvestor)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errors.New("corpo da requisição inconsistente").Error())
		return
	}
	walletDatas,err := wc.useCase.FetchDatasWalletInvestor(tokenInvestor.TokenInvestor)

	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(walletDatas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
}

func (wc *WalletController) FetchOperationsWallet(w http.ResponseWriter, r *http.Request) {
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

	var datasOperation entities.BodyAssetOperation

	err := json.NewDecoder(r.Body).Decode(&datasOperation)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errors.New("corpo da requisição inconsistente").Error())
		return
	}
	operationsDatas,err := wc.useCase.FetchOperationsWallet(datasOperation.TokenUser,datasOperation.OffSet)

	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(operationsDatas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
}
