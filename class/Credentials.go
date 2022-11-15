package class

func CloneCredentials(credentials *Credentials) *Credentials {
	if credentials == nil {
		return credentials
	}

	return credentials.Clone()
}

type Credentials struct {
	Validate     func() []error
	GetUsername  func() *string
	GetPassword  func() *string
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

func NewCredentials(username *string, password *string) (*Credentials, []error) {

	data := Map{
		"[username]": Map{"value": CloneString(username), "mandatory": true,
			FILTERS(): Array{Map{"values": GetCredentialsUsernameValidCharacters(), "function": getWhitelistCharactersFunc()}}},
		"[password]": Map{"value": CloneString(password), "mandatory": true,
			FILTERS(): Array{Map{"values": GetCredentialPasswordValidCharacters(), "function": getWhitelistCharactersFunc()}}},
	}

	validate := func() []error {
		data_cloned, data_cloned_errors := data.Clone()
		if data_cloned_errors != nil {
			return data_cloned_errors
		}

		return ValidateData(*data_cloned, "Credentials")
	}

	getUsername := func() *string {
		username_value, _ := data.M("[username]").GetString("value")
		return CloneString(username_value)
	}

	getPassword := func() *string {
		password_value, _ := data.M("[password]").GetString("value")
		return CloneString(password_value)
	}

	x := Credentials{
		Validate: func() []error {
			return validate()
		},
		GetUsername: func() *string {
			return getUsername()
		},
		GetPassword: func() *string {
			return getPassword()
		},
		ToJSONString: func() (*string, []error) {
			data_cloned, data_cloned_errors := data.Clone()
			if data_cloned_errors != nil {
				return nil, data_cloned_errors
			}
			return data_cloned.ToJSONString()
		},
		Clone: func() *Credentials {
			cloned, _ := NewCredentials(getUsername(), getPassword())
			return cloned
		},
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}
