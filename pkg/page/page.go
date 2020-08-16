package page

import (
	"io/ioutil"
	"strings"

	"github.com/bgokden/gowebian/pkg/component"
	"github.com/bgokden/gowebian/pkg/components"
	"golang.org/x/net/html"
)

type Page interface {
	component.Component
	SetTitle(title string)
	GetTitle() string
	SetHeader(key string, value component.Component)
	GetHeader(key string) component.Component
	Load(content string) error
}

type BasePage struct {
	component.BaseComponent
	Loader
}

func (bp *BasePage) SetTitle(title string) {
	bp.SetHeader("title", components.NewTitle(title))
}

func (bp *BasePage) GetTitle() string {
	return bp.GetHeader("title").GetValue()
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
	bp.SetKey("html")
	bp.SetAttribute("lang", "en")
	bp.SetChild("head", components.NewHead())
	bp.SetChild("body", components.NewBody())
	bp.SetTitle("Default")
	bp.SetHeader("charset", components.NewMeta(map[string]string{
		"charset": "utf-8",
	}))
	bp.SetHeader("viewport", components.NewMeta(map[string]string{
		"name":    "viewport",
		"content": "width=device-width, initial-scale=1, shrink-to-fit=no",
	}))
	bp.SetHeader("x-ua-compatible", components.NewMeta(map[string]string{
		"http-equiv": "x-ua-compatible",
		"content":    "ie=edge",
	}))
	bp.SetHeader("WebAssemblyinitialize_0", components.NewScriptFromSource("wasm_exec.js"))
	bp.SetHeader("WebAssemblyinitialize_1", components.NewScriptFromCode(
		`if (typeof go == 'undefined') {
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch('main.wasm'),go.importObject).then( res=> {
          go.run(res.instance)
        })
      }`))
}

func (bp *BasePage) SetHeader(key string, value component.Component) {
	bp.GetChild("head").SetChild(key, value)
}

func (bp *BasePage) GetHeader(key string) component.Component {
	return bp.GetChild("head").GetChild(key)
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
