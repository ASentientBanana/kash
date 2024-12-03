package definitions

import (
	"github.com/ASentientBanana/kash/definitions/commands"
)

type CommandType func([]string) error

var AvailableCommands = make(map[string]CommandType)

func InitializeCommandDefinitions() {
	AvailableCommands["cd"] = commands.Cd
}
