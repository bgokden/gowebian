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

func (hc *TextLabel) Render() string {
	return `<label id="{{.Id}}" for="...">{{.Value}}</label>`
}

func (hc *TextLabel) OnMessage(m *component.Message) component.Component {
	hc.SetProperty("innerHTML", fmt.Sprintf("%v", m.Value))
	return hc
}
