package class

import (
	"fmt"
)

type Host struct {
	Validate func() ([]error)
	GetCLSCommand func() (*string, []error)
	ToJSONString func() string
	Clone func() *Host
}

func CloneHost(host *Host) *Host {
	if host == nil {
		return host
	}

	return host.Clone()
}

func NewHost(host_name *string, port_number *string) (*Host, []error) {
	
	getHostNameValidCharacters := func() (*string) {
		temp := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890."
		return &temp
	}
	
	getValidPortCharacters := func() (*string) {
		temp := "1234567890"
		return &temp
	}
	
	data := Map {
		"host_name":Map{"type":"*string","value":CloneString(host_name),"mandatory":true,
		FILTERS(): Array{ Map {"values":getHostNameValidCharacters(),"function":getValidateCharacters() }}},
		"port_number":Map{"type":"*string","value":CloneString(port_number),"mandatory":true,
		FILTERS(): Array{ Map {"values":getValidPortCharacters(),"function":getValidateCharacters() }}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Host")
	}

	getCLSCommand := func() (*string, []error) {
		errors := validate()
		if len(errors) > 0 {
			return nil, errors
		}
	
		command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", *(data.M("host_name").S("value")), *(data.M("port_number").S("value")))
	
		return &command, nil
	 }

	errors := validate()

	if errors != nil {
		return nil, errors
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
			cloned, _ := NewHost(data.M("host_name").S("value"), data.M("port_number").S("value"))
			return cloned
		},
    }

	return &x, nil
}

 
