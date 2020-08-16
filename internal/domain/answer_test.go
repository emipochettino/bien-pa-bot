package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAnswer(t *testing.T) {
	answer, err := NewAnswer("answer_test")

	assert.Nil(t, err)
	assert.NotNil(t, answer)
	assert.Equal(t, "answer_test", answer.GetText())
}
