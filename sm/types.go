package sm

// State represents an extensible state type in the state machine
// The state machine will always be in one of these state types at any
// given time
type State string

func (s State) String() string {
	return string(s)
}

// Event represents an extensible event type. These events are passed through the
// state machine in order to trigger actions and state changes
type Event string

func (s Event) String() string {
	return string(s)
}

// EventData is data that's passed to each action when an event occurs. It's not required
// but can be used to pass data through the system, often via a pointer
type EventData any

// NoOp is a special event that will cause the state machine to stop executing
const NoOp Event = "NoOp"

// EventMap represents a mapping between events and the states they will trigger.
// For instance, the "failOrder" event may correspond with the OrderFailed state.
type EventMap map[Event]State

// StateMap represents a mapping of states and their implementations. Every state
// in the system will have an event map, which indicates how that state will respond
// to specific events, and an optional action, which fires when the state is entered
// in order to return new events and continue execution
type StateMap map[State]struct {
	// Action represents the action to be executed for a given state. It's fired when that state
	// is entered, and will return a new event, which will cause a new state to be entered, and so forth.
	// If a state does not have an action, then the state machine will treat that state as the end of the chain (like a no-op).
	Action interface {
		Execute(eventData EventData) (Event, error)
	}
	EventMap EventMap
}
