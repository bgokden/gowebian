package component

import (
	"bytes"
	"text/template"
)

type Component interface {
	GetChildren() map[string]Component
	SetChild(key string, child Component)
	Render() string
}

type DefaultComponent struct {
	Children map[string]Component
}

func NewDefaultComponent() Component {
	return &DefaultComponent{
		Children: make(map[string]Component),
	}
}

// Render html template
func (dc *DefaultComponent) Render() string {
	return "<div></div>"
}

func Generate(component Component) string {
	funcMap := template.FuncMap{
		"Generate": Generate,
	}
	templateText := component.Render()
	tmpl, err := template.New("Render").Funcs(funcMap).Parse(templateText)
	if err != nil {
		return err.Error()
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, component)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func (dc *DefaultComponent) GetChildren() map[string]Component {
	return dc.Children
}

func (dc *DefaultComponent) SetChild(key string, child Component) {
	if dc.Children == nil {
		dc.Children = make(map[string]Component)
	}
	dc.Children[key] = child
}

type HelloComponent struct {
	DefaultComponent
}

func NewHelloComponent() Component {
	return &HelloComponent{}
}

func (hc *HelloComponent) Render() string {
	return "<div><p>Hello</p><div>"
}
