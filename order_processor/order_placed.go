package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type OrderPlacedAction struct{}

func (a *OrderPlacedAction) Execute(eventCtx sm.EventContext) sm.Event {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Order placed, items:", order.items)
	return sm.NoOp
}
