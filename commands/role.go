package commands

import (
	"Bolts/lib"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"go.mongodb.org/mongo-driver/bson"
)

type Role struct{ Command }

func InitRole() Role {
	return Role{Init(&CommandItem{
		Name:        "role",
		Description: "Give User a role when a phrase is said",
		Usage:       "]role, listenchannel, role, phrase",
		Parameters: []Parameter{
			{
				Name:        "listenchannel",
				Description: "Channel to listen for phrase in.",
				Required:    true,
			},
			{
				Name:        "Role",
				Description: "Role to give",
				Required:    true,
			},
			{
				Name:        "Phrase",
				Description: "Phrase to listen for.",
				Required:    true,
			},
		},
		Admin: false,
	})}
}

func (c Role) Register() *atlas.Command {
	var guildFound string
	// var role disgord.Snowflake
	// var channel disgord.Snowflake
	var phrase string
	var phraseSet string
	var role []disgord.Snowflake
	var role1 string

	c.CommandInterface.Run = func(ctx atlas.Context) {

		guildString := ctx.Message.GuildID.String()

		phraseLookUp := lib.MonReturnOneListen(lib.GetClient(), bson.M{"GuildID": guildString})

		if ctx.Message.GuildID.String() == phraseLookUp.GuildID {
			phraseSet = phraseLookUp.Phrase
			guildFound = phraseLookUp.GuildID
			role1 = phraseLookUp.RoleID

		}

		fmt.Println(phrase)
		msg := strings.TrimPrefix(ctx.Message.Content, "]role ")
		if msg == phraseSet {
			roleStr := lib.StrToSnowflake(role1)
			role = append(role, roleStr)
			fmt.Printf("Guild: %s, Role: %x, Phrase: %s", guildFound, role, phraseLookUp.Phrase)
			_ = disgord.GuildMemberUpdate{Roles: role}
			_, _ = ctx.Message.Reply(ctx.Context, ctx.Atlas, "Acknowledged")
		} else {
			if len(ctx.Args) > 0 {
				phrase := strings.TrimPrefix(ctx.Message.Content, "]Role")
				phrase1 := strings.Split(phrase, ", ")
				fmt.Println(ctx.Args[0])
				if ctx.Args[0] == "new" {
					replyPhrase := fmt.Sprintf("Watching for ]role %v", phrase1[len(phrase1)-1])
					ctx.Message.Reply(ctx.Context, ctx.Atlas, replyPhrase)
					fmt.Println(phrase)
					fmt.Printf("%v, %v, %v, %v", ctx.Message.GuildID, ctx.Args[0], ctx.Args[1], phrase1[len(phrase1)-1])
					listenInsert := lib.RoleMeListen{
						GuildID:   ctx.Message.GuildID.String(),
						ChannelID: ctx.Args[1],
						RoleID:    ctx.Args[2],
						Phrase:    phrase1[len(phrase1)-1],
					}
					lib.MonListen("bolts", "listens", listenInsert)
				} else {
					listenLookup := lib.MonReturnOneListen(lib.GetClient(), bson.M{"GuildID": ctx.Message.GuildID})
					fmt.Println(listenLookup)
					fmt.Println(listenLookup.Phrase)

				}
			} else {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "blah")
			}
		}
	}
	return c.CommandInterface
}
