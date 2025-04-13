package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderFailedAction struct{}

func (a *orderFailedAction) Execute(ctx context.Context) (sm.Event, error) {
	order, err := sm.GetValueFromContext[Order](ctx, OrderContextKey)
	if err != nil {
		return sm.NoOp, err
	}

	fmt.Printf("Order %s failed: %v\n", order.ID, order.Err)
	return sm.NoOp, nil
}
