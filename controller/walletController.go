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

	code,token := getTokenJWT(r)
	if code != http.StatusOK {
		handleError.WriteHTTPStatus(w, code, token)
		return
	}

	walletDatas,err := wc.useCase.FetchDatasWalletInvestor(token)

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

	code,token := getTokenJWT(r)
	if code != http.StatusOK {
		handleError.WriteHTTPStatus(w, code, token)
		return
	}

	var datasOperation entities.BodyAssetOperation

	err := json.NewDecoder(r.Body).Decode(&datasOperation)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errors.New("corpo da requisição inconsistente").Error())
		return
	}
	datasOperation.TokenInvestor = token
	operationsDatas,err := wc.useCase.FetchOperationsWallet(datasOperation.TokenInvestor,datasOperation.OffSet)

	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(operationsDatas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (wc *WalletController) WithDraw(w http.ResponseWriter, r *http.Request) {
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
	var datasOperation entities.WalletOperationRequest

	err := json.NewDecoder(r.Body).Decode(&datasOperation)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errors.New("corpo da requisição inconsistente").Error())
		return
	}
	datasOperation.TokenInvestor = token
	code,message := wc.useCase.WithDraw(datasOperation)

	handleError.WriteHTTPStatus(w, code, message)

}

func (wc *WalletController) Deposit(w http.ResponseWriter, r *http.Request) {
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

	var datasOperation entities.WalletOperationRequest

	err := json.NewDecoder(r.Body).Decode(&datasOperation)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errors.New("corpo da requisição inconsistente").Error())
		return
	}
	datasOperation.TokenInvestor = token
	code,message := wc.useCase.Deposit(datasOperation)

	handleError.WriteHTTPStatus(w, code, message)

}
