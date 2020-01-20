package commands

import (
	"fmt"
	"github.com/pazuzu156/atlas"
	"github.com/subosito/shorturl"
	"reflect"
	"strings"
	"unsafe"
)

var provider = "tinyurl"

type Tiny struct { Command }

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

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

func (c Tiny) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		if len(ctx.Args) > 0 {
			fmt.Println(ctx.Message.Content)
			msg := strings.TrimPrefix(ctx.Message.Content, "]tiny ")
			u, err := shorturl.Shorten(msg, provider)
			if err == nil {
				fmt.Println(u)
			}
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Here's the URL you wanted "+" <@"+ctx.Message.Author.ID.String()+"> "+BytesToString([]byte(u)))

		} else {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Please Specify an URL to shorten")
		}
	}

	return c.CommandInterface
}