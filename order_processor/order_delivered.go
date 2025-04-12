package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderDeliveredAction struct{}

func (a *orderDeliveredAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	order, ok := eventCtx.(OrderCreationContext)
	if !ok {
		return "", fmt.Errorf("%w: order_delivered got %T but expected OrderCreationContext", ErrBadExecutionContext, order)
	}
	fmt.Println("Order delivered, items:", order.Items)
	return sm.NoOp, nil
}
