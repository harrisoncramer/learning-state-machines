package order_processor

import (
	"context"
	"errors"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderChargingCardAction struct{}

func (a *orderChargingCardAction) Execute(ctx context.Context) (sm.Event, error) {
	shipment, err := sm.GetValueFromContext[Shipment](ctx, ShipmentContextKey)
	if err != nil {
		return sm.NoOp, err
	}

	fmt.Println("Validating card, shipment:", shipment)
	if shipment.CardNumber == "" {
		shipment.Err = errors.New("card number is invalid")
		return FailTransaction, nil
	}
	return ShipOrder, nil
}
