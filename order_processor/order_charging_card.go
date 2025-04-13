package order_processor

import (
	"context"
	"errors"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderChargingCardAction struct{}

func (a *orderChargingCardAction) Execute(ctx context.Context) (sm.Event, error) {
	shipment, ok := ctx.Value("shipment").(*OrderShipmentContext)
	if !ok {
		return "", fmt.Errorf("%w: charging_card did not have shipment", ErrBadExecutionContext)
	}

	fmt.Println("Validating card, shipment:", shipment)
	if shipment.cardNumber == "" {
		shipment.err = errors.New("card number is invalid")
		return FailTransaction, nil
	}
	return ShipOrder, nil
}
