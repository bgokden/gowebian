//go:generate go run main.go
package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/bgokden/gowebian/pkg/component"
	"github.com/bgokden/gowebian/pkg/components"
	"github.com/bgokden/gowebian/pkg/events"
	"github.com/bgokden/gowebian/pkg/page"
)

func main() {
	pg := page.NewBasePage()
	pg.SetTitle("Empty Page")
	textInput := components.NewTextInput("...")
	textlabel := components.NewTextLabel("enter input")
	textInput.SetCallback("change", func(e interface{}) {
		events.Emit(&component.Message{
			From:  textInput,
			To:    textlabel,
			Title: "input Change",
			Value: e,
		})
	})
	pg.SetChild("label", textlabel)
	pg.SetChild("input", textInput)
	textlabel.SetAttribute("for", textInput.GetId())
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
