package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

type DatabaseCreateOptions struct {
	GetCharacterSet func() (*string, []error)
	GetCollate 		func() (*string, []error)
	Validate     func() []error
}

func newDatabaseCreateOptions(verify *validate.Validator, character_set *string, collate *string) (*DatabaseCreateOptions, []error) {
	struct_type := "*dao.DatabaseCreateOptions"

	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[character_set]", character_set)
	map_system_fields.SetObjectForMap("[collate]", collate)
	data.SetMapValue("[system_fields]", map_system_fields)

	map_system_schema := json.NewMapValue()


	map_character_set_schema := json.NewMapValue()
	map_character_set_schema.SetStringValue("type", "*string")

	map_character_set_schema_filters := json.NewArrayValue()
	map_character_set_schema_filter := json.NewMapValue()
	map_character_set_schema_filter.SetObjectForMap("function",  verify.GetValidateCharacterSetFunc())
	map_character_set_schema_filters.AppendMapValue(map_character_set_schema_filter)
	map_character_set_schema.SetArrayValue("filters", map_character_set_schema_filters)
	map_system_schema.SetMapValue("[character_set]", map_character_set_schema)


	map_collate_schema := json.NewMapValue()
	map_collate_schema.SetStringValue("type", "*string")

	map_collate_schema_filters := json.NewArrayValue()
	map_collate_schema_filter := json.NewMapValue()
	map_collate_schema_filter.SetObjectForMap("function",  verify.GetValidateCollateFunc())
	map_collate_schema_filters.AppendMapValue(map_collate_schema_filter)
	map_collate_schema.SetArrayValue("filters", map_collate_schema_filters)
	map_system_schema.SetMapValue("[collate]", map_collate_schema)

	data.SetMapValue("[system_schema]", map_system_schema)


	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	get_character_set := func() (*string, []error) {
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]",  "[character_set]", "*string")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*string), temp_value_errors
	}

	get_collate := func() (*string, []error) {
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]",  "[collate]", "*string")
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
		Validate: func() []error {
			return validate()
		},
	}, nil
}
