package db_client

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
		valid_chars := json.NewMapValue()
		valid_chars.SetNil("0")
		valid_chars.SetNil("1")
		valid_chars.SetNil("2")
		valid_chars.SetNil("3")
		valid_chars.SetNil("4")
		valid_chars.SetNil("5")
		valid_chars.SetNil("6")
		valid_chars.SetNil("7")
		valid_chars.SetNil("8")
		valid_chars.SetNil("9")
		valid_chars.SetNil("A")
		valid_chars.SetNil("B")
		valid_chars.SetNil("C")
		valid_chars.SetNil("D")
		valid_chars.SetNil("E")
		valid_chars.SetNil("F")
		valid_chars.SetNil("G")
		valid_chars.SetNil("H")
		valid_chars.SetNil("I")
		valid_chars.SetNil("J")
		valid_chars.SetNil("K")
		valid_chars.SetNil("L")
		valid_chars.SetNil("M")
		valid_chars.SetNil("N")
		valid_chars.SetNil("O")
		valid_chars.SetNil("P")
		valid_chars.SetNil("Q")
		valid_chars.SetNil("R")
		valid_chars.SetNil("S")
		valid_chars.SetNil("T")
		valid_chars.SetNil("U")
		valid_chars.SetNil("V")
		valid_chars.SetNil("W")
		valid_chars.SetNil("X")
		valid_chars.SetNil("Y")
		valid_chars.SetNil("Z")
		valid_chars.SetNil("_")
		valid_chars.SetNil("-")
		valid_chars.SetNil("a")
		valid_chars.SetNil("b")
		valid_chars.SetNil("c")
		valid_chars.SetNil("d")
		valid_chars.SetNil("e")
		valid_chars.SetNil("f")
		valid_chars.SetNil("g")
		valid_chars.SetNil("h")
		valid_chars.SetNil("i")
		valid_chars.SetNil("j")
		valid_chars.SetNil("k")
		valid_chars.SetNil("l")
		valid_chars.SetNil("m")
		valid_chars.SetNil("n")
		valid_chars.SetNil("o")
		valid_chars.SetNil("p")
		valid_chars.SetNil("q")
		valid_chars.SetNil("r")
		valid_chars.SetNil("s")
		valid_chars.SetNil("t")
		valid_chars.SetNil("u")
		valid_chars.SetNil("v")
		valid_chars.SetNil("w")
		valid_chars.SetNil("x")
		valid_chars.SetNil("y")
		valid_chars.SetNil("z")
		valid_chars.SetNil(".")
		return valid_chars
	}

	getValidPortCharacters := func() json.Map {
		valid_chars := json.NewMapValue()
		valid_chars.SetNil("0")
		valid_chars.SetNil("1")
		valid_chars.SetNil("2")
		valid_chars.SetNil("3")
		valid_chars.SetNil("4")
		valid_chars.SetNil("5")
		valid_chars.SetNil("6")
		valid_chars.SetNil("7")
		valid_chars.SetNil("8")
		valid_chars.SetNil("9")
		return valid_chars
	}

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

	map_host_name_schema_filters := json.NewArrayValue()
	map_host_name_schema_filter := json.NewMapValue()
	map_host_name_schema_filter.SetObjectForMap("values", getHostNameValidCharacters())
	map_host_name_schema_filter.SetObjectForMap("function",  getWhitelistCharactersFunc())
	map_host_name_schema_filters.AppendMapValue(map_host_name_schema_filter)
	map_host_name_schema.SetArrayValue("filters", map_host_name_schema_filters)
	map_system_schema.SetMapValue("[host_name]", map_host_name_schema)
	// End host_name

	// Start port_number
	map_port_number_schema := json.NewMapValue()
	map_port_number_schema.SetStringValue("type", "string")

	map_port_number_schema_filters := json.NewArrayValue()
	map_port_number_schema_filter := json.NewMapValue()
	map_port_number_schema_filter.SetObjectForMap("values", getValidPortCharacters())
	map_port_number_schema_filter.SetObjectForMap("function",  getWhitelistCharactersFunc())
	map_port_number_schema_filters.AppendMapValue(map_port_number_schema_filter)
	map_port_number_schema.SetArrayValue("filters", map_port_number_schema_filters)
	map_system_schema.SetMapValue("[port_number]", map_port_number_schema)
	// End port_number


	data.SetMapValue("[system_schema]", map_system_schema)

	
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
