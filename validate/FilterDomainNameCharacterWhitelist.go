package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type DomainNameCharacterWhitelist struct {
	GetDomainNameCharacterWhitelist func() (*json.Map)
	ValidateDomainName func(domain_name string) ([]error)
	GetValidateDomainNameFunc func() (*func(domain_name string) []error)
}

func NewDomainNameCharacterWhitelist() (*DomainNameCharacterWhitelist) {
	valid_characters := validation_constants.GetValidDomainNameCharacters()
	valid_words := validation_constants.GET_ALLOWED_DOMAIN_NAMES()
	cache := make(map[string]interface{})

	validateDomainName := func(domain_name string) ([]error) {
		if _, found := cache[domain_name]; found {
			return nil
		}
		
		var errors []error
		if domain_name == "" {
			errors = append(errors, fmt.Errorf("domain_name is empty"))
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", domain_name)
		parameters.SetMap("values", &valid_characters)
		parameters.SetStringValue("label", "Validator.ValidateDomainName")
		parameters.SetStringValue("data_type", "host.domain_name")
		whitelist_errors := validation_functions.WhitelistCharacters(parameters)
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		
		parameters.SetMap("values", &valid_words)
		whitelist_word_errors := validation_functions.WhiteListString(parameters)
		if whitelist_word_errors != nil {
			errors = append(errors, whitelist_word_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		cache[domain_name] = nil
		return nil
	}


	x := DomainNameCharacterWhitelist {
		GetDomainNameCharacterWhitelist: func() (*json.Map) {
			v := validation_constants.GetValidDomainNameCharacters()
			return &v
		},
		ValidateDomainName: func(domain_name string) ([]error) {
			return validateDomainName(domain_name)
		},
		GetValidateDomainNameFunc: func() (*func(domain_name string) []error) {
			function := validateDomainName
			return &function
		},
	}

	return &x
}
