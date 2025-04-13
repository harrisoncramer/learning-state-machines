package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderFailedAction struct{}

func (a *orderFailedAction) Execute(ctx context.Context) (sm.Event, error) {
	order, ok := ctx.Value("order").(*Order)
	if !ok {
		return "", fmt.Errorf("%w: order_failed did not have order", ErrBadExecutionContext, order)
	}
	fmt.Printf("Order %s failed: %v\n", order.ID, order.Err)
	return sm.NoOp, nil
}
