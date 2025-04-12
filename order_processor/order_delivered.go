package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type OrderDeliveredAction struct{}

func (a *OrderDeliveredAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	order, ok := eventCtx.(*OrderCreationContext)
	if !ok {
		return "", ErrMissingOrderContext
	}
	fmt.Println("Order delivered, items:", order.items)
	return sm.NoOp, nil
}
