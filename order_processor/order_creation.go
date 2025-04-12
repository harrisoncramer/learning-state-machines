package order_processor

import (
	"errors"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

// creatingOrderAction includes the action that fires when we enter the creating order state
type creatingOrderAction struct{}

func (c *creatingOrderAction) Execute(event sm.EventContext) (sm.Event, error) {
	order, ok := event.(OrderCreationContext)
	if !ok {
		return "", fmt.Errorf("%w: order_creation got %T but expected OrderCreationContext", ErrBadExecutionContext, order)
	}

	err := doSomeBusinessLogic(order)
	if err != nil {
		order.err = err
		return FailOrder, nil
	}

	return sm.NoOp, nil
}

func doSomeBusinessLogic(order OrderCreationContext) error {
	_ = order
	return errors.New("fake error")
}
