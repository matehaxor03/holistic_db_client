package class

import (
	"strings"
	json "github.com/matehaxor03/holistic_json/json"
)

func GET_CHARACTER_SET_UTF8() string {
	return "utf8"
}

func GET_CHARACTER_SET_UTF8MB4() string {
	return "utf8mb4"
}

func GET_CHARACTER_SETS() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil(GET_CHARACTER_SET_UTF8())
	valid_chars.SetNil(GET_CHARACTER_SET_UTF8MB4())
	return valid_chars
}

func GET_COLLATE_UTF8_GENERAL_CI() string {
	return "utf8_general_ci"
}

func GET_COLLATE_UTF8MB4_0900_AI_CI() string {
	return "utf8mb4_0900_ai_ci"
}

func GET_COLLATES() json.Map {
	valid_chars := json.NewMapValue()
	valid_chars.SetNil(GET_COLLATE_UTF8_GENERAL_CI())
	valid_chars.SetNil(GET_COLLATE_UTF8MB4_0900_AI_CI())
	return valid_chars
}

type DatabaseCreateOptions struct {
	ToJSONString func(json *strings.Builder) ([]error)
	GetCharacterSet func() (*string, []error)
	GetCollate 		func() (*string, []error)
	Validate     func() []error
}

func newDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions, []error) {
	struct_type := "*class.DatabaseCreateOptions"

	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[character_set]", character_set)
	map_system_fields.SetObjectForMap("[collate]", collate)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.NewMapValue()


	map_character_set_schema := json.NewMapValue()
	map_character_set_schema.SetStringValue("type", "*string")

	map_character_set_schema_filters := json.NewArrayValue()
	map_character_set_schema_filter := json.NewMapValue()
	map_character_set_schema_filter.SetObjectForMap("values", GET_CHARACTER_SETS())
	map_character_set_schema_filter.SetObjectForMap("function",  getWhitelistStringFunc())
	map_character_set_schema_filters.AppendMapValue(map_character_set_schema_filter)
	map_character_set_schema.SetArrayValue("filters", map_character_set_schema_filters)
	map_system_schema.SetMapValue("[character_set]", map_character_set_schema)


	map_collate_schema := json.NewMapValue()
	map_collate_schema.SetStringValue("type", "*string")

	map_collate_schema_filters := json.NewArrayValue()
	map_collate_schema_filter := json.NewMapValue()
	map_collate_schema_filter.SetObjectForMap("values", GET_COLLATES())
	map_collate_schema_filter.SetObjectForMap("function",  getWhitelistStringFunc())
	map_collate_schema_filters.AppendMapValue(map_collate_schema_filter)
	map_collate_schema.SetArrayValue("filters", map_collate_schema_filters)
	map_system_schema.SetMapValue("[collate]", map_collate_schema)


	data.SetMapValue("[system_schema]", map_system_schema)

	/*
	data := json.Map{
		"[fields]": json.NewMapValue(),
		"[schema]": json.NewMapValue(),
		"[system_fields]":json.Map{"[character_set]":character_set, "[collate]":collate},
		"[system_schema]":json.Map{"[character_set]":json.Map{"type":"*string",
			"filters": json.Array{json.Map{"values": GET_CHARACTER_SETS(), "function": getWhitelistStringFunc()}}},
			"[collate]": json.Map{"type":"*string",
			"filters": json.Array{json.Map{"values": GET_COLLATES(), "function": getWhitelistStringFunc()}}},
		},
	}*/

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	get_character_set := func() (*string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[character_set]", "*string")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*string), temp_value_errors
	}

	get_collate := func() (*string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[collate]", "*string")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*string), temp_value_errors
	}

	validate_errors := validate()

	if len(validate_errors) > 0 {
		return nil, validate_errors
	}

	return &DatabaseCreateOptions{
		GetCharacterSet: func() (*string, []error) {
			return get_character_set()
		},
		GetCollate: func() (*string, []error) {
			return get_collate()
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
		},
		Validate: func() []error {
			return validate()
		},
	}, nil
}
