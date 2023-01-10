package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type RepositoryAccountNameCharacterWhitelist struct {
	ValidateRepositoryAccountName func(respository_account_name string) ([]error)
	GetValidateRepositoryAccountNameFunc func() (*func(string) []error)
}

func NewRepositoryAccountNameCharacterWhitelist() (*RepositoryAccountNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidRepositoryAccountNameCharacters()
	cache := make(map[string]interface{})

	validateRepositoryAccountName := func(respository_account_name string) ([]error) {
		if _, found := cache[respository_account_name]; found {
			return nil
		}
		
		var errors []error
		if respository_account_name == "" {
			errors = append(errors, fmt.Errorf("respository_account_name is empty"))
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", respository_account_name)
		parameters.SetMap("values", &valid_characters)
		parameters.SetStringValue("label", "Validator.ValidateRepostoryAccountName")
		parameters.SetStringValue("data_type", "git.repository_account_name")
		whitelist_errors := validation_functions.WhitelistCharacters(parameters)
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		cache[respository_account_name] = nil
		return nil
	}

	x := RepositoryAccountNameCharacterWhitelist {
		ValidateRepositoryAccountName: func(respository_account_name string) ([]error) {
			return validateRepositoryAccountName(respository_account_name)
		},
		GetValidateRepositoryAccountNameFunc: func() (*func(respository_account_name string) []error) {
			function := validateRepositoryAccountName
			return &function
		},
	}

	return &x
}
