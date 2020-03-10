package commands

import "github.com/pazuzu156/atlas"

type Echo struct { Command }
channel1 := ""
channel2 := ""

func InitEcho() Echo {
	return Echo {Init(&CommandItem{
		Name:        "Echo",
		Description: "Echos one discord channel to another till stopped",
		Usage:       "]copy channelid1 channelid2",
		Parameters:  []Parameter{
			{
				Name:        "Channel ID 1",
				Description: "Channel ID for the source Channel",
				Required:    true,
			},
			{
				Name:        "Channel ID 2",
				Description: "Channel ID for the channel you send to.",
				Required:    true,
			},
		},
		Admin:       true,
	})}
}

func (c Echo) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		ctx.Message.Reply(ctx.Context, ctx.Atlas, )
	}
}