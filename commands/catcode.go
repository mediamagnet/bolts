package commands

import (
	"fmt"
	"github.com/pazuzu156/atlas"
)

type CatCode struct{ Command }

func InitCatCode() CatCode {
	return CatCode{Init(&CommandItem{
		Name:        "catcode",
		Description: "Returns a cat picture for a http response code",
		Aliases: []string{"cat"},
		Usage:       "]cat 404",
		Parameters: []Parameter{
			{
				Name:        "string",
				Description: "The HTTP response code.",
				Required:    true,
			},
		},
	})}
}

func (c CatCode) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {

		httpcode := ctx.Args[0]

		_, _ = ctx.Message.Reply(ctx.Context, ctx.Atlas, fmt.Sprintf("https://http.cat/%v.jpg", httpcode))
	}
	return c.CommandInterface
}