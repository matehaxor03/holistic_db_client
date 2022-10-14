package class

func CloneCredentials(credentials *Credentials) *Credentials {
	if credentials == nil {
		return credentials
	}

	return credentials.Clone()
}

type Credentials struct {
	Validate func() ([]error)
	GetUsername func() (*string)
	GetPassword func() (*string)
	ToJSONString func() string 
	Clone func() *Credentials
}

func GetCredentialsUsernameValidCharacters() *string {
	temp := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	return &temp
}

func GetCredentialPasswordValidCharacters() *string {
	temp := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890.=_"
	return &temp
}

func NewCredentials(username *string, password *string) (*Credentials, []error) {
	
	data := Map {
		"[username]":Map{"value":CloneString(username),"mandatory":true,
		FILTERS(): Array{ Map {"values":GetCredentialsUsernameValidCharacters(),"function":getWhitelistCharactersFunc() }}},
		"[password]":Map{"value":CloneString(password),"mandatory":true,
		FILTERS(): Array{ Map {"values":GetCredentialPasswordValidCharacters(),"function":getWhitelistCharactersFunc() }}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Credentials")
	}

	getUsername := func () (*string) {
		return CloneString(data.M("[username]").S("value"))
	}

	getPassword := func () (*string) {
		return CloneString(data.M("[password]").S("value"))
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
		ToJSONString: func() string {
			return data.Clone().ToJSONString()
		},
		Clone: func() *Credentials {
			cloned, _ :=  NewCredentials(getUsername(), getPassword())
			return cloned
		},
    }

	errors := validate()

	if errors != nil {
		return nil, errors
	}

	return &x, nil
}

