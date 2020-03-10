package commands

import (
	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
)

type Help struct { Command }

func InitHelp() Help {
	return Help{Init(&CommandItem{
		Name:        "help",
		Description: "Help command",
		Usage:       "]help",
		Parameters: []Parameter{
			{
				Name:        "string",
				Description: "The help command",
				Required:    false,
			},
		},
	})}
}

func (c Help) Register() *atlas.Command {

	c.CommandInterface.Run = func(ctx atlas.Context) {
		ctx.Atlas.DeleteMessage(ctx.Context, ctx.Message.ChannelID, ctx.Message.ID)
		ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Fields: []*disgord.EmbedField{
					{Name: "Help",	Value: "Yer lookin at it."},
					{Name: "Tiny",	Value: "Shorten a url using TinyURL"},
					{Name: "Ping",  Value: "Echo's back what you said."},
				},
				Color: 0x00AAFF,
				Footer: &disgord.EmbedFooter{
					Text:         "Bolts: It's like a bucket of bolts in bot form.",
					IconURL:      "https://cdn.discordapp.com/app-icons/668208024862588928/102624a1cafd833ac883b94b71ba0e45.png?size=64",
					ProxyIconURL: "https://cdn.discordapp.com/app-icons/668208024862588928/102624a1cafd833ac883b94b71ba0e45.png?size=64",
				},
				Thumbnail: &disgord.EmbedThumbnail{
					URL:      "https://cdn.discordapp.com/app-icons/668208024862588928/102624a1cafd833ac883b94b71ba0e45.png?size=64",
					ProxyURL: "https://cdn.discordapp.com/app-icons/668208024862588928/102624a1cafd833ac883b94b71ba0e45.png?size=64",
					Height:   64,
					Width:    64,
				},
				Description: "Command character is ]",
				Author: &disgord.EmbedAuthor{
					Name:         "Bolts",
					URL:          "https://xkcd.com/no",
					IconURL:      "https://cdn.discordapp.com/app-icons/668208024862588928/102624a1cafd833ac883b94b71ba0e45.png?size=64",
					ProxyIconURL: "https://cdn.discordapp.com/app-icons/668208024862588928/102624a1cafd833ac883b94b71ba0e45.png?size=64",
				},
			},
		})
	}
	return c.CommandInterface
}

