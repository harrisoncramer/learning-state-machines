package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderPlacedAction struct{}

func (a *orderPlacedAction) Execute(ctx context.Context) (sm.Event, error) {
	order, ok := ctx.Value("order").(*OrderCreationContext)
	if !ok {
		return "", fmt.Errorf("%w: order_placed got %T but expected *OrderCreationContext", ErrBadExecutionContext, order)
	}
	fmt.Println("Order placed, items:", order.Items)
	return sm.NoOp, nil
}
