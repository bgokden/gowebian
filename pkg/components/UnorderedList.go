package components

import (
	"fmt"

	"github.com/bgokden/go-web-lib/pkg/component"
)

type UnorderedList struct {
	component.BaseComponent
}

func NewUnorderedList() component.Component {
	e := &UnorderedList{}
	e.Tag = "ul"
	return e
}

func (ul *UnorderedList) OnMessage(m *component.Message) component.Component {
	ul.SetProperty("innerHTML", fmt.Sprintf("%v", m.Value))
	return ul
}

type ListElement struct {
	component.BaseComponent
}

func NewListElement() component.Component {
	e := &ListElement{}
	e.Tag = "li"
	return e
}

func NewTextElement(text string) component.Component {
	e := &ListElement{}
	e.Tag = "li"
	e.SetChild("0", NewP(text))
	return e
}
