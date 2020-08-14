package domain

import "fmt"

type Answer interface {
	GetText() string
}

type answer struct {
	text string
}

func (o *answer) GetText() string {
	return o.text
}

func NewAnswer(text string) (Answer, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("the text should not be empty")
	}

	return &answer{text: text}, nil
}
