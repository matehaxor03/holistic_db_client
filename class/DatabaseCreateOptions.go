package class

import (
	"fmt"
)

func GET_CHARACTER_SET_UTF8() string {
	return "utf8"
}

func GET_CHARACTER_SET_UTF8MB4() string {
	return "utf8mb4"
}

func GET_CHARACTER_SETS() Map {
	return Map{GET_CHARACTER_SET_UTF8(): nil, GET_CHARACTER_SET_UTF8MB4(): nil}
}

func GET_COLLATE_UTF8_GENERAL_CI() string {
	return "utf8_general_ci"
}

func GET_COLLATE_UTF8MB4_0900_AI_CI() string {
	return "utf8mb4_0900_ai_ci"
}

func GET_COLLATES() Map {
	return Map{GET_COLLATE_UTF8_GENERAL_CI(): nil, GET_COLLATE_UTF8MB4_0900_AI_CI(): nil}
}

type DatabaseCreateOptions struct {
	GetSQL       func() (*string, []error)
	ToJSONString func() (*string, []error)
	Validate     func() []error
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions, []error) {

	data := Map{
		"[character_set]": Map{"value": CloneString(character_set), "mandatory": false,
			FILTERS(): Array{Map{"values": GET_CHARACTER_SETS(), "function": getWhitelistStringFunc()}}},
		"[collate]": Map{"value": CloneString(collate), "mandatory": false,
			FILTERS(): Array{Map{"values": GET_COLLATES(), "function": getWhitelistStringFunc()}}},
		"[validated]": Map{"value": false, "mandatory": true},
	}

	getData := func() *Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "DatabaseCreateOptions")
	}

	get_character_set := func() (*string, []error) {
		character_set_map, character_set_map_errors := data.GetMap("[character_set]")
		
		if character_set_map_errors != nil {
			return nil, character_set_map_errors
		}

		return character_set_map.GetString("value")
	}

	get_collate := func() (*string, []error) {
		collate_map, colalte_map_errors := data.GetMap("[collate]")
		
		if colalte_map_errors != nil {
			return nil, colalte_map_errors
		}

		return collate_map.GetString("value")
	}

	getSQL := func() (*string, []error) {
		errs := validate()
		if errs != nil {
			return nil, errs
		}

		sql_command := ""

		character_set, character_set_errors := get_character_set()
		if character_set_errors != nil {
			return nil, character_set_errors
		}

		if character_set != nil && *character_set != "" {
			sql_command += fmt.Sprintf("CHARACTER SET %s ", *character_set)
		}

		collate, collate_errors := get_collate()
		if collate_errors != nil {
			return nil, collate_errors
		}

		if collate != nil && *collate != "" {
			sql_command += fmt.Sprintf("COLLATE %s ", *collate)
		}

		return &sql_command, nil
	}

	x := DatabaseCreateOptions{
		GetSQL: func() (*string, []error) {
			return getSQL()
		},
		ToJSONString: func() (*string, []error) {
			return getData().ToJSONString()
		},
		Validate: func() []error {
			return validate()
		},
	}

	validate_errors := validate()

	if len(validate_errors) > 0 {
		return nil, validate_errors
	}

	return &x, nil
}
