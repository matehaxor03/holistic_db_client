package class

type Host struct {
	Validate      func() []error
	ToJSONString  func() (*string, []error)
	Clone         func() *Host
	GetHostName   func() (string, []error)
	GetPortNumber func() (string, []error)
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

	getHostName := func() (string, []error) {
		temp_host_name_map, temp_host_name_map_errors := data.GetMap("[host_name]")
		if temp_host_name_map_errors != nil {
			return "", temp_host_name_map_errors
		}
		temp_host_name, temp_host_name_errors := temp_host_name_map.GetString("value")
		if temp_host_name_errors != nil {
			return "", temp_host_name_errors
		}
		c := CloneString(temp_host_name)
		return *c, nil
	}

	getPortNumber := func() (string, []error) {
		temp_port_number_map, temp_port_number_map_errors := data.GetMap("[port_number]")
		if temp_port_number_map_errors != nil {
			return "", temp_port_number_map_errors
		}
		temp_port_number, temp_port_number_errors := temp_port_number_map.GetString("value")
		if temp_port_number_errors != nil {
			return "", temp_port_number_errors
		}
		c := CloneString(temp_port_number)
		return *c, nil
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
			temp_host_name, _ := getHostName()
			temp_port_number, _ :=  getPortNumber()
			cloned, _ := NewHost(temp_host_name, temp_port_number)
			return cloned
		},
		GetHostName: func() (string, []error) {
			return getHostName()
		},
		GetPortNumber: func() (string, []error) {
			return getPortNumber()
		},
	}

	return &x, nil
}
