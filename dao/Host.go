package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	helper "github.com/matehaxor03/holistic_db_client/helper"
)

type Host struct {
	Validate      func() []error
	GetHostName   func() (string, []error)
	GetPortNumber func() (string, []error)
}

func newHost(verify *validate.Validator, host_name string, port_number string) (*Host, []error) {
	data := json.NewMap()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetString("[host_name]", &host_name)
	map_system_fields.SetString("[port_number]", &port_number)
	data.SetMapValue("[system_fields]", map_system_fields)

	map_system_schema := json.NewMapValue()

	map_host_name_schema := json.NewMapValue()
	map_host_name_schema.SetStringValue("type", "string")
	map_host_name_schema_filters := json.NewArrayValue()
	map_host_name_schema_filter := json.NewMapValue()
	map_host_name_schema_filter.SetObjectForMap("function", verify.GetValidateDomainNameFunc())
	map_host_name_schema_filters.AppendMapValue(map_host_name_schema_filter)
	map_host_name_schema.SetArrayValue("filters", map_host_name_schema_filters)
	map_system_schema.SetMapValue("[host_name]", map_host_name_schema)

	map_port_number_schema := json.NewMapValue()
	map_port_number_schema.SetStringValue("type", "string")
	map_port_number_schema_filters := json.NewArrayValue()
	map_port_number_schema_filter := json.NewMapValue()
	map_port_number_schema_filter.SetObjectForMap("function", verify.GetValidatePortNumberFunc())
	map_port_number_schema_filters.AppendMapValue(map_port_number_schema_filter)
	map_port_number_schema.SetArrayValue("filters", map_port_number_schema_filters)
	map_system_schema.SetMapValue("[port_number]", map_port_number_schema)

	data.SetMapValue("[system_schema]", map_system_schema)

	getData := func() *json.Map {
		return data
	}

	getHostName := func() (string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]",  "[host_name]", "string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} 

		if len(errors) > 0 {
			return "", errors
		}
		return temp_value.(string), nil
	}

	getPortNumber := func() (string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[port_number]", "string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} 

		if len(errors) > 0 {
			return "", errors
		}
		return temp_value.(string), nil
	}

	validate := func() []error {
		return ValidateData(getData(), "Host")
	}

	validate_errors := validate()

	if validate_errors != nil {
		return nil, validate_errors
	}

	return &Host{
		Validate: func() []error {
			return validate()
		},
		GetHostName: func() (string, []error) {
			return getHostName()
		},
		GetPortNumber: func() (string, []error) {
			return getPortNumber()
		},
	}, nil
}
