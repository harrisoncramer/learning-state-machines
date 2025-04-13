package order_processor

import (
	"github.com/google/uuid"
)

// Shipment defines details about the shipping of the order
type Shipment struct {
	CardNumber string
	Address    string
	Err        error
}

// Order contains data about the actual order, including the items and any errors that pertain to it
type Order struct {
	ID    uuid.UUID
	Items []string
	Err   error
}
