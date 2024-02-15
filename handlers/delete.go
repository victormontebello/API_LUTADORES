package handlers

import (
	"encoding/json"
	"lutadores-modulo/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rows < 1 {
		http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{"id": id}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	return
}