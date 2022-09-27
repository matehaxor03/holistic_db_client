package class

import (
	"fmt"
	"strings"
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
	ToJSONString func() string
	Clone func() *Host
}

func NewHost(host_name *string, port_number *string) (*Host) {
	var host_name_copy string 
	if host_name != nil {
		host_name_copy = strings.Clone(*host_name)
	}
	var port_number_copy string
	if port_number != nil {
		port_number_copy = strings.Clone(*port_number)
	}
	
	data := Map {
		"host_name":Map{"type":"*string","value":host_name_copy,"mandatory":true,
		FILTERS(): Array{ Map {"values":GetHostNameValidCharacters(),"function":ValidateCharacters }}},
		"port_number":Map{"type":"*string","value":port_number_copy,"mandatory":true,
		FILTERS(): Array{ Map {"values":GetValidealidatePortCharacters(),"function":ValidateCharacters }}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data, "Host")
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
		ToJSONString: func() string {
			return data.Clone().ToJSONString()
		},
		Clone: func() *Host {
			return NewHost(data.M("host_name").S("value"), data.M("port_number").S("value"))
		},
    }

	return &x
}

 
