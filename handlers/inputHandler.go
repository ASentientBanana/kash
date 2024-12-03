package handlers

import (
	"os"
	"os/exec"
	"strings"

	"github.com/ASentientBanana/kash/definitions"
)

type ParsedInput struct {
	Cmd  string
	Args []string
}

func SplitInput(input string) ParsedInput {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	return ParsedInput{
		Args: args[1:],
		Cmd:  args[0],
	}
}

func ExecInput(input string) error {

	inp := SplitInput(input)

	//check for commands
	action := definitions.AvailableCommands[inp.Cmd]

	// check if actions is nil if yes pass as a commang to exec if no run a defined function

	if action == nil {
		cmd := exec.Command(inp.Cmd, inp.Args...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	// run a defined function
	// TODO(petar): add agument list and test
	return action(inp.Args)
}
