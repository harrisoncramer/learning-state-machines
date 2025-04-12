package order_processor

import (
	"errors"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderChargingCardAction struct{}

func (a *orderChargingCardAction) Execute(eventCtx sm.EventData) (sm.Event, error) {
	shipment, ok := eventCtx.(*OrderShipmentContext)
	if !ok {
		return "", fmt.Errorf("%w: charging_card got %T but expected *OrderShipmentContext", ErrBadExecutionContext, shipment)
	}

	fmt.Println("Validating card, shipment:", shipment)
	if shipment.cardNumber == "" {
		shipment.err = errors.New("card number is invalid")
		return FailTransaction, nil
	}
	return ShipOrder, nil
}
