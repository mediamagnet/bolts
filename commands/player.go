package commands

import "github.com/pazuzu156/atlas"

type Player struct{ Command }

func InitPlayer() Player {
	return Player{Init(&CommandItem{
		Name:        "player",
		Description: "Plays videos from youtube",
		Aliases:     nil,
		Usage:       "]player play <video url>",
		Parameters:  []Parameter{
			{
				Name: "Function",
				Description: "Play/Pause/Stop/Skip",
				Required: true,
			},
			{
				Name: "VideoURL",
				Description: "Video to play",
				Required: false,
			},
		} ,
		Admin:       false,
	})}
}

func (c Player) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		ctx.Message.Reply(ctx.Context, ctx.Atlas, "Future content here.")
	}
	return c.CommandInterface
}
