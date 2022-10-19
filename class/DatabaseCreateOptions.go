package class

import (
	"fmt"
)

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET() string {
	return "character_set"
}

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE() string {
	return "collate"
}

type DatabaseCreateOptions struct {
	GetSQL       func() (*string, []error)
	ToJSONString func() string
	Clone        func() *DatabaseCreateOptions
	Validate     func() []error
}

func NewDatabaseCreateOptions(character_set *string, collate *string) *DatabaseCreateOptions {
	GET_CHARACTER_SET_UTF8 := func() string {
		return "utf8"
	}

	GET_CHARACTER_SET_UTF8MB4 := func() string {
		return "utf8mb4"
	}

	GET_CHARACTER_SETS := func() Map {
		return Map{GET_CHARACTER_SET_UTF8(): nil, GET_CHARACTER_SET_UTF8MB4(): nil}
	}

	GET_COLLATE_UTF8_GENERAL_CI := func() string {
		return "utf8_general_ci"
	}

	GET_COLLATE_UTF8MB4_0900_AI_CI := func() string {
		return "utf8mb4_0900_ai_ci"
	}

	GET_COLLATES := func() Map {
		return Map{GET_COLLATE_UTF8_GENERAL_CI(): nil, GET_COLLATE_UTF8MB4_0900_AI_CI(): nil}
	}

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
		ToJSONString: func() string {
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
