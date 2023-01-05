package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	validation_functions "github.com/matehaxor03/holistic_db_client/validation_functions"
	helper "github.com/matehaxor03/holistic_db_client/helper"
)

func LOCALHOST_IP() string {
	return "127.0.0.1"
}

func GET_ALLOWED_DOMAIN_NAMES() json.Map {
	valid := json.NewMapValue()
	valid.SetNil(LOCALHOST_IP())
	return valid
}

type DomainName struct {
	Validate      func() []error
	GetDomainName func() (string, []error)
}

func NewDomainName(domain_name string) (*DomainName, []error) {
	struct_type := "*dao.DomainName"


	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[domain_name]", domain_name)
	data.SetMapValue("[system_fields]", map_system_fields)

	///

	map_system_schema := json.NewMapValue()


	map_domain_name_schema := json.NewMapValue()
	map_domain_name_schema.SetStringValue("type", "string")

	map_domain_name_schema_filters := json.NewArrayValue()
	map_domain_name_schema_filter := json.NewMapValue()
	map_domain_name_schema_filter.SetObjectForMap("values", GET_ALLOWED_DOMAIN_NAMES())
	map_domain_name_schema_filter.SetObjectForMap("function",  validation_functions.GetWhitelistStringFunc())
	map_domain_name_schema_filters.AppendMapValue(map_domain_name_schema_filter)
	map_domain_name_schema.SetArrayValue("filters", map_domain_name_schema_filters)
	map_system_schema.SetMapValue("[domain_name]", map_domain_name_schema)


	data.SetMapValue("[system_schema]", map_system_schema)

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "DomainName")
	}

	getDomainName := func() (string, []error) {
		temp_value, temp_value_errors := helper.GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[domain_name]", "string")
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
