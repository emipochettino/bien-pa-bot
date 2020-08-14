package domain

import (
	"fmt"
	"strings"
)

type InMessage interface {
	GetType() string
}

type incomingMessage struct {
	text string
}

func (i *incomingMessage) GetType() string {
	return IncomingMessageType
}

type greetingMessage struct {
	text string
}

func (i *greetingMessage) GetType() string {
	return GreetingMessageType
}

func NewInMessage(text string) (InMessage, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("the should not be empty")
	}

	incomingMessageTexts := []string{
		"vo pa",
	}
	for _, incomingMessageText := range incomingMessageTexts {
		if strings.Contains(text, incomingMessageText) {
			return &incomingMessage{text: text}, nil
		}
	}
	greetingMessageTexts := []string{
		"ola pa",
		"onda pa",
	}
	for _, greetingMessageText := range greetingMessageTexts {
		if strings.Contains(text, greetingMessageText) {
			return &greetingMessage{text: text}, nil
		}
	}

	return nil, fmt.Errorf("the text does not contains what it is needed")
}
