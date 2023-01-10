package validate

import (
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
)

type CollateWordWhitelist struct {
	ValidateCollate func(collate string) ([]error)
	GetValidateCollateFunc func() (*func(string) []error)
}

func NewCollateWordWhitelist() (*CollateWordWhitelist) {
	valid_words := validation_constants.GET_COLLATES()
	cache := make(map[string]interface{})

	validateCollate := func(collate string) ([]error) {
		if _, found := cache[collate]; found {
			return nil
		}

		whitelist_errors := validation_functions.WhiteListString(valid_words, collate, "Validator.ValidateCollate", "database.collate")
		if whitelist_errors != nil {
			return whitelist_errors
		}

		cache[collate] = nil
		return nil
	}

	x := CollateWordWhitelist {
		ValidateCollate: func(collate string) ([]error) {
			return validateCollate(collate)
		},
		GetValidateCollateFunc: func() (*func(collate string) []error) {
			function := validateCollate
			return &function
		},
	}
	return &x
}
