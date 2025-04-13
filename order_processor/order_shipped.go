package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderShippedAction struct{}

func (a *orderShippedAction) Execute(ctx context.Context) (sm.Event, error) {
	shipment, err := sm.GetValueFromContext[Shipment](ctx, ShipmentContextKey)
	if err != nil {
		return sm.NoOp, err
	}
	fmt.Println("Order shipped, address:", shipment.Address)
	return sm.NoOp, nil
}
