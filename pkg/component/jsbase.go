package component

import "log"

type JsBase struct{}

func (jb *JsBase) SetProperty(key string, value interface{}) {
	log.Printf("JsBase SetProperty: %v\n", value)
}
