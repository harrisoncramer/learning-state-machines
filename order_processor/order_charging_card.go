package order_processor

import (
	"errors"
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type ChargingCardAction struct{}

func (a *ChargingCardAction) Execute(eventCtx sm.EventContext) sm.Event {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Validating card, shipment:", shipment)
	if shipment.cardNumber == "" {
		shipment.err = errors.New("card number is invalid")
		return FailTransaction
	}
	return ShipOrder
}
