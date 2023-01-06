package dao

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
	validation_constants "github.com/matehaxor03/holistic_db_client/validation_constants"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	//"fmt"
)

type Host struct {
	Validate      func() []error
	ToJSONString  func(json *strings.Builder) ([]error)
	GetHostName   func() (string, []error)
	GetPortNumber func() (string, []error)
}

func NewHost(host_name string, port_number string) (*Host, []error) {
	struct_type := "*dao.Host"

	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[host_name]", host_name)
	map_system_fields.SetObjectForMap("[port_number]", port_number)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.NewMapValue()

	// Start host_name
	map_host_name_schema := json.NewMapValue()
	map_host_name_schema.SetStringValue("type", "string")


	//fmt.Println(fmt.Sprintf("hello"))
	//fmt.Println(fmt.Sprintf("%s", validation_functions.GetWhitelistCharactersFunc()))
	//fmt.Println(fmt.Sprintf("%T", validation_functions.GetWhitelistCharactersFunc()))

	map_host_name_schema_filters := json.NewArrayValue()
	map_host_name_schema_filter := json.NewMapValue()
	map_host_name_schema_filter.SetObjectForMap("values", validation_constants.GetValidDomainNameCharacters())
	map_host_name_schema_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
	map_host_name_schema_filters.AppendMapValue(map_host_name_schema_filter)
	map_host_name_schema.SetArrayValue("filters", map_host_name_schema_filters)
	map_system_schema.SetMapValue("[host_name]", map_host_name_schema)
	// End host_name

	// Start port_number
	map_port_number_schema := json.NewMapValue()
	map_port_number_schema.SetStringValue("type", "string")

	map_port_number_schema_filters := json.NewArrayValue()
	map_port_number_schema_filter := json.NewMapValue()
	map_port_number_schema_filter.SetObjectForMap("values", validation_constants.GetValidPortNumberCharacters())
	map_port_number_schema_filter.SetObjectForMap("function", validation_functions.GetWhitelistCharactersFunc())
	map_port_number_schema_filters.AppendMapValue(map_port_number_schema_filter)
	map_port_number_schema.SetArrayValue("filters", map_port_number_schema_filters)
	map_system_schema.SetMapValue("[port_number]", map_port_number_schema)
	// End port_number


	data.SetMapValue("[system_schema]", map_system_schema)

	
	getData := func() *json.Map {
		return &data
	}

	getHostName := func() (string, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[host_name]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		} 
		return temp_value.(string), temp_value_errors
	}

	getPortNumber := func() (string, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[port_number]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		}
		return temp_value.(string), temp_value_errors
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
