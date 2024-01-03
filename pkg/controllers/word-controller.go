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
	fmt.Println(len(dictionary))

	res, err := json.Marshal(dictionary)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func hasAudio(clearWords *[]models.WordTranslate) []models.WordTranslate {
// 	if clearWords == nil {
// 		return nil
// 	}

// 	var wg sync.WaitGroup
// 	var wordsWithAudio []models.WordTranslate
// 	var errors []error

// 	for _, word := range *clearWords {
// 		wg.Add(1)

// 		go func(word models.WordTranslate) {
// 			defer wg.Done()

// 			response, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word.WordEn)
// 			if err != nil {
// 				errors = append(errors, err)
// 				return
// 			}

// 			if response.StatusCode == http.StatusOK {
// 				wordsWithAudio = append(wordsWithAudio, word)
// 			}
// 		}(word)
// 	}

// 	wg.Wait()

// 	if len(errors) > 0 {
// 		return nil
// 	}

// 	return wordsWithAudio
// }
