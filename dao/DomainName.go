package dao

import (
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

type DomainName struct {
	Validate      func() []error
	GetDomainName func() (string)
}

func NewDomainName(verify *validate.Validator, domain_name string) (*DomainName, []error) {

	validate := func() []error {
		var errors []error
		
		if domain_name_errors := verify.ValidateDomainName(domain_name); domain_name_errors != nil {
			errors = append(errors, domain_name_errors...)
		}
		if len(errors) > 0 {
			return errors
		}
		
		return nil
	}

	getDomainName := func() (string) {
		return domain_name
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &DomainName{
		Validate: func() []error {
			return validate()
		},
		GetDomainName: func() (string) {
			return getDomainName()
		},
	}, nil
}
