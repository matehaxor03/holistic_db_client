package validate

import (
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type DatabaseReservedWordsBlackList struct {
	ValidateDatabaseReservedWord func(value string) ([]error)
	GetValidateDatabaseReservedWordFunc func() (*func(string) []error)
}

func NewDatabaseReservedWordsBlackList() (*DatabaseReservedWordsBlackList) {
	database_reserved_words := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()
	cache := make(map[string]interface{})

	validateDatabaseReservedWord := func(value string) ([]error) {
		if _, found := cache[value]; found {
			return nil
		}
		
		var errors []error
		if value == "" {
			errors = append(errors, fmt.Errorf("value is empty"))
		}

		whitelist_errors := validation_functions.BlackListStringToUpper(database_reserved_words, value, "Validator.ValidateDatabaseReservedWord", "database.cross_cutting_field_value" )
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		cache[value] = nil
		return nil
	}

	x := DatabaseReservedWordsBlackList{
		ValidateDatabaseReservedWord: func(value string) ([]error) {
			return validateDatabaseReservedWord(value)
		},
		GetValidateDatabaseReservedWordFunc: func() (*func(value string) []error) {
			function := validateDatabaseReservedWord
			return &function
		},
	}

	return &x
}
