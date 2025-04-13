package light_switch

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

// The light switch can be in either an lightOn or an lightOff state
const (
	lightOff sm.State = "Off"
	lightOn  sm.State = "On"
)

// The two types of events are a "switch off" or a "switch on"
const (
	SwitchOff sm.Event = "SwitchOff"
	SwitchOn  sm.Event = "SwitchOn"
)

// onAction represents the action executed on entering the On state.
type onAction struct{}

func (a *onAction) Execute(ctx context.Context) (sm.Event, error) {
	fmt.Println("Light turned on: 💡")
	if ctx != nil {
		fmt.Printf("Context provided: %+v\n", ctx)
	}
	return sm.NoOp, nil
}

func NewLightSwitch() (*sm.StateMachine, error) {
	return sm.NewStateMachine(lightOff, sm.StateMap{
		lightOff: {
			EventMap: sm.EventMap{
				SwitchOn: lightOn,
			},
		},
		lightOn: {
			Action: &onAction{},
			EventMap: sm.EventMap{
				SwitchOff: lightOff,
			},
		},
	})
}
