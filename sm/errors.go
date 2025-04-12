package sm

import "errors"

// The error when the state machine cannot process the event
var ErrEventRejected = errors.New("event rejected")

// The error when the state machine reaches an invalid state due to developer error during configuration
var ErrBadConfiguration = errors.New("state machine is misconfigured")
