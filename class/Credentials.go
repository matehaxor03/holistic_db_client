package class

import (
	"fmt"
)

type Credentials struct {
	Validate func() ([]error)
	GetUsername func() (*string)
	GetPassword func() (*string)
	GetCLSCommand func() (*string, []error)
	ToJSONString func() string 
	Clone func() *Credentials
}

func NewCredentials(username *string, password *string) (*Credentials) {
	
	getCredentialsUsernameValidCharacters := func() *string {
		temp := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890."
		return &temp
	}

	getCredentialPasswordValidCharacters := func() *string {
		temp := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890.="
		return &temp
	}
	
	data := Map {
		"username":Map{"type":"*string","value":CloneString(username),"mandatory":true,
		FILTERS(): Array{ Map {"values":getCredentialsUsernameValidCharacters(),"function":getValidateCharacters() }}},
		"password":Map{"type":"*string","value":CloneString(password),"mandatory":true,
		FILTERS(): Array{ Map {"values":getCredentialPasswordValidCharacters(),"function":getValidateCharacters() }}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Credentials")
	}

	getUsername := func () (*string) {
		return CloneString(data.M("username").S("value"))
	}

	getPassword := func () (*string) {
		return CloneString(data.M("password").S("value"))
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
		
			command := fmt.Sprintf("--user=%s --password=%s ", *(getUsername()), *(getPassword()))
		
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

