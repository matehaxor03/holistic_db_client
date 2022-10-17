package class

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type BashCommand struct {
	ExecuteUnsafeCommand func(command *string) (*string, []error)
}

func newBashCommand() *BashCommand {
	x := BashCommand{
		ExecuteUnsafeCommand: func(command *string) (*string, []error) {
			var errors []error

			if command == nil {
				errors = append(errors, fmt.Errorf("bash command is nil"))
			}

			if len(errors) > 0 {
				return nil, errors
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
			}

			if strings.TrimSpace(shell_output_errs) != "" {
				errors = append(errors, fmt.Errorf(shell_output_errs))
			}

			if len(errors) > 0 {
				return &shell_output, errors
			}

			return &shell_output, nil
		},
	}

	return &x
}
