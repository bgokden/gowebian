package page

import (
	"fmt"

	"github.com/bgokden/go-web-lib/pkg/component"
)

type Page interface {
	component.Component
	SetTitle(title string)
	GetTitle() string
	SetHeader(key, value string)
	GetHeader(key string) string
	GetHeaders() map[string]string
}

type BasePage struct {
	component.BaseComponent
	Headers map[string]string
}

func (dp *BasePage) SetTitle(title string) {
	dp.Headers["title"] = fmt.Sprintf("<title>%s</title>", title)
}

func (dp *BasePage) GetTitle() string {
	return dp.Headers["title"]
}

func (dp *BasePage) SetHeader(key, value string) {
	dp.Headers[key] = value
}

func (dp *BasePage) GetHeader(key string) string {
	return dp.Headers[key]
}

func (dp *BasePage) GetHeaders() map[string]string {
	return dp.Headers
}

func NewBasePage() Page {
	dp := &BasePage{
		Headers: make(map[string]string),
	}
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

/*
func (dp *BasePage) String() string {
	templateText := dp.Render()
	tmpl, err := template.New("Render").Parse(templateText)
	if err != nil {
		return err.Error()
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, dp)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
*/
