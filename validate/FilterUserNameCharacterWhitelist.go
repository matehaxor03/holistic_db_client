package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type UsernameCharacterWhitelist struct {
	GetUsernameCharacterWhitelist func() (*json.Map)
	ValidateUsername func(username string) ([]error)
	GetValidateUsernameFunc func() (*func(string) []error)
}

func NewUsernameCharacterWhitelist() (*UsernameCharacterWhitelist) {
	valid_username_characters := validation_constants.GetValidUsernameCharacters()
	valid_usernames_cache := make(map[string]interface{})

	validateUsername := func(username string) ([]error) {
		if _, found := valid_usernames_cache[username]; found {
			return nil
		}
		
		var errors []error
		if username == "" {
			errors = append(errors, fmt.Errorf("username is empty"))
		}

		if len(username) < 2 {
			errors = append(errors, fmt.Errorf("username is too short must be at least 2 characters"))
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", username)
		parameters.SetMap("values", &valid_username_characters)
		parameters.SetStringValue("label", "Validator.ValidateUsername")
		parameters.SetStringValue("data_type", "dao.User.username")
		whitelist_errors := validation_functions.WhitelistCharacters(parameters)
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		valid_usernames_cache[username] = nil
		return nil
	}

	x := UsernameCharacterWhitelist {
		GetUsernameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetValidUsernameCharacters()
			return &v
		},
		ValidateUsername: func(username string) ([]error) {
			return validateUsername(username)
		},
		GetValidateUsernameFunc: func() (*func(string) []error) {
			function := validateUsername
			return &function
		},
	}

	return &x
}
