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
	SetParent(c Component)
	GetParent() Component
	SetHeader(key, value string)
	GetHeader(key string) string
	GetHeaders() map[string]string
	GetChildren() map[string]Component
	SetChild(key string, child Component)
	Render() string
	FuncMap() template.FuncMap
	OnChange(e js.Value)
	OnClick(e js.Value)
}

type BaseComponent struct {
	Id       string
	Parent   Component
	Headers  map[string]string
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

func (bc *BaseComponent) GetChildren() map[string]Component {
	return bc.Children
}

func (bc *BaseComponent) SetChild(key string, child Component) {
	if bc.Children == nil {
		bc.Children = make(map[string]Component)
	}
	child.SetParent(bc)
	bc.Children[key] = child
}

func (bc *BaseComponent) GetId() string {
	return bc.Id
}

func (bc *BaseComponent) SetId(id string) {
	bc.Id = id
}

func (bc *BaseComponent) GetParent() Component {
	return bc.Parent
}

func (bc *BaseComponent) SetParent(c Component) {
	bc.Parent = c
}

func (bc *BaseComponent) SetHeader(key, value string) {
	if bc.Headers == nil {
		bc.Headers = make(map[string]string)
	}
	bc.Headers[key] = value
}

func (bc *BaseComponent) GetHeader(key string) string {
	if bc.Headers == nil {
		return ""
	}
	return bc.Headers[key]
}

func (bc *BaseComponent) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for _, child := range bc.GetChildren() {
		for key, value := range child.GetHeaders() {
			headers[key] = value
		}
	}
	for key, value := range bc.Headers {
		headers[key] = value
	}
	return headers
}

func (dc *BaseComponent) OnChange(e js.Value) {
	log.Printf("On Change e: %v\n", e)
}

func (dc *BaseComponent) OnClick(e js.Value) {
	log.Printf("On Click e: %v\n", e)
}
