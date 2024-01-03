package utils

import (
	"github.com/jinzhu/gorm"
)

// GetMeansOfAWordQuery returns a query that finds the means of a word.
func GetMeansOfAWordQuery(db *gorm.DB, word string) *gorm.DB {
	var englishIds []int
	db.Table("english").Select("id").Where("word LIKE ?", word).Pluck("id", &englishIds)

	return db.Table("translate").Select(
		"translate.id as id, english.word as word_en, turkish.word as word_tr, type.name as type, category.name as category",
	).Joins(
		"INNER JOIN english ON translate.english_id = english.id",
	).Joins(
		"INNER JOIN turkish ON translate.turkish_id = turkish.id",
	).Joins(
		"INNER JOIN type ON translate.type_id = type.id",
	).Joins(
		"INNER JOIN category ON translate.category_id = category.id",
	).Where(
		"translate.english_id IN (?)", englishIds,
	)
}
