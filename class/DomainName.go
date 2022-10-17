package class

import (
	"strings"
)

func CloneDomainName(domain_name *DomainName) *DomainName {
	if domain_name == nil {
		return nil
	}

	return domain_name.Clone()
}

func LOCALHOST_IP() string {
	return "127.0.0.1"
}

func GET_ALLOWED_DOMAIN_NAMES() Map {
	return Map{LOCALHOST_IP(): nil}
}

type DomainName struct {
	Clone         func() *DomainName
	Validate      func() []error
	GetDomainName func() *string
}

func NewDomainName(domain_name *string) (*DomainName, []error) {

	data := Map{
		"[domain_name]": Map{"value": CloneString(domain_name), "mandatory": true,
			FILTERS(): Array{Map{"values": GET_ALLOWED_DOMAIN_NAMES(), "function": getWhitelistStringFunc()}}},
	}

	validate := func() []error {
		return ValidateData(data.Clone(), "DomainName")
	}

	getDomainName := func() *string {
		ptr := data.M("[domain_name]").S("value")
		if ptr == nil {
			return nil
		}
		cloneString := strings.Clone(*ptr)
		return &cloneString
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	x := DomainName{
		Validate: func() []error {
			return validate()
		},
		GetDomainName: func() *string {
			return getDomainName()
		},
		Clone: func() *DomainName {
			cloned, _ := NewDomainName(getDomainName())
			return cloned
		},
	}

	return &x, nil
}
