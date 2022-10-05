package class

import (
	"strings"
	"fmt"
)

func CloneGrant(grant *Grant) *Grant {
	if grant == nil {
		return nil
	}

	return grant.Clone()
}

type Grant struct {
	Clone func() (*Grant)
	Validate func() ([]error)
	GetGrantValue func() (*string)
	GetFilter func() (*string)
	Grant func() (*string, []error) 
}

func NewGrant(client *Client, user *User, grantValue *string, filter *string) (*Grant, []error) {
	SQLCommand := newSQLCommand()

	ALL := func() string {
		return "ALL"
	}

	INSERT := func() string {
		return "INSERT"
	}

	UPDATE := func() string {
		return "UPDATE"
	}

	SELECT := func() string {
		return "SELECT"
	}
	
	GET_ALLOWED_GRANTS := func() Array {
		return Array{ALL(), INSERT(), UPDATE(), SELECT()}
	}
	
	data := Map {
		"client":Map{"type":"*Client","value":CloneClient(client),"mandatory":true},
		"user":Map{"type":"*User","value":CloneUser(user),"mandatory":true},		
		"grant":Map{"type":"*string","value":CloneString(grantValue),"mandatory":true,
		FILTERS(): Array{ Map {"values":GET_ALLOWED_GRANTS(),"function":ContainsExactMatch}}},
		"filter":Map{"type":"*string","value":CloneString(filter),"mandatory":true,
		FILTERS(): Array{ Map {"values":GetAllowedStringValues(),"function":ValidateCharacters}}},
	}

	validate := func() ([]error) {
		return ValidateGenericSpecial(data.Clone(), "Grant")
	}

	getClient := func () (*Client) {
		return CloneClient(data.M("client").GetObject("value").(*Client))
	}

	getUser := func () (*User) {
		return CloneUser(data.M("user").GetObject("value").(*User))
	}

	getGrantValue := func () (*string) {
		return CloneString(data.M("grant").S("value"))
	}

	getFilter := func () (*string) {
		return CloneString(data.M("filter").S("value"))
	}

	getSQL := func() (*string, []error) {
		errors := validate()
		if len(errors) > 0 {
			return nil, errors
		}

		database := getClient().GetDatabase()
		user := getUser()
		credentials := (*user).GetCredentials()
		domain_name := user.GetDomainName()

		sql := fmt.Sprintf("GRANT %s ON %s.%s To '%s'@'%s';", 
		*getGrantValue(), 
		(*(*database).GetDatabaseName()),  
		*getFilter(), 
		*((*credentials).GetUsername()),
		*((*domain_name).GetDomainName()))

		return &sql, nil
	}

	grant := func () (*string, []error) {
		var errors []error 
		sql_command, sql_command_errors := getSQL()
	
		if sql_command_errors != nil {
			return nil, sql_command_errors
		}

		stdout, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, true)
		
		if *stderr != "" {
			if strings.Contains(*stderr, "Operation CREATE USER failed for") {
				errors = append(errors, fmt.Errorf("create user failed most likely the user already exists"))
			} else {
				errors = append(errors, fmt.Errorf("unknown create user error" + *stderr))
			}
		}

		if len(errors) > 0 {
			return stdout, errors
		}

		return stdout, nil
	}	

	errors := validate()

	if errors != nil {
		return nil, errors
	}
	
	x := Grant{
		Validate: func() ([]error) {
			return validate()
		},
		GetGrantValue: func() (*string) {
			return getGrantValue()
		},
		GetFilter: func() (*string) {
			return getFilter()
		},
		Clone: func() (*Grant) {
			cloned, _ := NewGrant(getClient(), getUser(), getGrantValue(), getFilter())
			return cloned
		},
		Grant: func() (*string, []error) {
			return grant()
		},
    }

	return &x, nil
}