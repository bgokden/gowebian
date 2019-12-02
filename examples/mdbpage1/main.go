//go:generate go run main.go
package main

import (
	"log"

	"github.com/bgokden/gowebian/mdb"
	"github.com/bgokden/gowebian/pkg/component"
)

func main() {
	pg := mdb.NewPage()
	pageString := component.Generate(pg)
	err := pg.Load(pageString)
	if err != nil {
		log.Fatal(err)
	}
}
