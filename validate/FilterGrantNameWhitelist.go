package validate

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	"fmt"
)

type GrantNameWhitelist struct {
	ValidateGrant func(frant string) ([]error)
	GetValidateGrantFunc func() (*func(grant string) []error)
	GetValidateDatabaseNameFilterAllFunc func() (*func(string) []error)
	GetValidateTableNameFilterAllFunc func() (*func(string) []error)
}

func NewGrantNameWhitelist() (*GrantNameWhitelist) {
	valid_words := validation_constants.GET_ALLOWED_GRANTS()
	cache := make(map[string]interface{})

	validateGrant := func(grant string) ([]error) {
		if _, found := cache[grant]; found {
			return nil
		}

		parameters := json.NewMapValue()
		parameters.SetStringValue("value", grant)
		parameters.SetObjectForMap("values", &valid_words)
		parameters.SetStringValue("label", "Validator.validateGrant")
		parameters.SetStringValue("data_type", "dao.Grant.grant")
		whitelist_errors := validation_functions.WhiteListString(parameters)
		if whitelist_errors != nil {
			return whitelist_errors
		}

		cache[grant] = nil
		return nil
	}

	validateDatabaseNameFilterAll :=  func(database_name_filter string) ([]error) {
		if database_name_filter == "*" {
			return nil
		}
		var errors []error
		errors = append(errors, fmt.Errorf("database_name_filter not supported"))
		return errors
	}

	validateTableNameFilterAll :=  func(table_name_filter string) ([]error) {
		if table_name_filter == "*" {
			return nil
		}
		var errors []error
		errors = append(errors, fmt.Errorf("table_name_filter not supported"))
		return errors
	}

	x := GrantNameWhitelist {
		ValidateGrant: func(grant string) ([]error) {
			return validateGrant(grant)
		},
		GetValidateGrantFunc: func() (*func(grant string) []error) {
			function := validateGrant
			return &function
		},
		GetValidateDatabaseNameFilterAllFunc: func() (*func(database_name_filter string) []error) {
			function := validateDatabaseNameFilterAll
			return &function
		},
		GetValidateTableNameFilterAllFunc: func() (*func(table_name_filter string) []error) {
			function := validateTableNameFilterAll
			return &function
		},
	}

	return &x
}
