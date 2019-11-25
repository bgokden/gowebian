package events

import (
	"fmt"

	"github.com/bgokden/go-web-lib/pkg/component"
)

var messageChannel = make(chan *component.Message, 100)

func Emit(m *component.Message) {
	messageChannel <- m
}

func Listen() {
	for {
		m := <-messageChannel
		fmt.Println(m)
		m.To.OnMessage(m)
	}
}
