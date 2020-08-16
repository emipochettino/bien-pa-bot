package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGreetingMessageAnswerProvider(t *testing.T) {
	answerProvider := NewGreetingMessageAnswerProvider()
	assert.NotNil(t, answerProvider)
	answer := answerProvider.GetAnswer()
	assert.NotNil(t, answer)
	assert.True(t, len(answer) > 0)
}

func TestNewIncomingMessageAnswerProvider(t *testing.T) {
	answerProvider := NewIncomingMessageAnswerProvider()
	assert.NotNil(t, answerProvider)
	answer := answerProvider.GetAnswer()
	assert.NotNil(t, answer)
	assert.True(t, len(answer) > 0)
}
