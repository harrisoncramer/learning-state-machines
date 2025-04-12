package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type TransactionFailedAction struct{}

func (a *TransactionFailedAction) Execute(eventCtx sm.EventContext) sm.Event {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Transaction failed, err:", shipment.err)
	return sm.NoOp
}
