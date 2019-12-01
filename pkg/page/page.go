package page

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bgokden/go-web-lib/pkg/component"
	"golang.org/x/net/html"
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
	InitDefaults(bp)
	return bp
}

func NewPage() Page {
	return NewBasePage()
}

func InitDefaults(bp Page) {
	bp.SetKey("body")
	bp.SetTitle("Default")
	bp.SetHeader("charset", `<meta charset="utf-8">`)
	bp.SetHeader("viewport", `<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">`)
	bp.SetHeader("x-ua-compatible", `<meta http-equiv="x-ua-compatible" content="ie=edge">`)
	bp.SetHeader("WebAssembly initialize",
		`<script src="wasm_exec.js"></script>
    <script>
      if (typeof go == 'undefined') {
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch('main.wasm'),go.importObject).then( res=> {
          go.run(res.instance)
        })
      }
    </script>`)
}

func (dp *BasePage) Render() string {
	return `
  <!doctype html>
  <html lang="en">
  	<head>
    {{ range $key, $value := .GetHeaders }}
      {{ $value }}
    {{ end }}
  	<style>
  	</style>
  	</head>
  	<body {{ range $key, $value := .GetAttributes }} {{ printf "%s=\"%s\"" $key $value }} {{ end }}>
    {{ range $key, $value := .GetChildren }}
      {{ Generate $value }}
    {{ end }}
  	</body>
  </html>
  `
}

type Loader struct{}

func (ld *Loader) Load(content string) error {
	_, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./public/index.html", []byte(content), 0644)
	return err
}
