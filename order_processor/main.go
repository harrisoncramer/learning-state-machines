package order_processor

import (
	"errors"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

// Possible states of the state machine for a given order, e.g. card is charging, order is shipped, etc
const (
	// The default state of the machine, no order has been placed yet
	orderNotPlaced sm.State = "OrderNotPlaced"
	// The order is being created
	orderCreating sm.State = "OrderCreating"
	// The order creating failed and is now a failed order
	orderFailed sm.State = "OrderFailed"
	// The order is placed successfully
	orderPlaced sm.State = "OrderPlaced"
	// The credit card associated with the order is being charged
	orderChargingCard sm.State = "ChargingCard"
	// The transaction (financial processing) failed
	orderTransactionFailed sm.State = "OrderTransactionFailed"
	// The order has been shipped
	orderShipped sm.State = "OrderShipped"
	// The order has been delivered successfully
	orderDelivered sm.State = "OrderDelivered"
)

// Possible events that cause changes in the state machine
const (
	CreateOrder     sm.Event = "CreateOrder"
	FailOrder       sm.Event = "FailOrder"
	PlaceOrder      sm.Event = "PlaceOrder"
	ChargeCard      sm.Event = "ChargeCard"
	FailTransaction sm.Event = "FailTransaction"
	ShipOrder       sm.Event = "ShipOrder"
	DeliverOrder    sm.Event = "DeliverOrder"
)

var ErrMissingOrderContext = errors.New("missing order creation context")

func NewOrderProcessor(opts ...sm.Option) (*sm.StateMachine, error) {
	return sm.NewStateMachine(orderNotPlaced, sm.StateMap{
		orderNotPlaced: {
			EventMap: sm.EventMap{
				CreateOrder: orderCreating,
			},
		},
		orderCreating: {
			Action: &creatingOrderAction{},
			EventMap: sm.EventMap{
				FailOrder:  orderFailed,
				PlaceOrder: orderPlaced,
			},
		},
		orderFailed: {
			Action: &orderFailedAction{},
			EventMap: sm.EventMap{
				CreateOrder: orderCreating,
			},
		},
		orderPlaced: {
			Action: &orderPlacedAction{},
			EventMap: sm.EventMap{
				ChargeCard: orderChargingCard,
			},
		},
		orderChargingCard: {
			Action: &orderChargingCardAction{},
			EventMap: sm.EventMap{
				FailTransaction: orderTransactionFailed,
				ShipOrder:       orderShipped,
			},
		},
		orderTransactionFailed: {
			Action: &orderTransactionFailedAction{},
			EventMap: sm.EventMap{
				ChargeCard: orderChargingCard,
			},
		},
		orderShipped: {
			Action: &orderShippedAction{},
			EventMap: sm.EventMap{
				DeliverOrder: orderDelivered,
			},
		},
		orderDelivered: {
			Action: &orderDeliveredAction{},
		},
	}, opts...)
}
