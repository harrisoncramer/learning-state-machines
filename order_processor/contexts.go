package order_processor

import (
	"fmt"
	"strings"
)

// OrderShipmentContext defines details about the shipping of the order
type OrderShipmentContext struct {
	cardNumber string
	address    string
	err        error
}

func (c *OrderShipmentContext) String() string {
	return fmt.Sprintf("OrderShipmentContext [ cardNumber: %s, address: %s ]", c.cardNumber, c.address)
}

// OrderCreationContext contains data about the actual order, including the items and any errors that pertain to it
type OrderCreationContext struct {
	Items []string
	err   error
}

func (c OrderCreationContext) String() string {
	return fmt.Sprintf("OrderCreationContext: items: %s; err: %s", strings.Join(c.Items, ","), c.err)
}
