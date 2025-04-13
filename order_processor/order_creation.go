package order_processor

import (
	"context"
	"errors"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

// creatingOrderAction includes the action that fires when we enter the creating order state
type creatingOrderAction struct{}

func (c *creatingOrderAction) Execute(ctx context.Context) (sm.Event, error) {
	order, ok := ctx.Value("order").(*Order)
	if !ok {
		return "", fmt.Errorf("%w: order_creation did not have order", ErrBadExecutionContext)
	}

	err := doSomeBusinessLogic(*order)
	if err != nil {
		order.Err = err
		return FailOrder, nil
	}

	return sm.NoOp, nil
}

func doSomeBusinessLogic(order Order) error {
	_ = order
	return errors.New("fake error")
}
