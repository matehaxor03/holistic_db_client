package class

import (
	"strings"
)

func getValidateDomainNameCharacters() (string) {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.%"
}

type DomainName struct {
	Validate func() ([]error)
	GetDomainName func() (*string)
}

func NewDomainName(domain_name *string) (*DomainName) {
	data := Map {
		"domain_name":Map{"type":"*string","value":domain_name,"mandatory":true,
		FILTERS(): Array{ Map {"values":getValidateDomainNameCharacters(),"function":ValidateCharacters }}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "DomainName")
	}

	getDomainName := func () (*string) {
		ptr := data.M("domain_name").S("value")
		if ptr == nil {
			return nil
		}
		cloneString := strings.Clone(*ptr)
		return &cloneString
	}
	
	x := DomainName{
		Validate: func() ([]error) {
			return validate()
		},
		GetDomainName: func() (*string) {
			return getDomainName()
		},
    }

	return &x
}