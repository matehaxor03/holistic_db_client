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
	Validate func() ([]error)
	GetSQL func() (*string, []error)
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	data := Map {
		"character_set":Map{"type":"string","value":character_set,"mandatory":false,
		FILTERS(): Array{ Map {"values":GET_CHARACTER_SETS(),"function":ContainsExactMatch } }},
		"collate":Map{"type":"string","value":collate,"mandatory":false,
		FILTERS(): Array{ Map {"values":GET_COLLATES(),"function":ContainsExactMatch } }},
	}
	
	getData := func() Map {
		return data
    }

	getCharacterSet := func() (*string) {
		v := data.M("character_set").S("value")
		if v == nil {
			return nil
		}
		clone := strings.Clone(*v)
		return &clone
	}

	getCollate := func() (*string) {
		v := data.M("collate").S("value")
		if v == nil {
			return nil
		}
		clone := strings.Clone(*v)
		return &clone
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(getData(), "DatabaseCreateOptions")
	}

	getSQL := func() (*string, []error) {
		errs := validate()
		if errs != nil {
			return nil, errs
		}
		
		sql_command := ""

		character_set := getCharacterSet()
		if character_set != nil {
			sql_command += fmt.Sprintf("CHARACTER SET %s ", *character_set)
		}
		
		collate := getCollate()
		if collate != nil {
			sql_command += fmt.Sprintf("COLLATE %s ", *collate)
		}

		return &sql_command, nil
	}
	
	return &DatabaseCreateOptions{
		Validate: func() ([]error) {
			return validate()
		},
		GetSQL: func() (*string, []error) {
			return getSQL()
		},
    }
}
