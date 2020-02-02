package main

import (
	"log"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsToken := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + dsToken)
	if err != nil {
		log.Fatal("Error creating Discord Session,", err)
	}

	//message handles
	dg.AddHandler(messageCreate)

	//connect handle
	dg.AddHandler(connect)

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection,", err)
		return
	}
	fmt.Println("Bolts is connected press ctrl+c to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc

	_ = dg.Close()
}

func connect(s *discordgo.Session, c *discordgo.Connect){
	log.Print(c)
	var guildName = make([]string, 1)
	for _, v := range s.State.Guilds {
		guildName = append(guildName, v.Name)
	}
	for {
		guildCount := len(guildName)-1
		err := s.UpdateListeningStatus(fmt.Sprintf("commands in %v servers", guildCount))
		time.Sleep(15 * time.Minute)
		err = s.UpdateListeningStatus("Donnybrook talk about cosmic background radiation.")
		time.Sleep(15 * time.Minute)
		err = s.UpdateStatus(0, "Bolts v0.1")
		time.Sleep(15 * time.Minute)
		if err != nil {
			log.Println(err)
		}
	}

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Ignore self
	if m.Author.ID == s.State.User.ID {
		return
	}
}