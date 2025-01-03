package controller

import (
	"encoding/json"
	"errors"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/handleError"
	"jumpStart-backEnd/useCase/news"
	"net/http"
	"strconv"
)

type NewsController struct {
	useCase *news.NewsUseCase
}

func NewNewsController(useCaseC *news.NewsUseCase) *NewsController {
	return &NewsController{useCase: useCaseC}
}

func (nc *NewsController) FetchNews(w http.ResponseWriter, r *http.Request) {
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

	var assetOperation []entities.NewsStructure

	offsetStr := r.URL.Query().Get("offset")
	offset, errConvert := strconv.Atoi(offsetStr)

	if errConvert != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errConvert.Error())
		return
	}

	assetOperation,err := nc.useCase.FindAllNews(offset)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, err.Error())
		return
	}
	json.NewEncoder(w).Encode(assetOperation)

		
}

func (nc *NewsController) DeleteNews(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")

	code,token := getTokenJWT(r)
	if code != http.StatusOK {
		handleError.WriteHTTPStatus(w, code, token)
		return
	}

	var datasOperation entities.NewsDelete

	err := json.NewDecoder(r.Body).Decode(&datasOperation)
	if err != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, errors.New("corpo da requisição inconsistente").Error())
		return
	}
	datasOperation.TokenInvestor = token
	message := nc.useCase.DeleteNews(datasOperation.IdNews,datasOperation.TokenInvestor)
	if message != nil {
		handleError.WriteHTTPStatus(w, http.StatusNotAcceptable, message.Error())
		return
	}
	handleError.WriteHTTPStatus(w, http.StatusOK, "Noticia deletada com sucesso")

}