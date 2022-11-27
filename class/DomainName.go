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

func NewDomainName(domain_name string) (*DomainName, []error) {

	data := Map{
		"[domain_name]": Map{"value": CloneString(&domain_name), "mandatory": true,
			FILTERS(): Array{Map{"values": GET_ALLOWED_DOMAIN_NAMES(), "function": getWhitelistStringFunc()}}},
	}

	getData := func() *Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(data, "DomainName")
	}

	getDomainName := func() (*string, []error) {
		domain_name_map, domain_name_map_errors := getData().GetMap("[domain_name]")
		if domain_name_map_errors != nil {
			return nil, domain_name_map_errors
		}
		return domain_name_map.GetString("value")
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
