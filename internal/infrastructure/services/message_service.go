package application

import (
	"github.com/emipochettino/bien-pa-bot/internal/domain"
)

type messageService struct {
	messageProcessorManager MessageProcessorManager
}

//AnswerAMessage validate the incoming text and return an answer if all goes well
func (m messageService) AnswerAMessage(text string) (domain.Answer, error) {
	inMessage, err := domain.NewInMessage(text)
	if err != nil {
		return nil, err
	}
	messageProcessor, err := m.messageProcessorManager.Get(inMessage.GetType())
	if err != nil {
		return nil, err
	}

	return messageProcessor.Process(inMessage)
}

type MessageService interface {
	AnswerAMessage(text string) (domain.Answer, error)
}

func NewMessageService(messageProcessorManager MessageProcessorManager) MessageService {
	return &messageService{messageProcessorManager: messageProcessorManager}
}
