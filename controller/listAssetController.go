package controller

import (
	"encoding/json"
    "jumpStart-backEnd/useCase/listasset"

	"jumpStart-backEnd/handleError"
	"net/http"
)

type ListAssetController struct {
	useCase *listasset.ListAssetUseCase
}

func NewListAssetController(useCase *listasset.ListAssetUseCase) *ListAssetController {
	return &ListAssetController{useCase: useCase}
}

func (lac *ListAssetController) ListAsset(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")


	asset := r.URL.Query().Get("type")


	response,err := lac.useCase.ListAssetByType(asset)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response)

		
}

func (lac *ListAssetController) ListAssetRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")


	asset := r.URL.Query().Get("type")


	response,err := lac.useCase.ListAssetRequest(asset)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response)

		
}

