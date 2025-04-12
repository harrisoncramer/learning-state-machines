package order_processor

import (
	"github.com/harrisoncramer/learning-state-machines/sm"
)

// Possible states of the state machine for a given order, e.g. card is charging, order is shipped, etc
const (
	// The default state of the machine, no order has been placed yet
	OrderNotPlaced sm.State = ""
	// The order is being created
	OrderCreating sm.State = "OrderCreating"
	// The order creating failed and is now a failed order
	OrderFailed sm.State = "OrderFailed"
	// The order is placed successfully
	OrderPlaced sm.State = "OrderPlaced"
	// The credit card associated with the order is being charged
	OrderChargingCard sm.State = "ChargingCard"
	// The transaction (financial processing) failed
	OrderTransactionFailed sm.State = "OrderTransactionFailed"
	// The order has been shipped
	OrderShipped sm.State = "OrderShipped"
	// The order has been delivered successfully
	OrderDelivered sm.State = "OrderDelivered"
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

func NewOrderProcessor() *sm.StateMachine {
	return sm.NewStateMachine(OrderNotPlaced, sm.StateMap{
		OrderNotPlaced: {
			EventMap: sm.EventMap{
				PlaceOrder: OrderPlaced,
			},
		},
		OrderCreating: {
			Action: &CreatingOrderAction{},
			EventMap: sm.EventMap{
				FailOrder:  OrderFailed,
				PlaceOrder: OrderPlaced,
			},
		},
		OrderFailed: {
			Action: &OrderFailedAction{},
			EventMap: sm.EventMap{
				CreateOrder: OrderCreating,
			},
		},
		OrderPlaced: {
			Action: &OrderPlacedAction{},
			EventMap: sm.EventMap{
				ChargeCard: OrderChargingCard,
			},
		},
		OrderChargingCard: {
			Action: &ChargingCardAction{},
			EventMap: sm.EventMap{
				FailTransaction: OrderTransactionFailed,
				ShipOrder:       OrderShipped,
			},
		},
		OrderTransactionFailed: {
			Action: &TransactionFailedAction{},
			EventMap: sm.EventMap{
				ChargeCard: OrderChargingCard,
			},
		},
		OrderShipped: {
			Action: &OrderShippedAction{},
			EventMap: sm.EventMap{
				DeliverOrder: OrderDelivered,
			},
		},
		OrderDelivered: {
			Action: &OrderDeliveredAction{},
		},
	}, func(sm *sm.StateMachine) {
		// TODO: @harrisoncramer implement functional options pattern for state machine
	})
}
