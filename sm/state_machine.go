package sm

import (
	"fmt"
	"log"
	"sync"

	logger "github.com/harrisoncramer/learning-state-machines/sm/logger"
	"go.uber.org/zap"
)

// StateMachine holds all states, actions, and events relevant to a system. This includes:
// 1. The current state of the system
// 2. A stateMap that dictates for each possible state, an optional action that should occur, and an event map which indicates which state to transition to when
// the state machine recieves an event of that type in that state
// 3. A mutex which ensures that the state machine can only be processing one event at a given time.
type StateMachine struct {
	currentState State
	stateMap     StateMap
	mutex        sync.Mutex
	logger       *zap.Logger
}

type Option func(*StateMachine)

func NewStateMachine(initialState State, stateMap StateMap, opts ...Option) *StateMachine {
	logger, err := logger.GetZapLogger("info", "2006-01-02 15:04:05.000 MST")
	if err != nil {
		log.Fatalf("failed to configure logger: %v", err)
	}

	s := &StateMachine{
		currentState: initialState,
		stateMap:     stateMap,
		mutex:        sync.Mutex{},
		logger:       logger,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// NoOp is an event that does not affect the state of the system
const NoOp Event = "NoOp"

// State represents an extensible state type in the state machine
type State string

func (s State) String() string {
	return string(s)
}

// Event represents an extensible event type
type Event string

func (s Event) String() string {
	return string(s)
}

// EventContext is the data that's passed to each action
type EventContext any

// Action represents the action to be executed in a given state
type Action interface {
	Execute(eventData EventContext) Event
}

// EventMap represents a mapping of states and their implementations
type EventMap map[Event]State

// StateMap represents a mapping of states and their implementations
type StateMap map[State]StateSet
type StateSet struct {
	Action   Action
	EventMap EventMap
}

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.
func (s *StateMachine) getNextState(event Event) (State, error) {

	// Get current state, and double check that we can actually handle events
	// with it.
	state, ok := s.stateMap[s.currentState]
	if !ok || state.EventMap == nil {
		return "", ErrEventRejected
	}

	// Get the next state based on the event that occurred.
	// This "next" state is based on how the current state reacts to this specific event.
	// If we cannot react appropriately, reject the event.
	next, ok := state.EventMap[event]
	if !ok {
		return "", ErrEventRejected
	}

	// Return the next state
	return next, nil

}

// Send event sends an event to the state machine
func (s *StateMachine) SendEvent(event Event, eventContext EventContext) error {

	// Lock the current state so that state transitions are valid
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		// Get the next state type based on the event + current state
		nextState, err := s.getNextState(event)
		if err != nil {
			return fmt.Errorf("%w: %s not valid for %s", ErrEventRejected, event, s.currentState)
		}

		// Get the next nextState on the machine
		state, ok := s.stateMap[nextState]
		if !ok {
			return fmt.Errorf("%w: could not find '%s' state in state map", ErrBadConfiguration, nextState)
		}

		// Update the current state
		s.logger.Debug(
			"Updated current state",
			zap.Stringer("current_state", s.currentState),
			zap.Stringer("next_state", nextState),
		)

		s.currentState = nextState

		// If there is no action required when entering this state
		// then just return early
		if state.Action == nil {
			s.logger.Debug("No action found, exiting")
			return nil
		}

		// Otherwise, execute the next state's action and pass along the event context
		// This event context can be used for side effects or to pass along other arbitrary data
		// besides just the event name that may be needed by the action itself
		nextEvent := state.Action.Execute(eventContext)

		// If it returns a no-op, then exit
		if nextEvent == NoOp {
			s.logger.Debug("Action returned nil, exiting")
			return nil
		}

		s.logger.Debug(
			"Action returned next event",
			zap.Stringer("next_event", nextEvent),
		)

		// If it's not a no-op, assign the new event and loop again
		event = nextEvent
	}
}

// Gets the current state, for use outside the package
func (s *StateMachine) GetCurrentState() State {
	return s.currentState
}
