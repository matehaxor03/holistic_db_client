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
	Grant func() ([]error) 
}

func NewGrant(client *Client, user *User, grant_value *string, filter *string) (*Grant, []error) {
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
		"client":Map{"value":CloneClient(client),"mandatory":true},
		"user":Map{"value":CloneUser(user),"mandatory":true},		
		"grant":Map{"value":CloneString(grant_value),"mandatory":true,
		FILTERS(): Array{ Map {"values":GET_ALLOWED_GRANTS(),"function":getWhitelistStringFunc()}}},
		"filter":Map{"value":CloneString(filter),"mandatory":true},
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

		database := (*(getClient())).GetDatabase()
		user := getUser()
		credentials := (*user).GetCredentials()
		domain_name := (*user).GetDomainName()

		grant_value := *(getGrantValue())
		filter_value := *(getFilter())
		username_value := *((*credentials).GetUsername())
		domain_name_value := *((*domain_name).GetDomainName())
		database_name_value := *((*database).GetDatabaseName())

		sql := fmt.Sprintf("GRANT %s ON %s.%s To '%s'@'%s';", grant_value, database_name_value, filter_value, username_value, domain_name_value)

		return &sql, nil
	}

	grant := func () ([]error) {
		var errors []error 
		sql_command, sql_command_errors := getSQL()
	
		if sql_command_errors != nil {
			return sql_command_errors
		}

		_, stderr, errors := SQLCommand.ExecuteUnsafeCommand(getClient(), sql_command, Map{"use_file": true})
		
		if stderr != nil && *stderr != "" {
			if strings.Contains(*stderr, "Operation CREATE USER failed for") {
				errors = append(errors, fmt.Errorf("create user failed most likely the user already exists"))
			} else {
				errors = append(errors, fmt.Errorf(*stderr))
			}
		}

		if len(errors) > 0 {
			return errors
		}

		return nil
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
		Grant: func() ([]error) {
			return grant()
		},
    }

	return &x, nil
}