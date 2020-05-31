package commands

import (
	"fmt"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"

	// "github.com/rylio/ytdl"
	"github.com/sirupsen/logrus"
	// "github.com/yyewolf/dca-disgord"
)

// Player structs
type Player struct{ Command }

// InitPlayer sets up player command
func InitPlayer() Player {
	return Player{Init(&CommandItem{
		Name:        "player",
		Description: "Plays videos from youtube",
		Aliases:     nil,
		Usage:       "]player play <video url>",
		Parameters: []Parameter{
			{
				Name:        "Function",
				Description: "Play/Pause/Stop/Skip",
				Required:    true,
			},
			{
				Name:        "VideoURL",
				Description: "Video to play",
				Required:    false,
			},
		},
		Admin: false,
	})}
}

// Register player command
func (c Player) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		//Command stuff here
		setChan := disgord.NewSnowflake(676650064801824775)

		fmt.Printf("Context args: %v, %v \n", ctx.Args[0], ctx.Args[1])
		vChan, err := atlas.Disgord.VoiceConnect(ctx.Atlas.Disgord, ctx.Message.GuildID, setChan)
		if err != nil {
			logrus.Debug(err)
		}

		// var voice disgord.VoiceConnection

		switch doThing := ctx.Args[0]; doThing {
		case "join":
			err = vChan.MoveTo(setChan)
		case "leave":
			err = vChan.Close()
		case "play":

			f, err := os.Open("media/racestart.dca")
			defer f.Close()

			// err = voice.StartSpeaking() // Sending a speaking signal is mandatory before sending voice data
			err = vChan.StartSpeaking()
			err = vChan.SendDCA(f)     // Or use voice.SendOpusFrame, this blocks until done sending (realtime audio duration)
			err = vChan.StopSpeaking() // Tell Discord we are done sending data.
			err = vChan.Close()
			if err != nil {
				logrus.Fatal(err)
			}
		}
	}
	return c.CommandInterface
}
