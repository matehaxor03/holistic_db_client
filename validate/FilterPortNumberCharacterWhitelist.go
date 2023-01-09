package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type PortNumberCharacterWhitelist struct {
	GetPortNumberCharacterWhitelist func() (*json.Map)
	ValidatePortNumber func(port_number string) ([]error) 
	GetValidatePortNumberFunc func() (*func(string) []error)
}

func NewPortNumberCharacterWhitelist() (*PortNumberCharacterWhitelist) {
	valid_characters := validation_constants.GetValidPortNumberCharacters()
	cache := make(map[string]interface{})

	validatePortNumber := func(port_number string) ([]error) {
		if _, found := cache[port_number]; found {
			return nil
		}
		
		var errors []error
		if port_number == "" {
			errors = append(errors, fmt.Errorf("port_number is empty"))
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", port_number)
		parameters.SetMap("values", &valid_characters)
		parameters.SetStringValue("label", "Validator.ValidatePortNumber")
		parameters.SetStringValue("data_type", "host.port_number")
		whitelist_errors := validation_functions.WhitelistCharacters(parameters)
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		//todo: check port number range

		if len(errors) > 0 {
			return errors
		}

		cache[port_number] = nil
		return nil
	}

	x := PortNumberCharacterWhitelist {
		GetPortNumberCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetValidPortNumberCharacters()
			return &v
		},
		ValidatePortNumber: func(port_number string) ([]error) {
			return validatePortNumber(port_number)
		},
		GetValidatePortNumberFunc: func() (*func(port_number string) []error) {
			function := validatePortNumber
			return &function
		},
	}

	return &x
}
