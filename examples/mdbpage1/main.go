//go:generate go run main.go
package main

import (
	"log"

	"github.com/bgokden/go-web-lib/mdb"
	"github.com/bgokden/go-web-lib/pkg/component"
)

func main() {
	pg := mdb.NewPage()
	pageString := component.Generate(pg)
	err := pg.Load(pageString)
	if err != nil {
		log.Fatal(err)
	}
}
