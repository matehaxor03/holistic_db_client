package class

func CloneCredentials(credentials *Credentials) *Credentials {
	if credentials == nil {
		return credentials
	}

	return credentials.Clone()
}

type Credentials struct {
	Validate     func() []error
	GetUsername  func() (string, []error)
	GetPassword  func() (string, []error)
	ToJSONString func() (*string, []error)
	Clone        func() *Credentials
}

func GetCredentialsUsernameValidCharacters() Map {
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
		"A": nil,
		"B": nil,
		"C": nil,
		"D": nil,
		"E": nil,
		"F": nil,
		"G": nil,
		"H": nil,
		"I": nil,
		"J": nil,
		"K": nil,
		"L": nil,
		"M": nil,
		"N": nil,
		"O": nil,
		"P": nil,
		"Q": nil,
		"R": nil,
		"S": nil,
		"T": nil,
		"U": nil,
		"V": nil,
		"W": nil,
		"X": nil,
		"Y": nil,
		"Z": nil,
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
		"_": nil}
	return temp
}

func GetCredentialPasswordValidCharacters() Map {
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
		"A": nil,
		"B": nil,
		"C": nil,
		"D": nil,
		"E": nil,
		"F": nil,
		"G": nil,
		"H": nil,
		"I": nil,
		"J": nil,
		"K": nil,
		"L": nil,
		"M": nil,
		"N": nil,
		"O": nil,
		"P": nil,
		"Q": nil,
		"R": nil,
		"S": nil,
		"T": nil,
		"U": nil,
		"V": nil,
		"W": nil,
		"X": nil,
		"Y": nil,
		"Z": nil,
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
		".": nil,
		"=": nil,
		"_": nil}
	return temp
}

func NewCredentials(username string, password string) (*Credentials, []error) {

	data := Map{
		"[username]": Map{"value": CloneString(&username), "mandatory": true,
			FILTERS(): Array{Map{"values": GetCredentialsUsernameValidCharacters(), "function": getWhitelistCharactersFunc()}}},
		"[password]": Map{"value": CloneString(&password), "mandatory": true,
			FILTERS(): Array{Map{"values": GetCredentialPasswordValidCharacters(), "function": getWhitelistCharactersFunc()}}},
	}

	validate := func() []error {
		data_cloned, data_cloned_errors := data.Clone()
		if data_cloned_errors != nil {
			return data_cloned_errors
		}

		return ValidateData(*data_cloned, "Credentials")
	}

	getUsername := func() (string, []error) {
		temp_username_map, temp_username_map_errors :=  data.GetMap("[username]")
		if temp_username_map_errors != nil {
			return "", temp_username_map_errors
		}
		username_value, username_value_errors := temp_username_map.GetString("value")
		if username_value_errors != nil {
			return "", username_value_errors
		}
		return *(CloneString(username_value)), nil
	}

	getPassword := func() (string, []error) {
		temp_password_map, temp_password_map_errors := data.GetMap("[password]")
		if temp_password_map_errors != nil {
			return "", temp_password_map_errors
		}
		password_value, password_value_errors := temp_password_map.GetString("value")
		if password_value_errors != nil {
			return "", password_value_errors
		}
		return (*CloneString(password_value)), nil
	}

	x := Credentials{
		Validate: func() []error {
			return validate()
		},
		GetUsername: func() (string, []error) {
			return getUsername()
		},
		GetPassword: func() (string, []error) {
			return getPassword()
		},
		ToJSONString: func() (*string, []error) {
			data_cloned, data_cloned_errors := data.Clone()
			if data_cloned_errors != nil {
				return nil, data_cloned_errors
			}
			return data_cloned.ToJSONString()
		},
		Clone: func() (*Credentials) {
			temp_username, _ := getUsername()
			temp_password, _ := getPassword()

			cloned, _ := NewCredentials(temp_username, temp_password)
			return cloned
		},
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}
