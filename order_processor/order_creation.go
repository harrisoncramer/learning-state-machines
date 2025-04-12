package order_processor

import (
	"errors"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

var ErrMissingOrderContext = errors.New("missing order creation context")

// CreatingOrderAction includes the action that fires when we enter the creating order state
type CreatingOrderAction struct{}

func (c *CreatingOrderAction) Execute(event sm.EventContext) (sm.Event, error) {
	order, ok := event.(*OrderCreationContext)
	if !ok {
		return "", ErrMissingOrderContext
	}
	_ = order
	return sm.NoOp, nil
}
