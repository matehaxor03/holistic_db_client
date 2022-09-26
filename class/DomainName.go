package class

import (
	"fmt"
)

func getValidateDomainNameCharacters() (string) {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.%"
}

type DomainName struct {
	Validate func() ([]error)
}

func NewDomainName(domain_name *string) (*DomainName) {
	data := Map {
		"domain_name":Map{"type":"*string","value":domain_name,"mandatory":true,
		FILTERS(): Array{ Map {"values":getValidateDomainNameCharacters(),"function":ValidateCharacters }}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "DomainName")
	}
	
	
	x := DomainName{
		Validate: func() ([]error) {
			return validate()
		},
    }

	return &x
}