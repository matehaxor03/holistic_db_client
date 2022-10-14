package class

type Host struct {
	Validate func() ([]error)
	ToJSONString func() string
	Clone func() *Host
	GetHostName func() (*string)
	GetPortNumber func() (*string)
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
		"[host_name]":Map{"value":CloneString(host_name),"mandatory":true,
		FILTERS(): Array{ Map {"values":getHostNameValidCharacters(),"function":getWhitelistCharactersFunc() }}},
		"[port_number]":Map{"value":CloneString(port_number),"mandatory":true,
		FILTERS(): Array{ Map {"values":getValidPortCharacters(),"function":getWhitelistCharactersFunc() }}},
	}

	getHostName := func() (*string) {
		return CloneString(data.M("[host_name]").S("value"))
	}

	getPortNumber := func() (*string) {
		return CloneString(data.M("[port_number]").S("value"))
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Host")
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}
	
	x := Host{
		Validate: func() ([]error) {
			return validate()
		},
		ToJSONString: func() string {
			return data.Clone().ToJSONString()
		},
		Clone: func() *Host {
			cloned, _ := NewHost(getHostName(), getPortNumber())
			return cloned
		},
		GetHostName: func() (*string) {
			return getHostName()
		},
		GetPortNumber: func() (*string) {
			return getPortNumber()
		},
    }

	return &x, nil
}

 
