package class

type Host struct {
	Validate      func() []error
	ToJSONString  func() (*string, []error)
	GetHostName   func() (string, []error)
	GetPortNumber func() (string, []error)
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
		"[fields]": Map{"host_name":host_name, "port_number":port_number},
		"[schema]": Map{
			"host_name":Map{"type":"*string","mandatory": true, "validated":false,
			FILTERS(): Array{Map{"values": getHostNameValidCharacters(), "function": getWhitelistCharactersFunc()}}},
			"port_number": Map{"type":"*string", "mandatory": true, "validated":false,
			FILTERS(): Array{Map{"values": getValidPortCharacters(), "function": getWhitelistCharactersFunc()}}},
		},
	}

	getData := func() *Map {
		return &data
	}

	getHostName := func() (*string, []error) {
		return GetStringField(getData(), "host_name")
	}

	getPortNumber := func() (*string, []error) {
		return GetStringField(getData(), "port_number")
	}

	validate := func() []error {
		return ValidateData(getData(), "Host")
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
			return getData().ToJSONString()
		},
		GetHostName: func() (string, []error) {
			host_name_ptr, host_name_ptr_errors := getHostName()
			if host_name_ptr_errors != nil {
				return "", host_name_ptr_errors
			}
			return *host_name_ptr, nil
		},
		GetPortNumber: func() (string, []error) {
			port_number_ptr, port_number_ptr_errors := getPortNumber()
			if port_number_ptr_errors != nil {
				return "", port_number_ptr_errors
			}
			return *port_number_ptr, nil
		},
	}

	return &x, nil
}
