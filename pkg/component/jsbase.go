package component

import (
	"log"
)

type JsBase struct{}

func (jb *JsBase) SetProperty(key string, value interface{}) {
	log.Printf("JsBase SetProperty: %v\n", value)
}

func (bc *JsBase) SetPropertyWithId(id string, key string, value interface{}) {
	log.Printf("JsBase SetPropertyWithId: %v\n", value)
}

func (bc *JsBase) Register(c Component) {
}
