package class

import (
	"fmt"
	"strings"
)

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_CHARACTER_SET() string {
	return "character_set"
}

func GET_TABLE_NAME_DATABASE_CREATE_OPTIONS_FIELD_NAME_COLLATE() string {
	return "collate"
}

type DatabaseCreateOptions struct {
	GetSQL func() (*string, []error)
	ToJSONString func() string
	Clone func() *DatabaseCreateOptions
	Validate func() ([]error)
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	var character_set_copy string
	if character_set != nil {
		character_set_copy = strings.Clone(*character_set)
	}
	
	var collate_copy string
	if collate != nil {
		collate_copy = strings.Clone(*collate)
	}
	
	data := Map {
		"character_set":Map{"type":"string","value":&character_set_copy,"mandatory":false,
		FILTERS(): Array{ Map {"values":GET_CHARACTER_SETS(),"function":ContainsExactMatch } }},
		"collate":Map{"type":"string","value":&collate_copy,"mandatory":false,
		FILTERS(): Array{ Map {"values":GET_COLLATES(),"function":ContainsExactMatch } }},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "DatabaseCreateOptions")
	}

	getSQL := func() (*string, []error) {
		errs := validate()
		if errs != nil {
			return nil, errs
		}
		
		sql_command := ""

		character_set := data.M("character_set").S("value")
		if character_set != nil && *character_set != "" {
			sql_command += fmt.Sprintf("CHARACTER SET %s ", *character_set)
		}
		
		collate := data.M("collate").S("value")
		if collate != nil  && *collate != "" {
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
			return NewDatabaseCreateOptions(data.M("character_set").S("value"), data.M("collate").S("value"))
		},
		Validate: func() ([]error) {
			return validate()
		},
    }
}
