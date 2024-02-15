package handlers

import (
	"encoding/json"
	"lutadores-modulo/models"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	lutadores, err := models.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lutadores)

	return
}