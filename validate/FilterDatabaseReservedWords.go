package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
	"strings"
)

type DatabaseReservedWordsBlackList struct {
	GetDatabaseReservedWordsBlackList func() (*json.Map)
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

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", strings.ToUpper(value))
		parameters.SetMap("values", &database_reserved_words)
		parameters.SetStringValue("label", "Validator.ValidateDatabaseReservedWord")
		parameters.SetStringValue("data_type", "database.cross_cutting_field_value")
		whitelist_errors := validation_functions.BlackListStringToUpper(parameters)
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
		GetDatabaseReservedWordsBlackList: func() (*json.Map) {
			v := validation_constants.GetMySQLKeywordsAndReservedWordsInvalidWords()
			return &v
		},
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
