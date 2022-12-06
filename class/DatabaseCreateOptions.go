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
	return json.Map{GET_CHARACTER_SET_UTF8(): nil, GET_CHARACTER_SET_UTF8MB4(): nil}
}

func GET_COLLATE_UTF8_GENERAL_CI() string {
	return "utf8_general_ci"
}

func GET_COLLATE_UTF8MB4_0900_AI_CI() string {
	return "utf8mb4_0900_ai_ci"
}

func GET_COLLATES() json.Map {
	return json.Map{GET_COLLATE_UTF8_GENERAL_CI(): nil, GET_COLLATE_UTF8MB4_0900_AI_CI(): nil}
}

type DatabaseCreateOptions struct {
	ToJSONString func(json *strings.Builder) ([]error)
	GetCharacterSet func() (*string, []error)
	GetCollate 		func() (*string, []error)
	Validate     func() []error
}

func newDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions, []error) {
	struct_type := "*class.DatabaseCreateOptions"

	data := json.Map{
		"[fields]": json.Map{},
		"[schema]": json.Map{},
		"[system_fields]":json.Map{"[character_set]":character_set, "[collate]":collate},
		"[system_schema]":json.Map{"[character_set]":json.Map{"type":"*string",
			"filters": json.Array{json.Map{"values": GET_CHARACTER_SETS(), "function": getWhitelistStringFunc()}}},
			"[collate]": json.Map{"type":"*string",
			"filters": json.Array{json.Map{"values": GET_COLLATES(), "function": getWhitelistStringFunc()}}},
		},
	}

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
