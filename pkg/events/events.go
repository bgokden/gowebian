package events

import (
	"github.com/bgokden/go-web-lib/pkg/component"
)

var messageChannel = make(chan *component.Message, 100)

func Emit(m *component.Message) {
	messageChannel <- m
}

func Listen() {
	for {
		m := <-messageChannel
		go m.To.OnMessage(m)
	}
}
