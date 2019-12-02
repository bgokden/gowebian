//go:generate go run main.go
package main

import (
	"fmt"
	"log"

	"github.com/bgokden/gowebian/mdb"
	"github.com/bgokden/gowebian/pkg/component"
	"github.com/bgokden/gowebian/pkg/components"
)

func main() {
	pg := mdb.NewPage()
	pg.SetAttribute("class", "fixed-sn mdb-skin")
	header := mdb.NewHeader()
	pg.AddChild(header)
	main := mdb.NewMain()
	container := mdb.NewContainer()
	for i := 0; i < 3; i++ {
		row := mdb.NewRow()
		for j := 0; j < 3; j++ {
			col := mdb.NewCol()
			col.AddChild(components.NewText(fmt.Sprintf("row: %v - col: %v", i+1, j+1)))
			row.AddChild(col)
		}
		container.AddChild(row)
	}
	main.AddChild(container)
	pg.AddChild(main)
	pageString := component.Generate(pg)
	err := pg.Load(pageString)
	if err != nil {
		log.Fatal(err)
	}
}
