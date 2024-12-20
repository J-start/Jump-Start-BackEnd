package controller

import (
	"encoding/json"
    "jumpStart-backEnd/useCase/operation"

	"jumpStart-backEnd/handleError"
	"net/http"
)

type OperationAssetController struct {
	useCase *operation.OperationAssetUseCase
}

func NewOperationAssetController (useCaseC *operation.OperationAssetUseCase) *OperationAssetController {
	return &OperationAssetController{useCase: useCaseC}
}

func (oac *OperationAssetController) FetchHistoryOperationInvestor(w http.ResponseWriter, r *http.Request) {
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

	//TODO CREATE LOGIC TO OBTAIN ID INVESTOR AND OFFSET FROM body
	response,err := oac.useCase.FetchAssetHistoryByInvestor(1,0)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response)

		
}

