package controller

import (
	"encoding/json"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/handleError"
	"jumpStart-backEnd/useCase/investor"
	"net/http"
)

type InvestorController struct {
	useCase *investor.InvestorUseCase
}

func NewInvestorController(useCase *investor.InvestorUseCase) *InvestorController {
	return &InvestorController{useCase: useCase}
}

func (ic *InvestorController) CreateInvestor(w http.ResponseWriter, r *http.Request) {
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

	var investor entities.InvestorInsert

	err := json.NewDecoder(r.Body).Decode(&investor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	error := ic.useCase.CreateInvestor(investor)
	if error != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, error.Error())
		return
	}
	handleError.WriteHTTPStatus(w, http.StatusOK, "Investidor criado com sucesso")

}

func (ic *InvestorController) Login(w http.ResponseWriter, r *http.Request) {
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

	var investor entities.LoginInvestor

	err := json.NewDecoder(r.Body).Decode(&investor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token,error := ic.useCase.LoginInvestor(investor)
	if error != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, error.Error())
		return
	}

	json.NewEncoder(w).Encode(token)

}