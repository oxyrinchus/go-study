//go:build wireinject
// +build wireinject

package main

import (
	"github/Slowhigh/go-study/go_labs/wire/event"
	"github/Slowhigh/go-study/go_labs/wire/greeter"
	"github/Slowhigh/go-study/go_labs/wire/message"

	"github.com/google/wire"
)

func InitializeEvent() (event.Event, error) {
	wire.Build(event.NewEvent, greeter.NewGreeter, message.NewMessage)
	return event.Event{}, nil
}
