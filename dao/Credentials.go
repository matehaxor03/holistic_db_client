package dao

import (
	validate "github.com/matehaxor03/holistic_validator/validate"
)

type Credentials struct {
	Validate     func() []error
	GetUsername  func() (string)
	GetPassword  func() (string)
	Clone        func() *Credentials
}

func newCredentials(verify *validate.Validator, username string, password string) (*Credentials, []error) {
	
	validate := func() []error {
		var errors []error
		if username_errors := verify.ValidateUsername(username); username_errors != nil {
			errors = append(errors, username_errors...)
		}
		
		if len(errors) > 0 {
			return errors
		}

		return nil
	}

	getUsername := func() (string) {
		return username
	}

	getPassword := func() (string) {
		return password
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &Credentials{
		Validate: func() []error {
			return validate()
		},
		GetUsername: func() (string) {
			return getUsername() 
		},
		GetPassword: func() (string) {
			return getPassword()
		},
	}, nil
}
