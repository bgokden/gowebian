//go:generate go run main.go
package main

import (
	"log"

	"github.com/bgokden/gowebian/mdb"
	"github.com/bgokden/gowebian/pkg/component"
)

func main() {
	pg := mdb.NewPage()
	pg.GetChild("body").SetValue(`
			<div style="height: 100vh">
				<div class="flex-center flex-column">
					<h1 class="animated fadeIn mb-2">Material Design for Bootstrap</h1>

					<h5 class="animated fadeIn mb-1">Thank you for using our product. We're glad you're with us.</h5>

					<p class="animated fadeIn text-muted">MDB Team</p>
				</div>
			</div>
			`)
	pageString := component.Generate(pg)
	err := pg.Load(pageString)
	if err != nil {
		log.Fatal(err)
	}
}
