package controller

import (
	"encoding/json"
	"errors"
	"jumpStart-backEnd/entities"
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

	var assetOperation entities.BodyAssetOperation

	err := json.NewDecoder(r.Body).Decode(&assetOperation)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errors.New("corpo da requisição inconsistente").Error())
		return
	}

	response,err := oac.useCase.FetchAssetHistoryByInvestor(assetOperation.TokenUser,assetOperation.OffSet)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response)

		
}

