package telefleet

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gopkg.in/telebot.v3"
)

var (
	globalWaitGroup  sync.WaitGroup
	poller           = &telebot.LongPoller{Timeout: 10 * time.Second}
	singleBot, fleet *Fleet

	fleetTokens = []string{
		"TOKEN",
		"TOKEN",
	}
	fleetHandlers = []Handler{{
		telebot.OnText: handleText,
	}}
	fleetButtonHandlers = []ButtonHandler{{
		// ... buttonPointer, HandlerFunction here
	}}
	fleetMiddleware = []telebot.MiddlewareFunc{
		// {
		// // ... middleware factory function here
		// }
	}
	fleetConfig = FleetConfig{
		Name:           "multi bot fleet example",
		Tokens:         fleetTokens,
		Handlers:       fleetHandlers,
		ButtonHandlers: fleetButtonHandlers,
		Poller:         poller,
		Middleware:     fleetMiddleware,
	}
	botToken = []string{
		"TOKEN",
	}
	botHandlers = []Handler{{
		telebot.OnPhoto: handlePhoto,
	}}
	botButtonHandlers = []ButtonHandler{{
		// ... buttonPointer, HandlerFunction here
	}}
	botMiddleware = []telebot.MiddlewareFunc{
		// {
		// ... middleware factory function here
		// }
	}
	botFleetConfig = FleetConfig{
		Name:           "single bot, no poller example",
		Tokens:         botToken,
		Handlers:       botHandlers,
		ButtonHandlers: botButtonHandlers,
		Poller:         nil,
		Middleware:     botMiddleware,
	}
)

func handleText(c telebot.Context) error {
	c.Send(
		fmt.Sprintf("Hi %s! My name is [%s](https://t.me/%s). Your message has been received.", c.Sender().FirstName, c.Bot().Me.FirstName, c.Bot().Me.Username),
		&telebot.SendOptions{
			ParseMode: telebot.ModeMarkdown,
		},
	)
	return nil
}

func handlePhoto(c telebot.Context) error {
	c.Send(
		fmt.Sprintf("Hi %s! My name is [%s](https://t.me/%s). Your photo has been received.", c.Sender().FirstName, c.Bot().Me.FirstName, c.Bot().Me.Username),
		&telebot.SendOptions{
			ParseMode: telebot.ModeMarkdown,
		},
	)
	return nil
}

func main() {
	var err error
	messages := make(chan string, 100)

	go func() {
		for msg := range messages {
			fmt.Println(msg)
		}
	}()

	fleet, err = NewFleet(fleetConfig, messages)
	if err != nil {
		log.Fatal(err)
	}

	singleBot, err = NewFleet(botFleetConfig, messages)
	if err != nil {
		log.Fatal(err)
	}

	close(messages)

	StartFleets()
}
