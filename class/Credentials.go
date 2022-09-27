package class

import (
	"fmt"
	"strings"
)

func GetCredentialsUsernameValidCharacters() string{
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890."
}

func GetCredentialPasswordValidCharacters() string {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890.="
}

type Credentials struct {
	Validate func() ([]error)
	GetUsername func() (*string)
	GetPassword func() (*string)
	GetCLSCommand func() (*string, []error)
	ToJSONString func() string 
	Clone func() *Credentials
}

func NewCredentials(username *string, password *string) (*Credentials) {
	username_copy := strings.Clone(*username)
	password_copy := strings.Clone(*password)
	
	data := Map {
		"username":Map{"type":"*string","value":&username_copy,"mandatory":true,
		FILTERS(): Array{ Map {"values":GetCredentialsUsernameValidCharacters(),"function":ValidateCharacters }}},
		"password":Map{"type":"*string","value":&password_copy,"mandatory":true,
		FILTERS(): Array{ Map {"values":GetCredentialPasswordValidCharacters(),"function":ValidateCharacters }}},
	}

	//(data.M("username").A("filters")[0]).(Map).SetFunc(ValidateCharacters)

	//panic(data.ToJSONString())

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Credentials")
	}

	getUsername := func () (*string) {
		ptr := data.M("username").S("value")
		if ptr == nil {
			return nil
		}
		cloneString := strings.Clone(*ptr)
		return &cloneString
	}

	getPassword := func () (*string) {
		ptr := data.M("password").S("value")
		if ptr == nil {
			return nil
		}
		cloneString := strings.Clone(*ptr)
		return &cloneString
	}
	
	x := Credentials{
		Validate: func() ([]error) {
			return validate()
		},
		GetUsername: func() (*string) {
			return getUsername()
		},
		GetPassword: func() (*string) {
			return getPassword()
		},
		GetCLSCommand: func() (*string, []error) {
			errors := validate()
			if errors != nil {
				return nil, errors
			}
		
			command := fmt.Sprintf("--user=%s --password=%s ", getUsername(), getPassword())
		
			return &command, nil
		 },
		ToJSONString: func() string {
			return data.Clone().ToJSONString()
		},
		Clone: func() *Credentials {
			return NewCredentials(getUsername(), getPassword())
		},
    }

	return &x
}

