package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderPlacedAction struct{}

func (a *orderPlacedAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	order, ok := eventCtx.(*OrderCreationContext)
	if !ok {
		return "", ErrMissingOrderContext
	}
	fmt.Println("Order placed, items:", order.Items)
	return sm.NoOp, nil
}
