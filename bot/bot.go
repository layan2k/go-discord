package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/layan2k/go-discord/config"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Running!")
}

func messageHandler(s *discordgo.Session, m discordgo.MessageCreate){

	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _= s.ChannelMessageSend(m.ChannelID, "pong")
	}
}

