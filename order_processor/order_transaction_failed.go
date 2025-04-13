package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderTransactionFailedAction struct{}

func (a *orderTransactionFailedAction) Execute(ctx context.Context) (sm.Event, error) {
	shipment, ok := ctx.Value("shipment").(*Shipment)
	if !ok {
		return "", fmt.Errorf("%w: order_transaction_failed got %T but expected *OrderShipmentContext", ErrBadExecutionContext, shipment)
	}
	fmt.Println("Transaction failed, err:", shipment.Err)
	return sm.NoOp, nil
}
