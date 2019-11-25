package component

import (
	"bytes"
	"fmt"
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
	GetChild(key string) Component
	SetProperty(key string, value interface{})
	Render() string
	FuncMap() template.FuncMap
	OnChange(e interface{})
	OnClick(e interface{})
	OnMessage(message *Message)
}

type Message struct {
	From  Component
	To    Component
	Title string
	Value interface{}
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
	return `<div id="{{.Id}}">
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

func Register(c Component) {
	onChangeEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		evt := args[0]
		value := evt.Get("target").Get("value")
		c.OnChange(value)
		return nil
	})
	// defer onChangeEvt.Release()
	// Events
	doc := js.Global().Get("document")
	fmt.Println(c.GetId())
	element := doc.Call("getElementById", c.GetId())
	if element != js.Null() {
		element.Call("addEventListener", "change", onChangeEvt)
	} else {
		log.Printf("Couldn't find element %s\n", c.GetId())
	}
	for _, value := range c.GetChildren() {
		Register(value)
	}
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

func (bc *BaseComponent) OnChange(e interface{}) {
	log.Printf("On Change e: %v\n", e)
}

func (bc *BaseComponent) OnClick(e interface{}) {
	log.Printf("On Click e: %v\n", e)
}

func (bc *BaseComponent) OnMessage(message *Message) {
	log.Printf("On Message m: %v\n", message)
}

func (bc *BaseComponent) SetProperty(key string, value interface{}) {
	doc := js.Global().Get("document")
	fmt.Println(bc.GetId())
	element := doc.Call("getElementById", bc.GetId())
	if element != js.Null() {
		element.Set(key, value)
	}
}

func (bc *BaseComponent) GetSelfElement() js.Value {
	doc := js.Global().Get("document")
	fmt.Println(bc.GetId())
	element := doc.Call("getElementById", bc.GetId())
	if element != js.Null() {
		return element
	}
	return js.Value{}
}
