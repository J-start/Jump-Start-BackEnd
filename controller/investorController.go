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


func (ic *InvestorController) SendCodeEmailRecoverPassword(w http.ResponseWriter, r *http.Request) {
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

	var investor entities.SendCodeEmail

	err := json.NewDecoder(r.Body).Decode(&investor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	errA := ic.useCase.SendCodeToRecoverPassword(investor.Email)
	if errA != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, errA.Error())
		return
	}

	handleError.WriteHTTPStatus(w, http.StatusOK, "Código enviado com sucesso")

}

func (ic *InvestorController) VerifyCodeEmail(w http.ResponseWriter, r *http.Request) {
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

	var investor entities.CodeChangePassword

	err := json.NewDecoder(r.Body).Decode(&investor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	errB := ic.useCase.VerifyCode(investor.Email,investor.Code,investor.NewPassword)
	if errB != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, errB.Error())
		return
	}

	handleError.WriteHTTPStatus(w, http.StatusOK, "Sucesso!")

}

func (ic *InvestorController) GetNameAndBalance(w http.ResponseWriter, r *http.Request) {
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

	datas,errB := ic.useCase.NameAndBalanceInvestor(token)
	if errB != nil {
		handleError.WriteHTTPStatus(w, http.StatusBadRequest, errB.Error())
		return
	}
	json.NewEncoder(w).Encode(datas)

}