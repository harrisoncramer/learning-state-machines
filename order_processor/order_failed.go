package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type OrderFailedAction struct{}

func (a *OrderFailedAction) Execute(eventCtx sm.EventContext) sm.Event {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Order failed, err:", order.err)
	return sm.NoOp
}
