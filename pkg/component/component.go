package component

import (
	"bytes"
	"log"
	"syscall/js"
	"text/template"
)

type Component interface {
	GetId() string
	SetId(string)
	GetChildren() map[string]Component
	SetChild(key string, child Component)
	Render() string
	FuncMap() template.FuncMap
	OnChange(e js.Value)
	OnClick(e js.Value)
}

type BaseComponent struct {
	Id       string
	Children map[string]Component
}

func NewBaseComponent() Component {
	return &BaseComponent{
		Children: make(map[string]Component),
	}
}

// Render html template
func (dc *BaseComponent) Render() string {
	return `<div>
    {{ range $key, $value := .GetChildren }}
      {{ Generate $value }}
    {{ end }}
  </div>`
}

func (dc *BaseComponent) FuncMap() template.FuncMap {
	return template.FuncMap{
		"Generate": Generate,
	}
}

func Generate(c Component) string {
	funcMap := c.FuncMap()
	funcMap["Generate"] = Generate

	templateText := c.Render()
	tmpl, err := template.New("Render").Funcs(funcMap).Parse(templateText)
	if err != nil {
		return err.Error()
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, c)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func (dc *BaseComponent) GetChildren() map[string]Component {
	return dc.Children
}

func (dc *BaseComponent) SetChild(key string, child Component) {
	if dc.Children == nil {
		dc.Children = make(map[string]Component)
	}
	dc.Children[key] = child
}

func (dc *BaseComponent) GetId() string {
	return dc.Id
}

func (dc *BaseComponent) SetId(string id) {
	dc.Id = id
}

func (dc *BaseComponent) OnChange(e js.Value) {
	log.Printf("On Change e: %v\n", e)
}

func (dc *BaseComponent) OnClick(e js.Value) {
	log.Printf("On Click e: %v\n", e)
}
