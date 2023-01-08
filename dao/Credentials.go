package dao

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	common "github.com/matehaxor03/holistic_common/common"
	validate "github.com/matehaxor03/holistic_db_client/validate"

	"fmt"
)

type Credentials struct {
	Validate     func() []error
	GetUsername  func() (string, []error)
	GetPassword  func() (string, []error)
	ToJSONString func(json *strings.Builder) ([]error)
	Clone        func() *Credentials
}

func newCredentials(verify *validate.Validator, username string, password string) (*Credentials, []error) {
	struct_type := "*dao.Credentials"


	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[username]", username)
	map_system_fields.SetObjectForMap("[password]", password)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.NewMapValue()
	
	map_username_schema := json.NewMapValue()
	map_username_schema.SetStringValue("type", "*string")
	map_username_schema.SetIntValue("min_length", 1)
	array_username_filters := json.NewArrayValue()
	map_username_filter := json.NewMapValue()
	map_username_filter.SetObjectForMap("values", verify.GetUsernameCharacterWhitelist())
	map_username_filter.SetObjectForMap("function",  validation_functions.GetWhitelistCharactersFunc())
	array_username_filters.AppendMapValue(map_username_filter)
	map_username_schema.SetArrayValue("filters", array_username_filters)
	map_system_schema.SetMapValue("[username]", map_username_schema)

	map_password_schema := json.NewMapValue()
	map_password_schema.SetStringValue("type", "*string")
	map_password_schema.SetIntValue("min_length", 1)
	map_system_schema.SetMapValue("[password]", map_password_schema)
	
	data.SetMapValue("[system_schema]", map_system_schema)

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	getUsername := func() (string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[username]", "string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("username is nil"))
		}

		if len(errors) > 0 {
			return "", errors
		}
		return temp_value.(string), nil
	}

	getPassword := func() (string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[password]", "string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			return "", nil
		}

		if len(errors) > 0 {
			return "", errors
		}
		return temp_value.(string), nil
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &Credentials{
		Validate: func() []error {
			return validate()
		},
		GetUsername: func() (string, []error) {
			return getUsername() 
		},
		GetPassword: func() (string, []error) {
			return getPassword()
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
		},
	}, nil
}
