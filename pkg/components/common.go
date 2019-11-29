package components

import (
	"github.com/bgokden/go-web-lib/pkg/component"
)

type P struct {
	component.BaseComponent
	Value string
}

func NewP(text string) component.Component {
	e := &P{
		Value: text,
	}
	e.Tag = "p"
	return e
}

func (p *P) Render() string {
	return `<{{.Tag}} id="{{.Id}}">{{.Value}}</{{.Tag}}>`
}

type Button struct {
	component.BaseComponent
	Value string
}

func NewButton(text string) component.Component {
	e := &Button{
		Value: text,
	}
	e.Tag = "button"
	return e
}

func (b *Button) Render() string {
	return `<{{.Tag}} id="{{.Id}}" type="button">{{.Value}}</{{.Tag}}>`
}
