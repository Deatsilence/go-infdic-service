package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Deatsilence/go-infdic-service/pkg/models"
	"github.com/gorilla/mux"
)

func GetDictionary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	dictionary := models.GetMeansOfTheWord(word)

	res, err := json.Marshal(dictionary)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func HandleGetWordDefinition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	word := vars["word"]

	resp, err := http.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word))
	if err != nil {
		fmt.Println(err)
		return
	}

	var wordDefinition []models.WordInformations
	if err := json.NewDecoder(resp.Body).Decode(&wordDefinition); err != nil {
		fmt.Println(err)
		return
	}
	res, err := json.Marshal(wordDefinition)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
