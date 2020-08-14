package application

import (
	"fmt"
	"github.com/emipochettino/bien-pa-bot/internal/domain"
	providers "github.com/emipochettino/bien-pa-bot/internal/infrastructure/adapters/providers"
)

type messageService struct {
	greetingAnswerProvider providers.AnswerProvider
	incomingAnswerProvider providers.AnswerProvider
}

//AnswerAMessage validate the incoming text and return an answer if all goes well
func (m messageService) AnswerAMessage(text string) (domain.Answer, error) {
	inMessage, err := domain.NewInMessage(text)
	if err != nil {
		return nil, err
	}
	//TODO implement strategy pattern here
	switch inMessage.GetType() {
	case domain.GreetingMessageType:
		return domain.NewAnswer(m.greetingAnswerProvider.GetAnswer())
	case domain.IncomingMessageType:
		return domain.NewAnswer(m.incomingAnswerProvider.GetAnswer())
	default:
		return nil, fmt.Errorf("could not find an implementation for message type %s", inMessage.GetType())
	}
}

type MessageService interface {
	AnswerAMessage(text string) (domain.Answer, error)
}

func NewMessageService(
	greetingAnswerProvider providers.AnswerProvider,
	incomingAnswerProvider providers.AnswerProvider, ) MessageService {
	return &messageService{
		greetingAnswerProvider: greetingAnswerProvider,
		incomingAnswerProvider: incomingAnswerProvider}
}
