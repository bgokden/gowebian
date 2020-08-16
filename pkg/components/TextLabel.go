package components

import (
	"fmt"

	"github.com/bgokden/gowebian/pkg/component"
)

type TextLabel struct {
	component.BaseComponent
}

func NewTextLabel(text string) *TextLabel {
	tl := &TextLabel{}
	tl.SetTag("label")
	tl.SetValue(text)
	return tl
}

func (tl *TextLabel) OnMessage(m *component.Message) component.Component {
	tl.SetProperty("innerHTML", fmt.Sprintf("%v", m.Value))
	return tl
}
