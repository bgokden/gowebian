package components

import (
	"fmt"

	"github.com/bgokden/gowebian/pkg/component"
)

type UnorderedList struct {
	component.BaseComponent
}

func NewUnorderedList() *UnorderedList {
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

func NewListElement() *ListElement {
	e := &ListElement{}
	e.Tag = "il"
	return e
}

func NewTextElement(text string) component.Component {
	e := &ListElement{}
	e.Tag = "li"
	e.SetValue(text)
	return e
}
