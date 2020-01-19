package main

import (
	"Bolts/commands"
	"log"
	"github.com/andersfylling/disgord"
	"github.com/joho/godotenv"
	"github.com/pazuzu156/atlas"
	"os"
)

var pingCommand = atlas.NewCommand("ping").SetDescription("Ping/Pong command")

func main() {
	// Load .env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := atlas.New(&atlas.Options{
		DisgordOptions: disgord.Config{
			BotToken: os.Getenv("DISCORD_TOKEN"),
			Logger: disgord.DefaultLogger(false),
		},
	})

	client.Use(atlas.DefaultLogger())
	client.GetPrefix = func(m *disgord.Message) string {
		return "]"
	}
	if err := client.Init(); err != nil {
		panic(err)
	}
}

func init() {
	atlas.Use(commands.InitPing().Register())
}