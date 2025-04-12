package state_machine

import (
	"errors"
	"fmt"
	"sync"
)

// The error when the state machine cannot process the event
var ErrEventRejected = errors.New("event rejected")

// The error when the state machine reaches an invalid state due to developer error during configuration
var ErrBadConfiguration = errors.New("state machine is misconfigured")

// StateType represents an extensible state type in the state machine
type StateType string

// Default is the default state of the system
const Default StateType = ""

// EventType represents an extensible event type
type EventType string

// NoOp is an event that does not affect the state of the system
const NoOp EventType = "NoOp"

// The data passed along to the action
type Event any

// Action represents the action to be executed in a given state
type Action interface {
	Execute(event Event) EventType
}

// State binds a state with an action and a set of events it can handle
type State struct {
	Action Action
	Events Events
}

// Events represents a mapping of states and their implementations
type Events map[EventType]StateType

// States represents a mapping of states and their implementations
type States map[StateType]State

type StateMachine struct {
	previous StateType
	current  StateType
	states   States
	mutex    sync.Mutex
}

func NewStateMachine(s States) StateMachine {
	return StateMachine{
		previous: "",
		states:   s,
		mutex:    sync.Mutex{},
	}
}

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {

	// Get current state, and double check that we can actually handle events
	// with it.
	state, ok := s.states[s.current]
	if !ok || state.Events == nil {
		return Default, ErrEventRejected
	}

	// Get the next state based on the event that occurred.
	// This "next" state is based on how the current state reacts to this specific event.
	// If we cannot react appropriately, reject the event.
	next, ok := state.Events[event]
	if !ok {
		return Default, ErrEventRejected
	}

	// Return the next state
	return next, nil

}

// Send event sends an event to the state machine
func (s *StateMachine) SendEvent(event EventType) error {

	// Lock the current state so that state transitions are valid
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		// Get the next state type based on the event + current state
		nextState, err := s.getNextState(event)
		if err != nil {
			return fmt.Errorf("%w: %s not valid for %s", ErrEventRejected, event, s.current)
		}

		// Get the next nextState on the machine
		state, ok := s.states[nextState]
		if !ok || state.Action == nil {
			return ErrBadConfiguration
		}

		// Update the previous and current states
		fmt.Printf("transitioning from '%s' to '%s'\n", s.current, nextState)
		s.previous = s.current
		s.current = nextState

		// Execute the next state's action + loop until
		// getting a no-op
		nextEvent := state.Action.Execute(event)
		if nextEvent == NoOp {
			return nil
		}

		// If not a no-op, assign the new event and loop again
		event = nextEvent
	}
}

// Gets the current state, for use outside the package
func (s *StateMachine) GetCurrentState() StateType {
	return s.current
}
