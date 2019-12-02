package components

import (
	"github.com/bgokden/gowebian/pkg/component"
)

type TextInput struct {
	component.BaseComponent
	Value string
}

func NewTextInput() *TextInput {
	return &TextInput{
		Value: "...",
	}
}

func (ti *TextInput) Render() string {
	return `<input id="{{.GetId}}" value="{{.Value}}" placeholder="Name input">`
}
