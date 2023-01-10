package validate

import (
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type DomainNameCharacterWhitelist struct {
	ValidateDomainName func(domain_name string) ([]error)
	GetValidateDomainNameFunc func() (*func(string) []error)
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

		whitelist_errors := validation_functions.WhitelistCharacters(valid_characters, domain_name,  "Validator.ValidateDomainName",  "host.domain_name")
		if whitelist_errors != nil {
			errors = append(errors, whitelist_errors...)
		}

		
		whitelist_word_errors := validation_functions.WhiteListString(valid_words, domain_name,  "Validator.ValidateDomainName",  "host.domain_name")
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
		ValidateDomainName: func(domain_name string) ([]error) {
			return validateDomainName(domain_name)
		},
		GetValidateDomainNameFunc: func() (*func(string) []error) {
			function := validateDomainName
			return &function
		},
	}

	return &x
}
