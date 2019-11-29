package components

import (
	"log"

	"github.com/bgokden/go-web-lib/pkg/component"
	"github.com/bgokden/go-web-lib/pkg/events"
)

type TextInput struct {
	component.BaseComponent
	Value string
}

func NewTextInput() component.Component {
	return &TextInput{
		Value: "...",
	}
}

func (hc *TextInput) OnChange(e interface{}) component.Component {
	log.Printf("e: %v\n", e)
	events.Emit(&component.Message{
		From:  hc,
		To:    hc.GetParent().GetChild("label"),
		Title: "input Change",
		Value: e,
	})
	return hc
}

func (hc *TextInput) Render() string {
	return `<input id="{{.Id}}" value="{{.Value}}" placeholder="Name input">`
}
