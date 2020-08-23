package main

import (
	"crypto/tls"
	"fmt"
	infrastructure "github.com/emipochettino/bien-pa-bot/internal/infrastructure/adapters/providers"
	application "github.com/emipochettino/bien-pa-bot/internal/infrastructure/services"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if len(token) == 0 {
		panic(fmt.Errorf("TELEGRAM_TOKEN not set"))
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	bot, err := tgbotapi.NewBotAPIWithClient(token, client)
	if err != nil {
		panic(err)
	}

	bot.Debug = strings.EqualFold("dev", os.Getenv("PROFILE"))

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	greetingMessageProcessor := application.NewGreetingMessageProcessor(infrastructure.NewGreetingMessageAnswerProvider())
	incomingMessageProcessor := application.NewIncomingMessageProcessor(infrastructure.NewIncomingMessageAnswerProvider())
	messageServiceManager := application.NewMessageProcessorManager(greetingMessageProcessor, incomingMessageProcessor)
	messageService := application.NewMessageService(messageServiceManager)

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		innerUpdate := update
		go func() {
			if innerUpdate.Message == nil {
				return
			}
			answer, err := messageService.AnswerAMessage(innerUpdate.Message.Text)
			if err != nil {
				return
			}
			log.Printf("[%s] %s", innerUpdate.Message.From.UserName, innerUpdate.Message.Text)

			msg := tgbotapi.NewMessage(innerUpdate.Message.Chat.ID, answer.GetText())
			msg.ReplyToMessageID = innerUpdate.Message.MessageID
			bot.Send(msg)
		}()
	}
}
