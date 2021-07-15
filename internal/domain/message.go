package domain

import (
	"fmt"
	"os"
	"strings"
)

const (
	IncomingMessageType    = "IncomingMessageType"
	GreetingMessageType    = "GreetingMessageType"
	VaccinationMessageType = "VaccinationMessageType"
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

type vaccinationMessage struct {
	text string
}

func (i *vaccinationMessage) GetType() string {
	return VaccinationMessageType
}

func NewInMessage(text string, name string) (InMessage, error) {
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
		"bien pa",
	}
	for _, greetingMessageText := range greetingMessageTexts {
		if strings.Contains(text, greetingMessageText) {
			return &greetingMessage{text: text}, nil
		}
	}
	vaccinationMessageTexts := []string{
		"sputnik",
		"la rusa",
		"moderna",
		"pfizer",
		"dosis",
		"la china",
		"astrazeneca",
		"vacuna",
		"fiebre",
		"virus",
		"inmunidad",
		"covid",
		"corona",
		"coronavirus",
	}
	toAnswers := strings.Split(os.Getenv("TO_ANSWER_LIST"), ",")
	for _, toAnswer := range toAnswers {
		if toAnswer == name {
			for _, vaccinationMessageText := range vaccinationMessageTexts {
				if strings.Contains(strings.ToLower(text), vaccinationMessageText) {
					return &vaccinationMessage{text: text}, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("the text does not contains what it is needed")
}
