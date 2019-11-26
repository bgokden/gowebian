package components

import (
	"fmt"
	"log"

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

func (hc *TextLabel) OnChange(e interface{}) {
	log.Printf("e: %v\n", e)
}

func (hc *TextLabel) Render() string {
	return `<label id="{{.Id}}" for="...">{{.Value}}</label>`
}

func (hc *TextLabel) OnMessage(m *component.Message) {
	hc.SetProperty("innerHTML", fmt.Sprintf("%v", hc.Value))
	log.Printf("TextLabel m: %v\n", m.Value)
}
