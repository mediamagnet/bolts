package commands

import (
	"fmt"
	"strings"

	"github.com/pazuzu156/atlas"
	"github.com/subosito/shorturl"
)

var provider = "tinyurl"

// Tiny struct
type Tiny struct{ Command }

// InitTiny sets up tiny
func InitTiny() Tiny {
	return Tiny{Init(&CommandItem{
		Name:        "tiny",
		Description: "generate a short url using bitly",
		Usage:       "]tiny <long url>",
		Parameters: []Parameter{
			{
				Name:        "string",
				Description: "The url to be shortened",
				Required:    true,
			},
		},
	})}
}

// Register registers tiny with atlas.
func (c Tiny) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		if len(ctx.Args) > 0 {
			fmt.Println(ctx.Message.Content)
			msg := strings.TrimPrefix(ctx.Message.Content, "]tiny ")
			u, err := shorturl.Shorten(msg, provider)
			if err == nil {
				fmt.Println(u)
			}
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Here's the URL you wanted "+" <@"+ctx.Message.Author.ID.String()+"> "+string([]byte(u)))

		} else {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Please Specify an URL to shorten")
		}
	}

	return c.CommandInterface
}
