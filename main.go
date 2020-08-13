package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if len(token) == 0 {
		panic(fmt.Errorf("TELEGRAM_TOKEN not set"))
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = strings.EqualFold("dev", os.Getenv("PROFILE"))

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !strings.Contains(update.Message.Text, "vo pa") {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, getAnswer())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}

func getAnswer() string {
	answers := []string{
		"bien pa vo pa?",
		"bieeeeeeeen pa vo pa?",
		"feró pa vo pa?",
		"ahí andamo pa vo pa?",
	}
	return answers[rand.Intn(len(answers)-1)]
}
