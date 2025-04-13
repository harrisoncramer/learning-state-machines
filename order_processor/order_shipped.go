package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderShippedAction struct{}

func (a *orderShippedAction) Execute(ctx context.Context) (sm.Event, error) {
	shipment, ok := ctx.Value("shipment").(*Shipment)
	if !ok {
		return "", fmt.Errorf("%w: order_shipped got %T but expected *OrderShipmentContext", ErrBadExecutionContext, shipment)
	}
	fmt.Println("Order shipped, address:", shipment.Address)
	return sm.NoOp, nil
}
