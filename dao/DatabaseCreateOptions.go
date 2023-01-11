package dao

import (
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

type DatabaseCreateOptions struct {
	GetCharacterSet func() (*string)
	GetCollate 		func() (*string)
	Validate     func() []error
}

func newDatabaseCreateOptions(verify *validate.Validator, character_set *string, collate *string) (*DatabaseCreateOptions, []error) {
	
	validate := func() []error {
		var errors []error
		if character_set != nil {
			if character_set_errors := verify.ValidateCharacterSet(*character_set); character_set_errors != nil {
				errors = append(errors, character_set_errors...)
			}
		}

		if collate != nil {
			if collate_errors := verify.ValidateCollate(*collate); collate_errors != nil {
				errors = append(errors, collate_errors...)
			}
		}

		if len(errors) > 0 {
			return errors
		}
		return nil
	}

	get_character_set := func() (*string) {
		return character_set
	}

	get_collate := func() (*string) {
		return collate
	}

	validate_errors := validate()

	if len(validate_errors) > 0 {
		return nil, validate_errors
	}

	return &DatabaseCreateOptions{
		GetCharacterSet: func() (*string) {
			return get_character_set()
		},
		GetCollate: func() (*string) {
			return get_collate()
		},
		Validate: func() []error {
			return validate()
		},
	}, nil
}
