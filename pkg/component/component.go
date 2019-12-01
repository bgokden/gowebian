package component

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
	"text/template"

	"golang.org/x/net/html"
)

type Component interface {
	GetId() string
	GetKey() string
	SetKey(string)
	GetTag() string
	SetTag(string)
	SetParent(c Component)
	GetParent() Component
	SetHeader(key, value string)
	GetHeader(key string) string
	GetHeaders() map[string]string
	GetChildren() map[string]Component
	HasChildren() bool
	SetChild(key string, child Component)
	GetChild(key string) Component
	AddChild(child Component)
	SetProperty(key string, value interface{})
	SetPropertyWithId(id string, key string, value interface{})
	Render() string
	FuncMap() template.FuncMap
	OnMessage(message *Message) Component
	RegisterOnClick(callback interface{}) Component
	SetCallback(key string, callback interface{})
	GetCallback(key string) interface{}
	GetCallbacks() map[string]interface{}
	Callback(event string, args ...interface{}) Component
	Register(c Component)
	SetAttribute(key, value string)
	GetAttribute(key string) string
	GetAttributes() map[string]string
}

type Message struct {
	From  Component
	To    Component
	Title string
	Value interface{}
}

type BaseComponent struct {
	JsBase
	Key        string
	Tag        string
	Parent     Component
	Headers    map[string]string
	Children   map[string]Component
	Iterator   uint
	Callbacks  map[string]interface{}
	Attributes map[string]string
}

func NewBaseComponent() Component {
	return &BaseComponent{
		Tag:      "div",
		Children: make(map[string]Component),
		Iterator: 0,
	}
}

// Render html template
func (dc *BaseComponent) Render() string {
	return `<{{.GetTag}} id="{{.GetId}}" {{ range $key, $value := .GetAttributes }} {{ printf "%s=\"%s\"" $key $value }} {{ end }}>
	{{ range $key, $value := .GetChildren }}
  	{{ Generate $value }}
	{{ end }}
</{{.GetTag}}>`
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

func (bc *BaseComponent) HasChildren() bool {
	if bc.Children == nil {
		return false
	}
	return len(bc.Children) > 0
}

func (bc *BaseComponent) SetChild(key string, child Component) {
	if bc.Children == nil {
		bc.Children = make(map[string]Component)
	}
	child.SetKey(key)
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
	if bc.Parent != nil {
		return fmt.Sprintf("%s.%s", bc.Parent.GetId(), bc.Key)
	}
	return bc.Key
}

func (bc *BaseComponent) GetKey() string {
	return bc.Key
}

func (bc *BaseComponent) SetKey(key string) {
	bc.Key = key
}

func (bc *BaseComponent) GetTag() string {
	if bc.Tag == "" {
		return "div"
	}
	return bc.Tag
}

func (bc *BaseComponent) SetTag(tag string) {
	bc.Tag = tag
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

func (bc *BaseComponent) SetAttribute(key, value string) {
	if key != "id" {
		if bc.Attributes == nil {
			bc.Attributes = make(map[string]string)
		}
		bc.Attributes[key] = value
	}
}
func (bc *BaseComponent) GetAttribute(key string) string {
	if key == "id" {
		return bc.GetId()
	}
	if bc.Attributes == nil {
		return ""
	}
	return bc.Attributes[key]
}

func (bc *BaseComponent) GetAttributes() map[string]string {
	if bc.Attributes == nil {
		return make(map[string]string)
	}
	return bc.Attributes
}

func (bc *BaseComponent) Callback(event string, args ...interface{}) Component {
	fnVal := reflect.ValueOf(bc.GetCallback(event))
	valIn := make([]reflect.Value, len(args), len(args))
	for idx, elt := range args {
		valIn[idx] = reflect.ValueOf(elt)
	}
	fnVal.Call(valIn)
	// ReRender(bc)
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

func (bc *BaseComponent) GetCallbacks() map[string]interface{} {
	return bc.Callbacks
}

func ReRender(c Component) error {
	content := Generate(c)
	_, err := html.Parse(strings.NewReader(content))
	if err != nil {
		log.Println(err)
		return err
	}
	c.SetProperty("outerHTML", content)
	return nil
}

func NewComponent(tag, class, style string) Component {
	c := NewBaseComponent()
	c.SetTag(tag)
	c.SetAttribute("class", class)
	c.SetAttribute("style", style)
	return c
}
