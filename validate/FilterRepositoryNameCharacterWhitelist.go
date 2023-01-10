package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type RepositoryNameCharacterWhitelist struct {
	GetValidateRepositoryNameFunc func() (*func(string) []error) 
	ValidateRepositoryName func(respository_name string) ([]error)
}

func NewRepositoryNameCharacterWhitelist() (*RepositoryNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidRepositoryNameCharacters()
	cache := make(map[string]interface{})

	validateRepositoryName := func(respository_name string) ([]error) {
		if _, found := cache[respository_name]; found {
			return nil
		}
		
		var errors []error
		if respository_name == "" {
			errors = append(errors, fmt.Errorf("respository_name is empty"))
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", respository_name)
		parameters.SetObjectForMap("values", &valid_characters)
		parameters.SetStringValue("label", "Validator.ValidateRepostoryName")
		parameters.SetStringValue("data_type", "git.repository_name")
		whitelist_errors := validation_functions.WhitelistCharacters(parameters)
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		cache[respository_name] = nil
		return nil
	}


	x := RepositoryNameCharacterWhitelist {
		ValidateRepositoryName: func(respository_name string) ([]error) {
			return validateRepositoryName(respository_name)
		},
		GetValidateRepositoryNameFunc: func() (*func(respository_name string) []error) {
			function := validateRepositoryName
			return &function
		},
	}

	return &x
}
