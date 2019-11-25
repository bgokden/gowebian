package main

import (
	"fmt"
	"syscall/js"

	"github.com/bgokden/go-web-lib/pkg/component"
	"github.com/bgokden/go-web-lib/pkg/components"
	"github.com/bgokden/go-web-lib/pkg/events"
	"github.com/bgokden/go-web-lib/pkg/page"
)

func main() {
	// var isGenenrate = flag.Bool("generate", false, "Generate page code")
	fmt.Println("main wasm run")
	pg := page.NewBasePage()
	pg.SetTitle("Empty Page")
	pg.SetChild("label", components.NewTextLabel())
	pg.SetChild("input", components.NewTextInput())
	pageString := component.Generate(pg)
	fmt.Printf("main wasm run 2 %v\n", len(pageString))
	fmt.Println(pageString)

	doc := js.Global().Get("document")
	doc.Call("write", pageString)
	component.Register(pg)
	events.Listen()
}
