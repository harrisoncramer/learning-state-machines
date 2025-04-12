package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderTransactionFailedAction struct{}

func (a *orderTransactionFailedAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	shipment, ok := eventCtx.(*OrderShipmentContext)
	if !ok {
		return "", fmt.Errorf("%w: order_transaction_failed got %T but expected OrderShipmentContext", ErrBadExecutionContext, shipment)
	}
	fmt.Println("Transaction failed, err:", shipment.err)
	return sm.NoOp, nil
}
