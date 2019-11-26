package page

import (
	"fmt"
	"io/ioutil"

	"github.com/bgokden/go-web-lib/pkg/component"
)

type Page interface {
	component.Component
	SetTitle(title string)
	GetTitle() string
	Load(content string) error
}

type BasePage struct {
	component.BaseComponent
	Loader
	Title string
}

func (bp *BasePage) SetTitle(title string) {
	bp.Title = title
	bp.SetHeader("title", fmt.Sprintf("<title>%s</title>", title))
}

func (bp *BasePage) GetTitle() string {
	return bp.Title
}

func NewBasePage() Page {
	bp := &BasePage{}
	bp.SetId("body")
	bp.SetTitle("Default")
	bp.SetHeader("charset", `<meta charset="utf-8">`)
	bp.SetHeader("viewport", `<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">`)
	bp.SetHeader("wasm_exec", `<script src="wasm_exec.js"></script>`)
	bp.SetHeader("WebAssembly initialize",
		`<script>
    if (!go) {
        const go = new Go();
      WebAssembly.instantiateStreaming(fetch('main.wasm'),go.importObject).then( res=> {
        go.run(res.instance)
      })
    }
  </script>`)
	return bp
}

func (dp *BasePage) Render() string {
	return `
  <!doctype html>
  <html lang="en">
  	<head>
    {{ range $key, $value := .Headers }}
      {{ $value }}
    {{ end }}
  	<style>
  	</style>
  	</head>
  	<body>
    {{ range $key, $value := .GetChildren }}
      {{ Generate $value }}
    {{ end }}
  	</body>
  </html>
  `
}

type Loader struct{}

func (ld *Loader) Load(content string) error {
	err := ioutil.WriteFile("./index.html", []byte(content), 0644)
	return err
}
