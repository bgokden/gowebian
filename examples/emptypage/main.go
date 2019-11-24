package main

import (
	"fmt"
	"syscall/js"

	"github.com/bgokden/go-web-lib/pkg/component"
	"github.com/bgokden/go-web-lib/pkg/page"
)

func main() {
	// var isGenenrate = flag.Bool("generate", false, "Generate page code")
	fmt.Println("main wasm run")
	pg := page.NewDefaultPage()
	pg.SetTitle("Empty Page")
	pg.SetChild("hello", component.NewHelloComponent())
	pageString := component.Generate(pg)
	fmt.Printf("main wasm run 2 %v\n", len(pageString))
	fmt.Println(pageString)

	doc := js.Global().Get("document")
	doc.Call("write", pageString)
}
