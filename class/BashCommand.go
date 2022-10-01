package class


import (
	"bytes"
	"os/exec"
	"fmt"
)

type BashCommand struct {
	ExecuteUnsafeCommand func(command *string) (*string, *string, []error)
}

func newBashCommand() (*BashCommand) {			    
	x := BashCommand{
		ExecuteUnsafeCommand: func(command *string) (*string, *string, []error) {
			var errors []error 

			if command == nil {
				errors = append(errors, fmt.Errorf("bash command is nil"))
				return nil, nil, errors
			}	

			var stdout bytes.Buffer
			var stderr bytes.Buffer
			cmd := exec.Command("bash", "-c", *command)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			command_err := cmd.Run()

			shell_output := stdout.String()
			shell_output_errs := stderr.String()	
			
			if command_err != nil {
				errors = append(errors, command_err)	
				return &shell_output, &shell_output_errs, errors
			}

			return &shell_output, &shell_output_errs, nil
		},
    }

	return &x
}