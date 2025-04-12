package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type OrderShippedAction struct{}

func (a *OrderShippedAction) Execute(eventCtx sm.EventContext) sm.Event {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Order shipped, address:", shipment.address)
	return sm.NoOp
}
