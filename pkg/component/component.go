package component

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/yosssi/gohtml"
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
	GetChildren() map[string]Component
	GetChildrenList() []Component
	HasChildren() bool
	SetChild(key string, child Component)
	GetChild(key string) Component
	AddChild(child Component)
	SetProperty(key string, value interface{})
	SetPropertyWithId(id string, key string, value interface{})
	Render() string
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
	GetValue() string
	SetValue(value string)
	GetPostValue() string
	SetPostValue(value string)
	IsSelfClosing() bool
	SetSelfClosing(bool)
}

type Message struct {
	From  Component
	To    Component
	Title string
	Value interface{}
}

type BaseComponent struct {
	JsBase
	Key              string
	Tag              string
	Value            string
	PostValue        string
	Parent           Component
	Children         map[string]Component
	ChildrenIndexMap []string
	Iterator         uint
	Callbacks        map[string]interface{}
	Attributes       map[string]string
	SelfClosing      bool
}

func NewBaseComponent() Component {
	return &BaseComponent{
		Tag:      "div",
		Children: make(map[string]Component),
		Iterator: 0,
	}
}

func (bc *BaseComponent) Render() string {
	var b strings.Builder
	fmt.Fprintf(&b, "<%v id=\"%v\"", bc.GetTag(), bc.GetId())
	for k, v := range bc.GetAttributes() {
		fmt.Fprintf(&b, " %v=\"%v\" ", k, v)
	}
	if bc.IsSelfClosing() {
		b.WriteString("/")
	}
	b.WriteString(">")
	if !bc.IsSelfClosing() {
		b.WriteString(bc.GetValue())
		for _, v := range bc.ChildrenIndexMap {
			b.WriteString(Generate(bc.Children[v]))
		}
		b.WriteString(bc.GetPostValue())
		fmt.Fprintf(&b, "</%v>", bc.GetTag())
	}
	return b.String()
}

func Generate(c Component) string {
	return gohtml.Format(c.Render())
}

func GenerateChild(c Component, key string) string {
	child := c.GetChild(key)
	if child == nil {
		return ""
	}
	return Generate(child)
}

func (bc *BaseComponent) GetChildren() map[string]Component {
	return bc.Children
}

func (bc *BaseComponent) GetChildrenList() []Component {
	list := make([]Component, 0, len(bc.Children))
	for _, v := range bc.ChildrenIndexMap {
		list = append(list, bc.Children[v])
	}
	return list
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
	if bc.ChildrenIndexMap == nil {
		bc.ChildrenIndexMap = make([]string, 0)
	}
	bc.ChildrenIndexMap = append(bc.ChildrenIndexMap, key)
	bc.Iterator++
}

func (bc *BaseComponent) GetChild(key string) Component {
	if bc.Children == nil {
		return nil
	}
	return bc.Children[key]
}

func (bc *BaseComponent) AddChild(child Component) {
	key := fmt.Sprintf("%d", bc.Iterator)
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
	go fnVal.Call(valIn)
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

func (bc *BaseComponent) GetValue() string {
	return bc.Value
}

func (bc *BaseComponent) SetValue(value string) {
	bc.Value = value
}

func (bc *BaseComponent) GetPostValue() string {
	return bc.PostValue
}

func (bc *BaseComponent) SetPostValue(value string) {
	bc.PostValue = value
}

func (bc *BaseComponent) IsSelfClosing() bool {
	return bc.SelfClosing
}

func (bc *BaseComponent) SetSelfClosing(value bool) {
	bc.SelfClosing = value
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
