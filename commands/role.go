package commands

import (
	"Bolts/lib"
	"context"
	"fmt"
	"strings"

	"github.com/pazuzu156/atlas"
	"go.mongodb.org/mongo-driver/bson"
)

// Role Struct
type Role struct{ Command }

// InitRole sets up the role command
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

// Register command
func (c Role) Register() *atlas.Command {
	// var guildFound string
	// var role disgord.Snowflake
	// var channel disgord.Snowflake
	var phrase = ""
	var roleID = ""

	c.CommandInterface.Run = func(ctx atlas.Context) {

		guildString := ctx.Message.GuildID.String()

		phraseLookUp := lib.MonReturnOneListen(lib.GetClient(), bson.M{"GuildID": guildString})

		msg := strings.TrimPrefix(ctx.Message.Content, "]role ")

		fmt.Printf("MSG: %s \n", msg)

		phrases := lib.MonReturnAllListen(lib.GetClient(), bson.M{})
		if ctx.Message.GuildID.String() == phraseLookUp.GuildID {
			for _, v := range phrases {
				if v.Phrase == msg {
					phrase = v.Phrase
					roleID = v.RoleID
				}
			}
		}

		fmt.Printf("Phrase: %s \n", phrase)

		fmt.Println(phrase)
		if msg == phrase {
			fmt.Printf("Role: %s \n", roleID)

			roleStr := lib.StrToSnowflake(roleID)
			fmt.Printf("Snowflake: %v, Converted: %v \n", roleStr, lib.SnowflakeToUInt64(roleStr))
			_ = atlas.Disgord.AddGuildMemberRole(ctx.Atlas.Disgord, context.Background(), ctx.Message.GuildID, ctx.Message.Author.ID, roleStr)
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
						ChannelID: strings.TrimSuffix(ctx.Args[1], ","),
						RoleID:    strings.TrimSuffix(ctx.Args[2], ","),
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
