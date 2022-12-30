package class

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
)

type Host struct {
	Validate      func() []error
	ToJSONString  func(json *strings.Builder) ([]error)
	GetHostName   func() (string, []error)
	GetPortNumber func() (string, []error)
}

func newHost(host_name string, port_number string) (*Host, []error) {
	struct_type := "*Host"

	getHostNameValidCharacters := func() json.Map {
		temp := json.Map{"a": nil,
			"b": nil,
			"c": nil,
			"d": nil,
			"e": nil,
			"f": nil,
			"g": nil,
			"h": nil,
			"i": nil,
			"j": nil,
			"k": nil,
			"l": nil,
			"m": nil,
			"n": nil,
			"o": nil,
			"p": nil,
			"q": nil,
			"r": nil,
			"s": nil,
			"t": nil,
			"u": nil,
			"v": nil,
			"w": nil,
			"x": nil,
			"y": nil,
			"z": nil,
			"0": nil,
			"1": nil,
			"2": nil,
			"3": nil,
			"4": nil,
			"5": nil,
			"6": nil,
			"7": nil,
			"8": nil,
			"9": nil,
			"-": nil,
			".": nil}
		return temp
	}

	getValidPortCharacters := func() json.Map {
		temp := json.Map{"0": nil,
			"1": nil,
			"2": nil,
			"3": nil,
			"4": nil,
			"5": nil,
			"6": nil,
			"7": nil,
			"8": nil,
			"9": nil}
		return temp
	}

	data := json.Map{}
	data.SetMapValue("[fields]", json.Map{})
	data.SetMapValue("[schema]", json.Map{})

	map_system_fields := json.Map{}
	map_system_fields.SetObject("[host_name]", host_name)
	map_system_fields.SetObject("[port_number]", port_number)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.Map{}

	// Start host_name
	map_host_name_schema := json.Map{}
	map_host_name_schema.SetStringValue("type", "string")

	map_host_name_schema_filters := json.Array{}
	map_host_name_schema_filter := json.Map{}
	map_host_name_schema_filter.SetObject("values", getHostNameValidCharacters())
	map_host_name_schema_filter.SetObject("function",  getWhitelistCharactersFunc())
	map_host_name_schema_filters.AppendMapValue(map_host_name_schema_filter)
	map_host_name_schema.SetArrayValue("filters", map_host_name_schema_filters)
	map_system_schema.SetMapValue("[host_name]", map_host_name_schema)
	// End host_name

	// Start port_number
	map_port_number_schema := json.Map{}
	map_port_number_schema.SetStringValue("type", "string")

	map_port_number_schema_filters := json.Array{}
	map_port_number_schema_filter := json.Map{}
	map_port_number_schema_filter.SetObject("values", getValidPortCharacters())
	map_port_number_schema_filter.SetObject("function",  getWhitelistCharactersFunc())
	map_port_number_schema_filters.AppendMapValue(map_port_number_schema_filter)
	map_port_number_schema.SetArrayValue("filters", map_port_number_schema_filters)
	map_system_schema.SetMapValue("[port_number]", map_port_number_schema)
	// End port_number


	data.SetMapValue("[system_schema]", map_system_schema)
	/*
	data := json.Map{
		"[fields]": json.Map{},
		"[schema]": json.Map{},
		"[system_fields]": json.Map{"[host_name]":host_name, "[port_number]":port_number},
		"[system_schema]": json.Map{
			"[host_name]":json.Map{"type":"string",
			"filters": json.Array{json.Map{"values": getHostNameValidCharacters(), "function": getWhitelistCharactersFunc()}}},
			"[port_number]": json.Map{"type":"string",
			"filters": json.Array{json.Map{"values": getValidPortCharacters(), "function": getWhitelistCharactersFunc()}}},
		},
	}*/

	getData := func() *json.Map {
		return &data
	}

	getHostName := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[host_name]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		} 
		return temp_value.(string), temp_value_errors
	}

	getPortNumber := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[port_number]", "string")
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
