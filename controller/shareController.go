package controller

import (
	"encoding/json"
	"jumpStart-backEnd/useCase"
	"net/http"
	"strconv"
)

type ShareController struct {
	useCase *usecase.ShareUseCase
}

func NewShareController(useCase *usecase.ShareUseCase) *ShareController {
	return &ShareController{useCase: useCase}
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
func (c *ShareController) GetTodaySharesJSON(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
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

func (c *ShareController) GetSharesSpecifyOffSet(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	offsetStr := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shares, err := c.useCase.ListSharesBasedOffSet(offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(shares); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
