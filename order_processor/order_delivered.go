package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderDeliveredAction struct{}

func (a *orderDeliveredAction) Execute(ctx context.Context) (sm.Event, error) {
	order, ok := ctx.Value("order").(*OrderCreationContext)
	if !ok {
		return "", fmt.Errorf("%w: order_delivered got %T but expected OrderCreationContext", ErrBadExecutionContext, order)
	}
	fmt.Println("Order delivered, items:", order.Items)
	return sm.NoOp, nil
}
