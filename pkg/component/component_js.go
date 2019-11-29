// +build js,wasm
package component

import (
	"log"
	"syscall/js"
)

func (bc *BaseComponent) Register(c Component) {
	// Events
	doc := js.Global().Get("document")
	element := doc.Call("getElementById", c.GetId())
	if element != js.Null() {
		callbacks := c.GetCallbacks()
		if callbacks != nil {
			for key, _ := range callbacks {
				onCallbackEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
					evt := args[0]
					value := evt.Get("target").Get("value")
					c.Callback(key, value)
					return nil
				})
				element.Call("addEventListener", key, onCallbackEvt)
				// defer onCallbackEvt.Release() todo Register this release
			}
		}
	} else {
		if c.GetId() != "body" {
			log.Printf("Couldn't find element %s\n", c.GetId())
		}
	}

	for _, value := range c.GetChildren() {
		value.Register(value)
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

func (bc *BaseComponent) SetPropertyWithId(id string, key string, value interface{}) {
	doc := js.Global().Get("document")
	element := doc.Call("getElementById", id)
	if element != js.Null() {
		element.Set(key, value)
	} else {
		if id != "body" {
			log.Printf("Couldn't find element %s\n", id)
		}
	}
}
