package components

import (
	"github.com/bgokden/go-web-lib/pkg/component"
)

type P struct {
	component.BaseComponent
	Value string
}

func NewP(text string) *P {
	e := &P{
		Value: text,
	}
	e.Tag = "p"
	return e
}

func (p *P) Render() string {
	return `<{{.Tag}} id="{{.GetId}}">{{.Value}}</{{.Tag}}>`
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
	return `<{{.Tag}} id="{{.GetId}}" type="button">{{.Value}}</{{.Tag}}>`
}

type Text struct {
	component.BaseComponent
	Value string
}

func NewText(text string) *Text {
	t := &Text{
		Value: text,
	}
	t.SetTag("")
	return t
}

func (t *Text) Render() string {
	return `{{.Value}}`
}

type Title struct {
	component.BaseComponent
	Value string
}

func NewTitle(text string) *Title {
	c := &Title{
		Value: text,
	}
	c.SetTag("title")
	return c
}

func (c *Title) Render() string {
	return `<{{.Tag}}>{{.Value}}</{{.Tag}}>`
}

type Meta struct {
	component.BaseComponent
}

func NewMeta(attributeMap map[string]string) *Meta {
	c := &Meta{}
	c.SetTag("meta")
	if attributeMap != nil {
		for k, v := range attributeMap {
			c.SetAttribute(k, v)
		}
	}
	return c
}

func (c *Meta) Render() string {
	return `<{{.Tag}} {{ range $key, $value := .GetAttributes }} {{ printf "%s=\"%s\"" $key $value }} {{ end }}>`
}

type Script struct {
	component.BaseComponent
	Code string
}

func NewScript(attributeMap map[string]string, code string) *Script {
	c := &Script{}
	c.SetTag("script")
	if attributeMap != nil {
		for k, v := range attributeMap {
			c.SetAttribute(k, v)
		}
	}
	if code != "" {
		c.AddChild(NewText(code))
	}
	return c
}

func NewScriptFromSource(src string) *Script {
	return NewScript(map[string]string{
		"src": src,
	}, "")
}

func NewScriptFromCode(code string) *Script {
	return NewScript(nil, code)
}
