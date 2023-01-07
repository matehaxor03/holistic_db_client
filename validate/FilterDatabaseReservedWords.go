package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
)

type DatabaseReservedWordsBlackList struct {
	GetDatabaseReservedWordsBlackList func() (*json.Map)
}

func NewDatabaseReservedWordsBlackList() (*DatabaseReservedWordsBlackList) {
	database_reserved_words := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()

	x := DatabaseReservedWordsBlackList{
		GetDatabaseReservedWordsBlackList: func() (*json.Map) {
			return &database_reserved_words
		},
	}

	return &x
}
