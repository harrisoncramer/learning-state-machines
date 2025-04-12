package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderShippedAction struct{}

func (a *orderShippedAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	shipment, ok := eventCtx.(*OrderShipmentContext)
	if !ok {
		return "", fmt.Errorf("%w: order_shipped got %T but expected OrderShipmentContext", ErrBadExecutionContext, shipment)
	}
	fmt.Println("Order shipped, address:", shipment.address)
	return sm.NoOp, nil
}
