package light_switch

import (
	"fmt"

	sm "github.com/harrisoncramer/learning-state-machines/state_machine"
)

const (
	// The light switch can be in either an on or an off state
	Off sm.StateType = "Off"
	On  sm.StateType = "On"

	// The two types of events are a "switch off" or a "switch on"
	SwitchOff sm.EventType = "SwitchOff"
	SwitchOn  sm.EventType = "SwitchOn"
)

// OffAction represents the action executed on entering the Off state.
type OffAction struct{}

func (a *OffAction) Execute(event sm.Event) sm.EventType {
	fmt.Println("The light has been switched off")
	return sm.NoOp
}

// OnAction represents the action executed on entering the On state.
type OnAction struct{}

func (a *OnAction) Execute(event sm.Event) sm.EventType {
	fmt.Println("The light has been switched on")
	return sm.NoOp
}

func NewLightSwitch() sm.StateMachine {
	return sm.NewStateMachine(sm.States{
		sm.Default: sm.State{
			Events: sm.Events{
				SwitchOff: Off,
			},
		},
		Off: sm.State{
			Action: &OffAction{},
			Events: sm.Events{
				SwitchOn: On,
			},
		},
		On: sm.State{
			Action: &OnAction{},
			Events: sm.Events{
				SwitchOff: Off,
			},
		},
	})
}
