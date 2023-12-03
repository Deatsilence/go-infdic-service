package utils

import (
	"github.com/jinzhu/gorm"
)

func GetAllWordsQuery(db *gorm.DB) *gorm.DB {
	return db.Table("translate").Select("translate.id as id, english.word as word_en, turkish.word as word_tr, type.name as type, category.name as category").Joins("JOIN english ON translate.english_id = english.id").Joins("JOIN turkish ON translate.turkish_id = turkish.id").Joins("JOIN type ON translate.type_id = type.id").Joins("JOIN category ON translate.category_id = category.id").Order("word_en ASC").Limit(20)
}
