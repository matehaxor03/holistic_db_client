package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type BranchNameCharacterWhitelist struct {
	GetBranchNameCharacterWhitelist func() (*json.Map)
	ValidateBranchName func(branch_name string) ([]error)
}

func NewBranchNameCharacterWhitelist() (*BranchNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidBranchNameCharacters()
	cache := make(map[string]interface{})

	x := BranchNameCharacterWhitelist {
		GetBranchNameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetValidBranchNameCharacters()
			return &v
		},
		ValidateBranchName: func(branch_name string) ([]error) {
			if _, found := cache[branch_name]; found {
				return nil
			}
			
			var errors []error
			if branch_name == "" {
				errors = append(errors, fmt.Errorf("branch_name is empty"))
			}

			parameters := json.NewMapValue()
			parameters.SetStringValue("value", branch_name)
			parameters.SetMap("values", &valid_characters)
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
		},
	}

	return &x
}
