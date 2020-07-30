package commands

import "github.com/pazuzu156/atlas"

// Roll I have a comment
type Roll struct{ Command }

// InitRoll Sets up command
func InitRoll() Roll {
	return Roll{Init(&CommandItem{
		Name:        "roll",
		Description: "Rolls dice.",
		Usage:       "]roll xdy+z>!afb",
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
func (c Roll) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {

	}
	return c.CommandInterface
}
