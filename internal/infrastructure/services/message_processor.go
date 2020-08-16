package application

import (
	"github.com/emipochettino/bien-pa-bot/internal/domain"
	infrastructure "github.com/emipochettino/bien-pa-bot/internal/infrastructure/adapters/providers"
	"strings"
)

type MessageProcessor interface {
	shouldProcess(messageType string) bool
	Process(message domain.InMessage) (domain.Answer, error)
}

type GreetingMessageProcessor struct {
	answerProvider infrastructure.AnswerProvider
}

func (m *GreetingMessageProcessor) shouldProcess(messageType string) bool {
	return strings.EqualFold(domain.GreetingMessageType, messageType)
}

func (m *GreetingMessageProcessor) Process(message domain.InMessage) (domain.Answer, error) {
	return domain.NewAnswer(m.answerProvider.GetAnswer())
}

func NewGreetingMessageProcessor(answerProvider infrastructure.AnswerProvider) MessageProcessor {
	return &GreetingMessageProcessor{answerProvider: answerProvider}
}

type IncomingMessageProcessor struct {
	answerProvider infrastructure.AnswerProvider
}

func (i IncomingMessageProcessor) shouldProcess(messageType string) bool {
	return strings.EqualFold(domain.IncomingMessageType, messageType)
}

func (i IncomingMessageProcessor) Process(message domain.InMessage) (domain.Answer, error) {
	return domain.NewAnswer(i.answerProvider.GetAnswer())
}

func NewIncomingMessageProcessor(answerProvider infrastructure.AnswerProvider) MessageProcessor {
	return &IncomingMessageProcessor{answerProvider: answerProvider}
}
