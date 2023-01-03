package db_client

import (
	"fmt"
)

type Context struct {
	LogErrors func([]error)
	LogError  func(error)
	HasErrors func() bool
}

func NewContext() *Context {
	errors := []error{}

	getErrors := func() []error {
		return errors
	}

	logError := func(err error) {
		errors = append(errors, err)
		fmt.Println(err)
	}

	logErrors := func(errs []error) {
		for i := 0; i < len((errs)); i++ {
			err := errs[i]
			logError(err)
		}
	}

	hasErrors := func() bool {
		return len(getErrors()) > 0
	}

	return &Context{
		LogErrors: func(errs []error) {
			logErrors(errs)
		},
		LogError: func(err error) {
			logError(err)
		},
		HasErrors: func() bool {
			return hasErrors()
		},
	}
}
