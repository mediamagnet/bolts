package commands

import (
	"strings"
	"github.com/pazuzu156/atlas"
)

type Echo struct { Command }
var channel1 = ""
var channel2 = ""

func InitEcho() Echo {
	return Echo {Init(&CommandItem{
		Name:        "Echo",
		Description: "Echos one discord channel to another till stopped",
		Usage:       "]echo channelid1 channelid2",
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
		if strings.Contains(ctx.Message.Content, "]echo") == true {
			_, _ = ctx.Message.Reply(ctx.Context, ctx.Atlas, "Now echoing the channels you wanted")
		}
	}
	return c.CommandInterface
}