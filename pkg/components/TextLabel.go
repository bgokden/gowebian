package components

import (
	"fmt"

	"github.com/bgokden/go-web-lib/pkg/component"
)

type TextLabel struct {
	component.BaseComponent
	Value string
}

func NewTextLabel() component.Component {
	return &TextLabel{
		Value: "...",
	}
}

func (tl *TextLabel) Render() string {
	return `<label id="{{.GetId}}" for="...">{{.Value}}</label>`
}

func (tl *TextLabel) OnMessage(m *component.Message) component.Component {
	tl.SetProperty("innerHTML", fmt.Sprintf("%v", m.Value))
	return tl
}
