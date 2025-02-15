package controller

import (
	"encoding/json"
	"jumpStart-backEnd/entities"
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

func (lac *ListAssetController) ListAssets(w http.ResponseWriter, r *http.Request) {
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
	code,token := getTokenJWT(r)
	if code != http.StatusOK {
		handleError.WriteHTTPStatus(w, code, token)
		return
	}


	response,err := lac.useCase.GetListAssets(token)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response)

		
}

func (lac *ListAssetController) UpdateUrlImage(w http.ResponseWriter, r *http.Request) {
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

	var datas entities.UpdateUrlImage
	err := json.NewDecoder(r.Body).Decode(&datas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	errUpdate := lac.useCase.UpdateUrlImage(token,datas)
	if errUpdate != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errUpdate.Error())
		return
	}

	handleError.WriteHTTPStatus(w, http.StatusOK, "Imagem atualizada com sucesso")

		
}


func (lac *ListAssetController) CreateNewAsset(w http.ResponseWriter, r *http.Request) {
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

	var datas entities.NewAsset
	err := json.NewDecoder(r.Body).Decode(&datas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	errUpdate := lac.useCase.AddNewAsset(token,datas)
	if errUpdate != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errUpdate.Error())
		return
	}

	handleError.WriteHTTPStatus(w, http.StatusOK, "ativo adicionado com sucesso")

		
}

func (lac *ListAssetController) GetHistoryCrypto(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	crypto := r.URL.Query().Get("crypto")

	w.Header().Set("Content-Type", "application/json")

	response,err := lac.useCase.GetHistoryCrypto(crypto)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response)

		
}


