package main

import (
	"fmt"
	"github.com/layan2k/go-discord/config"
	"github.com/layan2k/go-discord/bot"
)
func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{}) 
	return
}