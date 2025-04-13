package order_processor

import (
	"context"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderTransactionFailedAction struct{}

func (a *orderTransactionFailedAction) Execute(ctx context.Context) (sm.Event, error) {
	shipment, err := sm.GetValueFromContext[Shipment](ctx, ShipmentContextKey)
	if err != nil {
		return sm.NoOp, err
	}
	fmt.Println("Transaction failed, err:", shipment.Err)
	return sm.NoOp, nil
}
