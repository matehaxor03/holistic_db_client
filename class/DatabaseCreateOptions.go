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
	Clone        func() *DatabaseCreateOptions
	Validate     func() []error
}

func NewDatabaseCreateOptions(character_set *string, collate *string) *DatabaseCreateOptions {

	data := Map{
		"[character_set]": Map{"value": CloneString(character_set), "mandatory": false,
			FILTERS(): Array{Map{"values": GET_CHARACTER_SETS(), "function": getWhitelistStringFunc()}}},
		"[collate]": Map{"value": CloneString(collate), "mandatory": false,
			FILTERS(): Array{Map{"values": GET_COLLATES(), "function": getWhitelistStringFunc()}}},
	}

	validate := func() []error {
		return ValidateData(data, "DatabaseCreateOptions")
	}

	getSQL := func() (*string, []error) {
		errs := validate()
		if errs != nil {
			return nil, errs
		}

		sql_command := ""

		character_set, _ := data.M("[character_set]").GetString("value")
		if character_set != nil && *character_set != "" {
			sql_command += fmt.Sprintf("CHARACTER SET %s ", *character_set)
		}

		collate, _ := data.M("[collate]").GetString("value")
		if collate != nil && *collate != "" {
			sql_command += fmt.Sprintf("COLLATE %s ", *collate)
		}

		return &sql_command, nil
	}

	return &DatabaseCreateOptions{
		GetSQL: func() (*string, []error) {
			return getSQL()
		},
		ToJSONString: func() (*string, []error) {
			return data.Clone().ToJSONString()
		},
		Clone: func() *DatabaseCreateOptions {
			character_set, _ := data.M("[character_set]").GetString("value")
			collate, _ := data.M("[collate]").GetString("value")
			return NewDatabaseCreateOptions(character_set, collate)
		},
		Validate: func() []error {
			return validate()
		},
	}
}
