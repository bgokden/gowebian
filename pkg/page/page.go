package page

import (
	"fmt"

	"github.com/bgokden/go-web-lib/pkg/component"
)

type Page interface {
	component.Component
	SetTitle(title string)
	GetTitle() string
}

type BasePage struct {
	component.BaseComponent
	Title string
}

func (dp *BasePage) SetTitle(title string) {
	dp.Title = title
	dp.SetHeader("title", fmt.Sprintf("<title>%s</title>", title))
}

func (dp *BasePage) GetTitle() string {
	return dp.Title
}

func NewBasePage() Page {
	dp := &BasePage{}
	dp.SetId("body")
	dp.SetTitle("Default")
	dp.SetHeader("charset", `<meta charset="utf-8">`)
	dp.SetHeader("viewport", `<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">`)
	dp.SetHeader("wasm_exec", `<script src="wasm_exec.js"></script>`)
	dp.SetHeader("WebAssembly initialize",
		`<script>
    if (!go) {
        const go = new Go();
      WebAssembly.instantiateStreaming(fetch('main.wasm'),go.importObject).then( res=> {
        go.run(res.instance)
      })
    }
  </script>`)
	return dp
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
