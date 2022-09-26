package class

import (
	"fmt"
	"reflect"
)

func GetHostNameValidCharacters() (string) {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890."
}

func  GetValidealidatePortCharacters() (string) {
	return "1234567890"
}


type Host struct {
	Validate func() ([]error)
	GetCLSCommand func() (*string, []error)
}

func NewHost(host_name *string, port_number *string) (*Host) {
	data := Map {
		"host_name":Map{"type":"*string","value":host_name,"mandatory":true,
		FILTERS(): Array{ Map {"values":GetHostNameValidCharacters(),"function":ValidateCharacters }}},
		"port_number":Map{"type":"*string","value":port_number,"mandatory":true,
		FILTERS(): Array{ Map {"values":GetValidealidatePortCharacters(),"function":ValidateCharacters }}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Host")
	}

	getCLSCommand := func() (*string, []error) {
		errors := validate()
		if len(errors) > 0 {
			return nil, errors
		}
	
		command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", data.M("host_name").S("value"),  data.M("port_number").S("value"))
	
		return &command, nil
	 }
	
	
	x := Host{
		Validate: func() ([]error) {
			return validate()
		},
		GetCLSCommand: func() (*string, []error) {
			return getCLSCommand()
		},
    }

	return &x
}

 
