package telefleet

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/telebot.v3"
)

var wg sync.WaitGroup

// NewFleet creates a new Fleet instance with the provided configuration
func NewFleet(config FleetConfig, messages chan<- string) (*Fleet, error) {
	fleet := &Fleet{}       // Initialize an empty Fleet struct
	var bots []*telebot.Bot // Initialize an empty slice to hold the bots
	messages <- fmt.Sprintf("Processing %s config...", config.Name)

	// Loop through each token in the config
	for _, token := range config.Tokens {
		// Create a new bot with the token and poller settings
		bot, err := telebot.NewBot(telebot.Settings{
			Token:  token,
			Poller: config.Poller,
		})
		if err != nil {
			return nil, err // If there's an error creating the bot, return the error
		}
		bots = append(bots, bot) // Add the bot to the slice of bots
		// Apply all middleware to the bot
		for _, middleware := range config.Middleware {
			bot.Use(middleware)
		}
	}

	// Create a new Fleet instance with the slice of bots and middleware
	fleet = &Fleet{Bots: bots, Middleware: config.Middleware}

	// Register handlers and button handlers for the fleet
	messages <- fmt.Sprintf("%s Starting registration of event handlers for '%s'... Registered handlers: { ", time.Now().Format("2006/01/02 15:04:05"), config.Name)
	fleet.RegisterHandlers(config.Handlers, messages)
	messages <- "}"
	messages <- fmt.Sprintf("%s Starting registration of button handlers for '%s'... Registered handlers: { ", time.Now().Format("2006/01/02 15:04:05"), config.Name)
	fleet.RegisterButtonHandlers(config.ButtonHandlers, messages)
	messages <- "}"
	// Initialize the fleet by starting all the bots in goroutines
	fleet.InitializeFleet(messages)

	return fleet, nil // Return the fleet and nil error
}

// InitializeFleet starts all the bots in the Fleet in separate goroutines
// It uses a WaitGroup to ensure that the main function waits for all the bots to finish their execution
// It also sends messages back to the caller using a channel
func (f *Fleet) InitializeFleet(messages chan<- string) {
	// Loop through each bot in the Fleet
	for _, bot := range f.Bots {
		// Increment the WaitGroup counter
		wg.Add(1)
		// Start a new goroutine for the bot
		go func(b *telebot.Bot) {
			// Decrement the WaitGroup counter when the goroutine finishes
			defer wg.Done()
			// Send a message back to the caller
			messages <- fmt.Sprintf("%s (https://t.me/%s) started...", b.Me.FirstName, b.Me.Username)
			// Start the bot
			b.Start()
		}(bot)
	}
}

// StartFleets function waits for all the bots in all fleets to finish their execution
func StartFleets() {
	// Wait for the WaitGroup counter to reach zero
	wg.Wait()
}
