package controller

import (
	"encoding/json"
	"jumpStart-backEnd/useCase"
	"net/http"
)


type ShareController struct {
	useCase *usecase.ShareUseCase
}

func NewShareController(useCase *usecase.ShareUseCase) *ShareController {
	return &ShareController{useCase: useCase}
}

func (c *ShareController) GetTodaySharesJSON(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")

	shares, err := c.useCase.FindAllShares()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	if err := json.NewEncoder(w).Encode(shares); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}