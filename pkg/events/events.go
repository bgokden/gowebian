package events

import (
	"github.com/bgokden/gowebian/pkg/component"
)

var messageChannel = make(chan *component.Message, 100)

func Emit(m *component.Message) {
	messageChannel <- m
}

func Listen() error {
	for {
		m := <-messageChannel
		go m.To.OnMessage(m)
	}
	return nil
}
