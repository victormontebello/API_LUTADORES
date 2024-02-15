package handlers

import (
	"encoding/json"
	"log"
	"lutadores-modulo/models"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var lutador models.Lutador

	err := json.NewDecoder(r.Body).Decode(&lutador)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(lutador)

	var resp map[string]any

	if err != nil {
		log.Println(err)
		resp = map[string]any{"erro": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		resp = map[string]any{"id": id}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	return
}
