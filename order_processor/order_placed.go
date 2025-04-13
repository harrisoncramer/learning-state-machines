package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderPlacedAction struct{}

func (a *orderPlacedAction) Execute(ctx context.Context) (sm.Event, error) {
	order, err := sm.GetValueFromContext[Order](ctx, OrderContextKey)
	if err != nil {
		return sm.NoOp, err
	}
	fmt.Println("Order placed, items:", order.Items)
	return sm.NoOp, nil
}
