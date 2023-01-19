package dao

import (
	validate "github.com/matehaxor03/holistic_validator/validate"
)

type Host struct {
	Validate      func() []error
	GetHostName   func() (string)
	GetPortNumber func() (string)
}

func newHost(verify *validate.Validator, host_name string, port_number string) (*Host, []error) {
	
	getHostName := func() (string) {
		return host_name
	}

	getPortNumber := func() (string) {
		return port_number
	}

	validate := func() []error {
		var errors []error
		if hostname_errors := verify.ValidateDomainName(host_name); hostname_errors != nil {
			errors = append(errors, hostname_errors...)
		}

		if port_number_errors := verify.ValidatePortNumber(port_number); port_number_errors != nil {
			errors = append(errors, port_number_errors...)
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	validate_errors := validate()

	if validate_errors != nil {
		return nil, validate_errors
	}

	return &Host{
		Validate: func() []error {
			return validate()
		},
		GetHostName: func() (string) {
			return getHostName()
		},
		GetPortNumber: func() (string) {
			return getPortNumber()
		},
	}, nil
}
