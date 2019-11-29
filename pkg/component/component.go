package component

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
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
	GetChild(key string) Component
	AddChild(child Component)
	SetProperty(key string, value interface{})
	Render() string
	FuncMap() template.FuncMap
	OnChange(e interface{}) Component
	OnClick(e ...interface{}) Component
	OnMessage(message *Message) Component
	RegisterOnClick(callback interface{}) Component
	SetCallback(key string, callback interface{})
	GetCallback(key string) interface{}
}

type Message struct {
	From  Component
	To    Component
	Title string
	Value interface{}
}

type BaseComponent struct {
	JsBase
	Id        string
	Tag       string
	Parent    Component
	Headers   map[string]string
	Children  map[string]Component
	Iterator  uint
	Callbacks map[string]interface{}
}

func NewBaseComponent() Component {
	return &BaseComponent{
		Children: make(map[string]Component),
		Iterator: 0,
	}
}

// Render html template
func (dc *BaseComponent) Render() string {
	return `<{{.Tag}} id="{{.Id}}">
	{{ range $key, $value := .GetChildren }}
  	{{ Generate $value }}
	{{ end }}
</{{.Tag}}>`
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
	child.SetId(bc.GetId() + "." + key)
	child.SetParent(bc)
	bc.Children[key] = child
}

func (bc *BaseComponent) GetChild(key string) Component {
	if bc.Children == nil {
		return nil
	}
	return bc.Children[key]
}

func (bc *BaseComponent) AddChild(child Component) {
	key := fmt.Sprintf("%d", bc.Iterator)
	bc.Iterator++
	bc.SetChild(key, child)
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

func (bc *BaseComponent) OnChange(e interface{}) Component {
	log.Printf("On Change e: %v\n", e)
	return bc
}

func (bc *BaseComponent) OnClick(args ...interface{}) Component {
	fnVal := reflect.ValueOf(bc.GetCallback("click"))
	valIn := make([]reflect.Value, len(args), len(args))
	for idx, elt := range args {
		valIn[idx] = reflect.ValueOf(elt)
	}
	fnVal.Call(valIn)
	return bc
}

func (bc *BaseComponent) OnMessage(message *Message) Component {
	log.Printf("On Message m: %v\n", message)
	return bc
}

func (bc *BaseComponent) RegisterOnClick(callback interface{}) Component {
	bc.SetCallback("click", callback)
	return bc
}

func (bc *BaseComponent) SetCallback(key string, callback interface{}) {
	if bc.Callbacks == nil {
		bc.Callbacks = make(map[string]interface{})
	}
	bc.Callbacks[key] = callback
}

func (bc *BaseComponent) GetCallback(key string) interface{} {
	if bc.Callbacks == nil {
		return func(value interface{}) {
			log.Printf("value: %v\n", value)
		}
	}
	return bc.Callbacks[key]
}

type JsBase struct{}

func (jb *JsBase) SetProperty(key string, value interface{}) {
	log.Printf("JsBase SetProperty: %v\n", value)
}
