package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderFailedAction struct{}

func (a *orderFailedAction) Execute(ctx context.Context) (sm.Event, error) {
	order, ok := ctx.Value("order").(*OrderCreationContext)
	if !ok {
		return "", fmt.Errorf("%w: order_failed got %T but expected *OrderCreationContext", ErrBadExecutionContext, order)
	}
	fmt.Printf("Order failed: %v\n", order.err)
	return sm.NoOp, nil
}
