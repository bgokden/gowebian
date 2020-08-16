package components

import (
	"github.com/bgokden/gowebian/pkg/component"
)

type TextInput struct {
	component.BaseComponent
	Value string
}

func NewTextInput(text string) *TextInput {
	ti := &TextInput{}
	ti.SetTag("input")
	ti.SetValue(text)
	return ti
}

func (ti *TextInput) Render() string {
	return `<{{.Tag}} id="{{.GetId}}" value="{{.GetValue}}" placeholder="{{.GetValue}}">`
}
