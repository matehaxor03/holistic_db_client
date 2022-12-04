package class

func LOCALHOST_IP() string {
	return "127.0.0.1"
}

func GET_ALLOWED_DOMAIN_NAMES() Map {
	return Map{LOCALHOST_IP(): nil}
}

type DomainName struct {
	Validate      func() []error
	GetDomainName func() (string, []error)
}

func newDomainName(domain_name string) (*DomainName, []error) {
	struct_type := "*DomainName"

	data := Map{
		"[fields]": Map{},
		"[schema]": Map{},
		"[system_fields]": Map{"[domain_name]": domain_name},
		"[system_schema]": Map{"[domain_name]":Map{"mandatory": true,
			FILTERS(): Array{Map{"values": GET_ALLOWED_DOMAIN_NAMES(), "function": getWhitelistStringFunc()}}},
		},
	}

	getData := func() *Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), "DomainName")
	}

	getDomainName := func() (*string, []error) {
		return GetStringField(struct_type, getData(), "[system_schema]", "[system_fields]", "[domain_name]")
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	x := DomainName{
		Validate: func() []error {
			return validate()
		},
		GetDomainName: func() (string, []error) {
			domain_name_ptr, domain_name_ptr_errors := getDomainName()
			if domain_name_ptr_errors != nil {
				return "", domain_name_ptr_errors
			}
			return *domain_name_ptr, nil
		},
	}

	return &x, nil
}
