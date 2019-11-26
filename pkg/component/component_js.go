// +build js,wasm
package component

import (
	"log"
	"syscall/js"
)

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
	element := doc.Call("getElementById", c.GetId())
	if element != js.Null() {
		element.Call("addEventListener", "change", onChangeEvt)
	} else {
		if c.GetId() != "body" {
			log.Printf("Couldn't find element %s\n", c.GetId())
		}
	}
	for _, value := range c.GetChildren() {
		Register(value)
	}
}

func (bc *BaseComponent) SetProperty(key string, value interface{}) {
	doc := js.Global().Get("document")
	element := doc.Call("getElementById", bc.GetId())
	if element != js.Null() {
		element.Set(key, value)
	} else {
		if bc.GetId() != "body" {
			log.Printf("Couldn't find element %s\n", bc.GetId())
		}
	}
}
