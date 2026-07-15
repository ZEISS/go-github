package apps

import (
	"log"

	"github.com/google/go-github/v89/github"
	hooks "github.com/zeiss/fiber-hooks/v3"
)

var _ hooks.Dispatcher = (*dispatcherImpl)(nil)

type dispatcherImpl struct {
	config *Config
}

// NewDispatcher creates a new dispatcher instance.
func NewDispatcher(config *Config) *dispatcherImpl {
	return &dispatcherImpl{config: config}
}

// Dispatch dispatches a hook event to the configured handler.
func (d *dispatcherImpl) Dispatch(event hooks.Event) error {
	parsedEvent, err := github.ParseWebHook(event.EventType, event.Payload)
	if err != nil {
		return err
	}

	switch parsedEvent.(type) {
	case *github.IssueCommentEvent:
		log.Print(parsedEvent)
	case *github.IssueEvent:
		log.Print(parsedEvent)
	}

	return nil
}
