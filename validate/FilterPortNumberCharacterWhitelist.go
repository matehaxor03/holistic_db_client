package validate

import (
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type PortNumberCharacterWhitelist struct {
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

		whitelist_errors := validation_functions.WhitelistCharacters(valid_characters, port_number, "Validator.ValidatePortNumber", "host.port_number")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		cache[port_number] = nil
		return nil
	}

	x := PortNumberCharacterWhitelist {
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
