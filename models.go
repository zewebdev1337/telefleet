package telefleet

import "gopkg.in/telebot.v3"

// Fleet represents a group of telegram bots.
// It contains a slice of telebot.Bot pointers and a slice of telebot.MiddlewareFunc.
type Fleet struct {
	Bots       []*telebot.Bot           // A slice of pointers to telebot.Bot instances.
	Middleware []telebot.MiddlewareFunc // A slice of middleware functions for the bots.
}

// FleetConfig represents the configuration for a group of telegram bots.
// It contains the name of the fleet, a slice of tokens, a slice of handlers, a slice of button handlers,
// a slice of middleware functions, and a poller.
type FleetConfig struct {
	Name           string                   // The name of the fleet.
	Tokens         []string                 // A slice of tokens for the bots in the fleet.
	Handlers       []Handler                // A slice of handlers for the bots in the fleet.
	ButtonHandlers []ButtonHandler          // A slice of button handlers for the bots in the fleet.
	Middleware     []telebot.MiddlewareFunc // A slice of middleware functions for the bots in the fleet.
	Poller         telebot.Poller           // The poller for the bots in the fleet.
}

// Handler associates a string with a function that handles telebot.Context.
// It is used to handle commands or messages.
type Handler map[string]func(telebot.Context) error

// ButtonHandler associates a *telebot.InlineButton with a function that handles telebot.Context.
// It is used to handle button presses.
type ButtonHandler map[*telebot.InlineButton]func(telebot.Context) error
