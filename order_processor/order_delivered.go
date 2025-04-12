package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type OrderDeliveredAction struct{}

func (a *OrderDeliveredAction) Execute(eventCtx sm.EventContext) sm.Event {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Order delivered, items:", order.items)
	return sm.NoOp
}
