package telefleet

import (
	"fmt"
)

// RegisterButtonHandlers registers button handlers for each bot in the fleet.
// It takes a slice of ButtonHandler maps as input. Each ButtonHandler map
// contains a button pointer and a corresponding handler function.
// The function wraps each handler function with the middleware functions
// defined in the fleet and registers the wrapped handler with the bot.
func (f *Fleet) RegisterButtonHandlers(handlers []ButtonHandler, messages chan<- string) {
	for _, bot := range f.Bots { // Iterate over each bot in the fleet
		for _, handlerMap := range handlers { // Iterate over each handler map
			for buttonPtr, handlerFunc := range handlerMap { // Iterate over each button pointer and handler function pair
				wrappedHandler := handlerFunc             // Initialize the wrapped handler with the handler function
				for _, middleware := range f.Middleware { // Iterate over each middleware function
					wrappedHandler = middleware(wrappedHandler) // Wrap the handler function with the middleware function
				}
				messages <- fmt.Sprintf("%s, ", buttonPtr.Unique)
				bot.Handle(buttonPtr, wrappedHandler) // Register the wrapped handler with the bot
			}
		}
	}
}

// RegisterHandlers registers event handlers for each bot in the fleet.
// It takes a slice of Handler maps as input. Each Handler map
// contains an event and a corresponding handler function.
// The function wraps each handler function with the middleware functions
// defined in the fleet and registers the wrapped handler with the bot.
func (f *Fleet) RegisterHandlers(handlers []Handler, messages chan<- string) {
	for _, bot := range f.Bots { // Iterate over each bot in the fleet
		for _, handlerMap := range handlers { // Iterate over each handler map
			for event, handlerFunc := range handlerMap { // Iterate over each event and handler function pair
				wrappedHandler := handlerFunc             // Initialize the wrapped handler with the handler function
				for _, middleware := range f.Middleware { // Iterate over each middleware function
					wrappedHandler = middleware(wrappedHandler) // Wrap the handler function with the middleware function
				}
				messages <- fmt.Sprintf("%s, ", event)
				bot.Handle(event, wrappedHandler) // Register the wrapped handler with the bot
			}
		}
	}
}
