package components

import (
	"github.com/bgokden/gowebian/pkg/component"
)

type P struct {
	component.BaseComponent
}

func NewP(text string) *P {
	p := &P{}
	p.SetTag("p")
	p.SetValue(text)
	return p
}

type Button struct {
	component.BaseComponent
}

func NewButton(text string) component.Component {
	b := &Button{}
	b.SetTag("button")
	b.SetAttribute("type", "button")
	b.SetValue(text)
	return b
}

type Title struct {
	component.BaseComponent
	Value string
}

func NewTitle(text string) *Title {
	t := &Title{
		Value: text,
	}
	t.SetTag("title")
	return t
}

type Meta struct {
	component.BaseComponent
}

func NewMeta(attributeMap map[string]string) *Meta {
	m := &Meta{}
	m.SetTag("meta")
	if attributeMap != nil {
		for k, v := range attributeMap {
			m.SetAttribute(k, v)
		}
	}
	m.SetSelfClosing(true)
	return m
}

/*
func (m *Meta) Render() string {
	return `<{{.Tag}}{{ range $key, $value := .GetAttributes }} {{ printf "%s=\"%s\"" $key $value }} {{ end }}>`
}
*/

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
	c.SetValue(code)
	return c
}

/*
func (c *Script) Render() string {
	return `<{{.GetTag}}{{ range $key, $value := .GetAttributes }} {{ printf "%s=\"%s\"" $key $value }} {{ end }}>
	{{ .GetValue }}
	{{ range $key, $value := .GetChildren }}
  	{{ Generate $value }}
	{{ end }}
</{{.GetTag}}>`
}
*/

func NewScriptFromSource(src string) *Script {
	return NewScript(map[string]string{
		"src": src,
	}, "")
}

func NewScriptFromCode(code string) *Script {
	return NewScript(nil, code)
}
