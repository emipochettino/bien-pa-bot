package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInMessageWithTypeGreeting(t *testing.T) {
	inMessage, err := NewInMessage("Hola pa", "user_test")

	assert.Nil(t, err)
	assert.NotNil(t, inMessage)
	assert.Equal(t, GreetingMessageType, inMessage.GetType())
}

func TestNewInMessageWithTypeIncoming(t *testing.T) {
	inMessage, err := NewInMessage("Bien pa vo pa", "user_test")

	assert.Nil(t, err)
	assert.NotNil(t, inMessage)
	assert.Equal(t, IncomingMessageType, inMessage.GetType())
}

func TestNewInMessageReturnErrorForInvalidText(t *testing.T) {
	inMessage, err := NewInMessage("hola que tal", "user_test")

	assert.Nil(t, inMessage)
	assert.NotNil(t, err)
}