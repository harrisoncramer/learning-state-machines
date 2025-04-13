package order_processor

import (
	"context"
	"errors"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

// creatingOrderAction includes the action that fires when we enter the creating order state
type creatingOrderAction struct{}

func (c *creatingOrderAction) Execute(ctx context.Context) (sm.Event, error) {
	order, err := sm.GetValueFromContext[Order](ctx, OrderContextKey)
	if err != nil {
		return sm.NoOp, err
	}

	err = doSomeBusinessLogic(order)
	if err != nil {
		order.Err = err
		return FailOrder, nil
	}

	return sm.NoOp, nil
}

func doSomeBusinessLogic(order *Order) error {
	_ = order
	return errors.New("fake error")
}
