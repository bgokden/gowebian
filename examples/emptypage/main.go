//go:generate go run main.go
package main

import (
	"github.com/bgokden/go-web-lib/pkg/component"
	"github.com/bgokden/go-web-lib/pkg/components"
	"github.com/bgokden/go-web-lib/pkg/page"
)

func main() {
	pg := page.NewBasePage()
	pg.SetTitle("Empty Page")
	pg.SetChild("label", components.NewTextLabel())
	pg.SetChild("input", components.NewTextInput())
	pageString := component.Generate(pg)
	pg.Load(pageString)
}
