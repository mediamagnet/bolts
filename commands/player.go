package commands

import (
	"fmt"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/rylio/ytdl"
	"github.com/yyewolf/dca-disgord"

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
		var volume int

		fmt.Printf("Context args: %v, %v \n", ctx.Args[0], ctx.Args[1])
		// vChan, err := atlas.Disgord.VoiceConnect(ctx.Atlas.Disgord, ctx.Message.GuildID, setChan)

		vChan, err := disgord.Session.VoiceConnect(ctx.Atlas.Disgord, ctx.Message.GuildID, setChan)
		if err != nil {
			logrus.Debug(err)
		}

		switch doThing := ctx.Args[0]; doThing {
		case "join":
			err = vChan.MoveTo(setChan)
			if err != nil {
				logrus.Fatal(err)
			}
		case "leave":
			err = disgord.VoiceConnection.Close(vChan)
			if err != nil {
				logrus.Fatal(err)
			}
		case "volume":
			if ctx.Args[1] == "" {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "Volume is set to: "+strconv.Itoa(volume))
			} else {
				volume, _ = strconv.Atoi(ctx.Args[1])
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "Volume is set to: "+strconv.Itoa(volume))
			}

		case "play":

			// err = voice.StartSpeaking() // Sending a speaking signal is mandatory before sending voice data
			options := dca.StdEncodeOptions
			options.RawOutput = true
			options.Bitrate = 128
			options.Application = "lowdelay"
			options.Volume = volume

			videoInfo, err := ytdl.GetVideoInfo(ctx.Args[1])
			if err != nil {
				logrus.Fatal(err)
			}

			format := videoInfo.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
			downloadURL, err := videoInfo.GetDownloadURL(format)
			if err != nil {
				logrus.Error(err)
			}

			fmt.Println(downloadURL)

			encodingSession, err := dca.EncodeFile(downloadURL.String(), options)
			if err != nil {
				logrus.Error(err)
			}
			defer encodingSession.Cleanup()

			done := make(chan error)
			err = vChan.StartSpeaking()
			dca.NewStream(encodingSession, vChan, done)
			err = vChan.StopSpeaking()

			/*
				decoder := dca.NewDecoder()
				for {
					frame, err := decoder.OpusFrame()
					if err != nil {
						if err != io.EOF {
							logrus.Fatal(err)
						}
						break
					}
					vChan.SendOpusFrame(frame)
				}
				// err = vChan.SendDCA(f) // Or use voice.SendOpusFrame, this blocks until done sending (realtime audio duration)
				err = vChan.StopSpeaking() // Tell Discord we are done sending data.
				// err = vChan.Close() */
			if err != nil {
				logrus.Fatal(err)
			}
		}
	}
	return c.CommandInterface
}
