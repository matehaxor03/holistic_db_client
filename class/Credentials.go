package class

import (
	"strings"
)

type Credentials struct {
	Validate     func() []error
	GetUsername  func() (string, []error)
	GetPassword  func() (*string, []error)
	ToJSONString func(json *strings.Builder) ([]error)
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

func newCredentials(username string, password *string) (*Credentials, []error) {
	struct_type := "*Credentials"

	data := Map{
		"[fields]": Map{},
		"[schema]": Map{},
		"[system_fields]":Map{"[username]":username,"[password]":password},
		"[system_schema]":Map{"[username]":Map{"type":"*string","mandatory": true, 
			FILTERS(): Array{Map{"values": GetCredentialsUsernameValidCharacters(), "function": getWhitelistCharactersFunc()}}},
							 "[password]": Map{"type":"*string","mandatory": false, 
			FILTERS(): Array{Map{"values": GetCredentialPasswordValidCharacters(), "function": getWhitelistCharactersFunc()}}},
							},
	}

	getData := func() *Map {
		return &data
	}

	validate := func() []error {
		return ValidateData(getData(), struct_type)
	}

	getUsername := func() (string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[username]", "string")
		if temp_value_errors != nil {
			return "", temp_value_errors
		} else if temp_value == nil {
			return "", temp_value_errors
		}
		return temp_value.(string), nil
	}

	getPassword := func() (*string, []error) {
		temp_value, temp_value_errors := GetField(struct_type, getData(), "[system_schema]", "[system_fields]",  "[password]", "*string")
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}
		return temp_value.(*string), nil
	}

	x := Credentials{
		Validate: func() []error {
			return validate()
		},
		GetUsername: func() (string, []error) {
			return getUsername() 
		},
		GetPassword: func() (*string, []error) {
			return getPassword()
		},
		ToJSONString: func(json *strings.Builder) ([]error) {
			return getData().ToJSONString(json)
		},
	}

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}
