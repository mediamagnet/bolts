package commands

import "github.com/pazuzu156/atlas"

// Man I have a comment
type Man struct{ Command }

// InitMan Sets up command
func InitMan() Man {
	return Man{Init(&CommandItem{
		Name:        "man",
		Description: "Returns a man page for the searched command.",
		Usage:       "]man command",
		Parameters: []Parameter{
			{
				Name:        "string",
				Description: "Command String",
				Required:    true,
			},
		},
	})}
}

// Register Man
func (c Man) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {

	}
	return c.CommandInterface
}
