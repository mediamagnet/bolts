package commands

import (

	"github.com/pazuzu156/atlas"

)

// Command is the base command object for all commands.
type Command struct {
	CommandInterface *atlas.Command
}

// CommandItem is the base command item object for the help command.
type CommandItem struct {
	Name        string
	Description string
	Aliases     []string
	Usage       string
	Parameters  []Parameter
	Admin       bool
}

// Parameter is the base parameter object for the help command.
type Parameter struct {
	Name        string // parameter name
	Value       string // value representation
	Description string // parameter description
	Required    bool   // is parameter required?
}

// Init initializes atlas commands
func Init(t *CommandItem) Command {
	cmd := atlas.NewCommand(t.Name).SetDescription(t.Description)

	if t.Aliases != nil {
		cmd.SetAliases(t.Aliases...)
	}

	return Command{cmd}
}

// embedFooter returns a footer and timestamp for disgord embeds

// getBot returns the bot object.
