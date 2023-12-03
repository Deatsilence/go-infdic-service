package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Deatsilence/go-infdic-service/pkg/models"
)

func GetDictionary(w http.ResponseWriter, r *http.Request) {
	dictionary := models.GetAllWords()
	res, err := json.Marshal(dictionary)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
