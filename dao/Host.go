package dao

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	helper "github.com/matehaxor03/holistic_db_client/helper"
)

type Host struct {
	Validate      func() []error
	ToJSONString  func(json *strings.Builder) ([]error)
	GetHostName   func() (string, []error)
	GetPortNumber func() (string, []error)
}

func NewHost(host_name string, port_number string) (*Host, []error) {
	struct_type := "*dao.Host"

	data := json.NewMap()
	data.SetMap("[fields]", json.NewMap())
	data.SetMap("[schema]", json.NewMap())

	map_system_fields := json.NewMap()
	map_system_fields.SetStringValue("[host_name]", host_name)
	map_system_fields.SetStringValue("[port_number]", port_number)
	data.SetMap("[system_fields]", map_system_fields)

	map_system_schema := json.NewMap()

	map_host_name_schema := json.NewMap()
	map_host_name_schema.SetStringValue("type", "string")
	map_host_name_schema_filters := json.NewArray()
	map_host_name_schema_filter := json.NewMap()
	map_host_name_schema_filter.SetObjectForMap("values", validation_constants.GetValidDomainNameCharacters())
	map_host_name_schema_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
	map_host_name_schema_filters.AppendMap(map_host_name_schema_filter)
	map_host_name_schema.SetArray("filters", map_host_name_schema_filters)
	map_system_schema.SetMap("[host_name]", map_host_name_schema)

	map_port_number_schema := json.NewMap()
	map_port_number_schema.SetStringValue("type", "string")
	map_port_number_schema_filters := json.NewArray()
	map_port_number_schema_filter := json.NewMap()
	map_port_number_schema_filter.SetObjectForMap("values", validation_constants.GetValidPortNumberCharacters())
	map_port_number_schema_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
	map_port_number_schema_filters.AppendMap(map_port_number_schema_filter)
	map_port_number_schema.SetArray("filters", map_port_number_schema_filters)
	map_system_schema.SetMap("[port_number]", map_port_number_schema)

	data.SetMap("[system_schema]", map_system_schema)

	getData := func() *json.Map {
		return data
	}

	getHostName := func() (string, []error) {
		var errors []error
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[host_name]", "string")
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
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[port_number]", "string")
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

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &Host{
		Validate: func() []error {
			return validate()
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
		},
		GetHostName: func() (string, []error) {
			return getHostName()
		},
		GetPortNumber: func() (string, []error) {
			return getPortNumber()
		},
	}, nil
}
