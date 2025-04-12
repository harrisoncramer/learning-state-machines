package light_switch

import (
	"fmt"

	sm "github.com/harrisoncramer/learning-state-machines/state_machine"
)

const (
	// The light switch can be in either an on or an off state
	Off sm.State = "Off"
	On  sm.State = "On"

	// The two types of events are a "switch off" or a "switch on"
	SwitchOff sm.EventType = "SwitchOff"
	SwitchOn  sm.EventType = "SwitchOn"
)

// OnAction represents the action executed on entering the On state.
type OnAction struct{}

func (a *OnAction) Execute(event sm.EventType) sm.EventType {
	fmt.Printf("%s: Light turned on: 💡\n", event)
	return sm.NoOp
}

func NewLightSwitch() *sm.StateMachine {
	return sm.NewStateMachine(Off, sm.StateMap{
		Off: {
			EventMap: sm.EventMap{
				SwitchOn: On,
			},
		},
		On: {
			Action: &OnAction{},
			EventMap: sm.EventMap{
				SwitchOff: Off,
			},
		},
	})
}
