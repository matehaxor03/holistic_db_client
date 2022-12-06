package class

import (
	"strings"
)

type Host struct {
	Validate      func() []error
	ToJSONString  func(json *strings.Builder) ([]error)
	GetHostName   func() (string, []error)
	GetPortNumber func() (string, []error)
}

func newHost(host_name string, port_number string) (*Host, []error) {
	struct_type := "*Host"

	getHostNameValidCharacters := func() Map {
		temp := Map{"a": nil,
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

	getValidPortCharacters := func() Map {
		temp := Map{"0": nil,
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

	data := Map{
		"[fields]": Map{},
		"[schema]": Map{},
		"[system_fields]": Map{"[host_name]":host_name, "[port_number]":port_number},
		"[system_schema]": Map{
			"[host_name]":Map{"type":"string",
			"filters": Array{Map{"values": getHostNameValidCharacters(), "function": getWhitelistCharactersFunc()}}},
			"[port_number]": Map{"type":"string",
			"filters": Array{Map{"values": getValidPortCharacters(), "function": getWhitelistCharactersFunc()}}},
		},
	}

	getData := func() *Map {
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
