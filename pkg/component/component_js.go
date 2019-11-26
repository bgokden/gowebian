// +build js,wasm
package component

import (
	"fmt"
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

func (bc *BaseComponent) SetProperty(key string, value interface{}) {
	doc := js.Global().Get("document")
	fmt.Println(bc.GetId())
	element := doc.Call("getElementById", bc.GetId())
	if element != js.Null() {
		element.Set(key, value)
	}
}
