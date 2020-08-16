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
	ti.SetSelfClosing(true)
	ti.SetAttribute("value", text)
	ti.SetAttribute("placeholer", text)
	return ti
}
