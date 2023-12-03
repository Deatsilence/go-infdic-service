package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type WordTranslate struct {
	gorm.Model
	Id       int    `json:"id"`
	WordEn   string `json:"word_en"`
	WordTr   string `json:"word_tr"`
	Type     string `json:"type"`
	Category string `json:"category"`
}

// func init()  {
// 	config.conne
// }
