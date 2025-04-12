package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type OrderPlacedAction struct{}

func (a *OrderPlacedAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	order, ok := eventCtx.(*OrderCreationContext)
	if !ok {
		return "", ErrMissingOrderContext
	}
	fmt.Println("Order placed, items:", order.items)
	return sm.NoOp, nil
}
