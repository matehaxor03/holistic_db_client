package class

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func LOCALHOST_IP() string {
	return "127.0.0.1"
}

func GET_ALLOWED_DOMAIN_NAMES() json.Map {
	return json.Map{LOCALHOST_IP(): nil}
}

func get_domain_name_characters() json.Map {
	return json.Map{
		".": nil,
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
		"A": nil,
		"B": nil,
		"C": nil,
		"D": nil,
		"E": nil,
		"F": nil,
		"G": nil,
		"H": nil,
		"I": nil,
		"J": nil,
		"K": nil,
		"L": nil,
		"M": nil,
		"N": nil,
		"O": nil,
		"P": nil,
		"Q": nil,
		"R": nil,
		"S": nil,
		"T": nil,
		"U": nil,
		"V": nil,
		"W": nil,
		"X": nil,
		"Y": nil,
		"Z": nil,
		"_": nil,
		"a": nil,
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
		"z": nil}
}

type DomainName struct {
	Validate      func() []error
	GetDomainName func() (string, []error)
}

func NewDomainName(domain_name string) (*DomainName, []error) {
	struct_type := "*DomainName"


	data := json.Map{}
	data.SetMapValue("[fields]", json.Map{})
	data.SetMapValue("[schema]", json.Map{})

	map_system_fields := json.Map{}
	map_system_fields.SetObject("[domain_name]", domain_name)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.Map{}


	map_domain_name_schema := json.Map{}
	map_domain_name_schema.SetStringValue("type", "string")

	map_domain_name_schema_filters := json.Array{}
	map_domain_name_schema_filter := json.Map{}
	map_domain_name_schema_filter.SetObject("values", GET_ALLOWED_DOMAIN_NAMES())
	map_domain_name_schema_filter.SetObject("function",  getWhitelistStringFunc())
	map_domain_name_schema_filters.AppendMapValue(map_domain_name_schema_filter)
	map_domain_name_schema.SetArrayValue("filters", map_domain_name_schema_filters)
	map_system_schema.SetMapValue("[domain_name]", map_domain_name_schema)


	data.SetMapValue("[system_schema]", map_system_schema)

	/*
	data := json.Map{
		"[fields]": json.Map{},
		"[schema]": json.Map{},
		"[system_fields]": json.Map{"[domain_name]": domain_name},
		"[system_schema]": json.Map{"[domain_name]": json.Map{"type":"string",
			"filters": json.Array{json.Map{"values": GET_ALLOWED_DOMAIN_NAMES(), "function": getWhitelistStringFunc()}}},
		},
	}*/

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "DomainName")
	}

	getDomainName := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[domain_name]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		}
		return temp_value.(string), nil
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &DomainName{
		Validate: func() []error {
			return validate()
		},
		GetDomainName: func() (string, []error) {
			return getDomainName()
		},
	}, nil
}
