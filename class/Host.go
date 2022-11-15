package class

type Host struct {
	Validate      func() []error
	ToJSONString  func() (*string, []error)
	Clone         func() *Host
	GetHostName   func() string
	GetPortNumber func() string
}

func CloneHost(host *Host) *Host {
	if host == nil {
		return host
	}

	return host.Clone()
}

func NewHost(host_name string, port_number string) (*Host, []error) {

	getHostNameValidCharacters := func() Map {
		temp := Map{"a": nil,
			"b": nil,
			"c": nil,
			"d": nil,
			"e": nil,
			"f": nil,
			"g": nil,
			"h": nil,
			"i": nil,
			"j": nil,
			"k": nil,
			"l": nil,
			"m": nil,
			"n": nil,
			"o": nil,
			"p": nil,
			"q": nil,
			"r": nil,
			"s": nil,
			"t": nil,
			"u": nil,
			"v": nil,
			"w": nil,
			"x": nil,
			"y": nil,
			"z": nil,
			"0": nil,
			"1": nil,
			"2": nil,
			"3": nil,
			"4": nil,
			"5": nil,
			"6": nil,
			"7": nil,
			"8": nil,
			"9": nil,
			"-": nil,
			".": nil}
		return temp
	}

	getValidPortCharacters := func() Map {
		temp := Map{"0": nil,
			"1": nil,
			"2": nil,
			"3": nil,
			"4": nil,
			"5": nil,
			"6": nil,
			"7": nil,
			"8": nil,
			"9": nil}
		return temp
	}

	data := Map{
		"[host_name]": Map{"value": CloneString(&host_name), "mandatory": true,
			FILTERS(): Array{Map{"values": getHostNameValidCharacters(), "function": getWhitelistCharactersFunc()}}},
		"[port_number]": Map{"value": CloneString(&port_number), "mandatory": true,
			FILTERS(): Array{Map{"values": getValidPortCharacters(), "function": getWhitelistCharactersFunc()}}},
	}

	getHostName := func() string {
		host_name, _ := data.M("[host_name]").GetString("value")
		c := CloneString(host_name)
		return *c
	}

	getPortNumber := func() string {
		port_number, _ := data.M("[port_number]").GetString("value")
		c := CloneString(port_number)
		return *c
	}

	validate := func() []error {
		data_cloned, data_cloned_errors := data.Clone()
		if data_cloned_errors != nil {
			return data_cloned_errors
		}

		return ValidateData(*data_cloned, "Host")
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	x := Host{
		Validate: func() []error {
			return validate()
		},
		ToJSONString: func() (*string, []error) {
			data_cloned, data_cloned_errors := data.Clone()
			if data_cloned_errors != nil {
				return nil, data_cloned_errors
			}

			return data_cloned.ToJSONString()
		},
		Clone: func() *Host {
			cloned, _ := NewHost(getHostName(), getPortNumber())
			return cloned
		},
		GetHostName: func() string {
			return getHostName()
		},
		GetPortNumber: func() string {
			return getPortNumber()
		},
	}

	return &x, nil
}
