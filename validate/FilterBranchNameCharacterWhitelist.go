package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type BranchNameCharacterWhitelist struct {
	ValidateBranchName func(branch_name string) ([]error)
	GetValidateBranchNameFunc func() (*func(string) []error)
}

func NewBranchNameCharacterWhitelist() (*BranchNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidBranchNameCharacters()
	cache := make(map[string]interface{})

	validateBranchName := func(branch_name string) ([]error) {
		if _, found := cache[branch_name]; found {
			return nil
		}
		
		var errors []error
		if branch_name == "" {
			errors = append(errors, fmt.Errorf("branch_name is empty"))
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", branch_name)
		parameters.SetObjectForMap("values", &valid_characters)
		parameters.SetStringValue("label", "Validator.ValidateBranchName")
		parameters.SetStringValue("data_type", "git.branch_name")
		whitelist_errors := validation_functions.WhitelistCharacters(parameters)
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		cache[branch_name] = nil
		return nil
	}

	x := BranchNameCharacterWhitelist {
		ValidateBranchName: func(branch_name string) ([]error) {
			return validateBranchName(branch_name)
		},
		GetValidateBranchNameFunc: func() (*func(branch_name string) []error) {
			function := validateBranchName
			return &function
		},
	}

	return &x
}
