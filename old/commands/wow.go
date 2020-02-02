package commands

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/jonas747/dca"
	"github.com/pazuzu156/atlas"
	"github.com/rylio/ytdl"

)

var volume = 100

type Wow struct { Command }
func PlayYT(videoURL string){
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 128
	options.Application = "lowdelay"
	options.Volume = volume

	videoInfo, err := ytdl.GetVideoInfo(videoURL)
	if err != nil {
		fmt.Printf("Error, could not find video. %s", err)
	}

	format := videoInfo.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
	downloadURL, err := videoInfo.GetDownloadURL(format)
	if err != nil {
		fmt.Printf("Error unable to download. %s", err)
	}

	encodingSession, err := dca.EncodeFile(downloadURL.String(), options)
	if err != nil {
		fmt.Printf("Error could not encode %s", err)
	}
	defer encodingSession.Cleanup()

	done := make(chan error)
	dca.NewStream(encodingSession, disgord.VoiceConnection(),done)
}

func InitWow() Wow {
	return Wow{Init(&CommandItem{
		Name:        "Wow",
		Description: "A little wow",
		Aliases:     nil,
		Usage:       "]wow",
		Parameters:  []Parameter{
			{
				Name: "String",
				Description: "wow",
				Required: false,
			},
		},
	})}
}

func (c Wow) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{

		})
	}
}
