package components

import (
	"log"
	"syscall/js"

	"github.com/bgokden/go-web-lib/pkg/component"
)

type HelloComponent struct {
	component.BaseComponent
	Name string
}

func NewHelloComponent() component.Component {
	return &HelloComponent{
		Name: "Hello",
	}
}

func (hc *HelloComponent) OnChange(e js.Value) {
	log.Println("e: %v\n", e)
}

func (hc *HelloComponent) Render() string {
	return `<div>
    <p>{{.Name}}</p>
    <input value="{{.Name}}" placeholder="Name input" autofocus>
  <div>`
}
