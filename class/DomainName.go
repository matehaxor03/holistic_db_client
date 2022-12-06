package class

import (
	json "github.com/matehaxor03/holistic_json/json"
)

func LOCALHOST_IP() string {
	return "127.0.0.1"
}

func GET_ALLOWED_DOMAIN_NAMES() json.Map {
	return json.Map{LOCALHOST_IP(): nil}
}

type DomainName struct {
	Validate      func() []error
	GetDomainName func() (string, []error)
}

func NewDomainName(domain_name string) (*DomainName, []error) {
	struct_type := "*DomainName"

	data := json.Map{
		"[fields]": json.Map{},
		"[schema]": json.Map{},
		"[system_fields]": json.Map{"[domain_name]": domain_name},
		"[system_schema]": json.Map{"[domain_name]": json.Map{"type":"string",
			"filters": json.Array{json.Map{"values": GET_ALLOWED_DOMAIN_NAMES(), "function": getWhitelistStringFunc()}}},
		},
	}

	getData := func() *json.Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "DomainName")
	}

	getDomainName := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]", "[domain_name]", "string")
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
