package light_switch

import (
	"context"
	"errors"
	"testing"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type EventAndState struct {
	event *sm.Event
	state sm.State
	error error
}

func TestLightSwitch(t *testing.T) {

	switchOn := SwitchOn // needed for pointers...
	switchOff := SwitchOff

	tests := []struct {
		name   string
		events []EventAndState
	}{
		{
			"Should start off",
			[]EventAndState{
				{
					nil,
					lightOff,
					nil,
				},
			},
		},
		{
			"Should turn on",
			[]EventAndState{
				{
					&switchOn,
					lightOn,
					nil,
				},
			},
		},
		{
			"Should turn on and off",
			[]EventAndState{
				{
					&switchOn,
					lightOn,
					nil,
				},
				{
					&switchOff,
					lightOff,
					nil,
				},
			},
		},
		{
			"Should reject turning off an already off switch",
			[]EventAndState{
				{
					&switchOff,
					lightOff,
					sm.ErrEventRejected,
				},
				{
					&switchOn, // Still processes next event
					lightOn,
					nil,
				},
			},
		},
		{
			"Should reject turning on an already on switch",
			[]EventAndState{
				{
					&switchOn,
					lightOn,
					nil,
				},
				{
					&switchOn,
					lightOn,
					sm.ErrEventRejected,
				},
				{
					&switchOff, // Still processes subsequent events
					lightOff,
					nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ls, err := NewLightSwitch()
			if err != nil {
				t.Fatalf("failed to set up light switch: %v", err)
			}
			for _, eventAndState := range tt.events {
				if eventAndState.event != nil {
					err := ls.SendEvent(context.Background(), *eventAndState.event)
					if err != nil {
						if eventAndState.error == nil {
							t.Fatalf("Got unexpected error: %v", err)
						}
						if !errors.Is(err, eventAndState.error) {
							t.Fatalf("Wanted error %v but got error %v", eventAndState.error, err)
						}
					}
					if ls.GetCurrentState() != eventAndState.state {
						t.Fatalf("Expected %s but got %s", eventAndState.state, ls.GetCurrentState())
					}
				}
			}
		})
	}
}
