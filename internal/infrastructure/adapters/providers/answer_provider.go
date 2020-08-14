package infrastructure

import "math/rand"

type AnswerProvider interface {
	GetAnswer() string
}

type answerProvider struct {
	answers []string
}

func (p *answerProvider) GetAnswer() string {
	return p.answers[rand.Intn(len(p.answers))]
}

func NewIncomingMessageAnswerProvider() AnswerProvider {
	return &answerProvider{
		answers: []string{
			"bien pa vo pa?",
			"bieeeeeeeen pa vo pa?",
			"feró pa vo pa?",
			"ahí andamo pa vo pa?",
		},
	}
}

func NewGreetingMessageAnswerProvider() AnswerProvider {
	return &answerProvider{
		answers: []string{
			"Hola pa bien pa vo pa?",
		},
	}
}
