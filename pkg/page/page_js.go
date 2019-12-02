// +build js,wasm
package page

import (
	"github.com/bgokden/gowebian/pkg/events"
)

func (bp *BasePage) Load(content string) error {
	// doc := js.Global().Get("document")
	// doc.Call("write", content)
	bp.Register(bp)
	return events.Listen()
}
