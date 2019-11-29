//go:generate go run main.go
package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/bgokden/go-web-lib/pkg/component"
	"github.com/bgokden/go-web-lib/pkg/components"
	"github.com/bgokden/go-web-lib/pkg/page"
)

func main() {
	pg := page.NewBasePage()
	pg.SetTitle("Empty Page")
	pg.SetChild("label", components.NewTextLabel())
	pg.SetChild("input", components.NewTextInput())
	list := components.NewUnorderedList()
	textElement := components.NewTextElement("element 0")
	textElement.RegisterOnClick(func(e interface{}) {
		log.Println("todo: list component element pre re-render")
	})
	list.AddChild(textElement)
	list.AddChild(components.NewTextElement("element 1"))
	pg.SetChild("list", list)
	button := components.NewButton("click me!")
	button.RegisterOnClick(func(e interface{}) {
		list.AddChild(components.NewTextElement(fmt.Sprintf("element %d", rand.Int())))
		log.Println("todo: list component pre re-render")
		component.ReRender(list)
		list.Register(list)
	})
	pg.SetChild("button", button)
	pageString := component.Generate(pg)
	err := pg.Load(pageString)
	if err != nil {
		log.Fatal(err)
	}
}
