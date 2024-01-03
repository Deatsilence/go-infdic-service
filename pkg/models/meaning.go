package models

type Meaning struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []Definition  `json:"definitions"`
	Synonyms     []string      `json:"synonyms"`
	Antonyms     []interface{} `json:"antonyms"`
}
