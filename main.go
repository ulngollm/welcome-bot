package main

import (
	"fmt"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"welcome-bot/settings"
)

var bot *tele.Bot

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("godotenv.Load: %s", err)
		return
	}

	botToken := os.Getenv("TOKEN")
	pref := tele.Settings{
		Token:     botToken,
		ParseMode: tele.ModeMarkdown,
	}
	bot, err = tele.NewBot(pref)
	if err != nil {
		log.Fatalf("tele.NewBot: %s", err)
		return
	}
}

func main() {
	bot.Handle("/start", joinHandler)
	bot.Handle(tele.OnUserJoined, joinHandler)

	bot.Start()
}

func joinHandler(c tele.Context) error {
	desc, err := settings.GetBotDescription(*c.Bot())
	if err != nil {
		return fmt.Errorf("getDescription: %w", err)
	}
	return c.Send(desc)
}
