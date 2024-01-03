package models

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Deatsilence/go-infdic-service/pkg/config"
	"github.com/Deatsilence/go-infdic-service/pkg/utils"
)

var db *gorm.DB

type WordTranslate struct {
	Id       int    `json:"id"`
	WordEn   string `json:"word_en"`
	WordTr   string `json:"word_tr"`
	Type     string `json:"type"`
	Category string `json:"category"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetMeansOfTheWord(word string) []WordTranslate {
	var words []WordTranslate
	fmt.Println("word:", word)
	utils.GetMeansOfAWordQuery(db, word).Find(&words)

	return words
}
