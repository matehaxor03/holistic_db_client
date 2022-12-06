package class

import (
	json "github.com/matehaxor03/holistic_json/json"
)

type DatabaseReservedWords struct {
	GetDatabaseReservedWords func() (*json.Map)
}

func newDatabaseReservedWords() (*DatabaseReservedWords) {
	database_reserved_words := GetMySQLKeywordsAndReservedWordsInvalidWords()

	x := DatabaseReservedWords{
		GetDatabaseReservedWords: func() (*json.Map) {
			return &database_reserved_words
		},
	}

	return &x
}
