package dao

import (
	json "github.com/matehaxor03/holistic_json/json"
	validate "github.com/matehaxor03/holistic_db_client/validate"
	helper "github.com/matehaxor03/holistic_db_client/helper"
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
)

type DomainName struct {
	Validate      func() []error
	GetDomainName func() (string, []error)
}

func NewDomainName(verify *validate.Validator, domain_name string) (*DomainName, []error) {
	data := json.NewMapValue()
	data.SetMapValue("[fields]", json.NewMapValue())
	data.SetMapValue("[schema]", json.NewMapValue())

	map_system_fields := json.NewMapValue()
	map_system_fields.SetObjectForMap("[domain_name]", domain_name)
	data.SetMapValue("[system_fields]", map_system_fields)

	map_system_schema := json.NewMapValue()


	map_domain_name_schema := json.NewMapValue()
	map_domain_name_schema.SetStringValue("type", "*string")

	map_domain_name_schema_filters := json.NewArrayValue()
	map_domain_name_schema_filter := json.NewMapValue()
	map_domain_name_schema_filter.SetObjectForMap("function",  verify.GetValidateDomainNameFunc())
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
		var errors []error
		temp_value, temp_value_errors := helper.GetField(*getData(), "[system_schema]", "[system_fields]", "[domain_name]", "string")
		if temp_value_errors != nil {
			errors = append(errors, temp_value_errors...)
		} else if common.IsNil(temp_value) {
			errors = append(errors, fmt.Errorf("domain name is nil"))
		}
		if len(errors) > 0 {
			return "", errors
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
