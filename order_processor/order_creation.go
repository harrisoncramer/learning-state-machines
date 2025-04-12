package order_processor

import (
	"fmt"
	"strings"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

// OrderCreationContext contains data about the actual order, including the items and any errors that pertain to it
type OrderCreationContext struct {
	items []string
	err   error
}

func (c OrderCreationContext) String() string {
	return fmt.Sprintf("OrderCreationContext: items: %s; err: %s", strings.Join(c.items, ","), c.err)
}

// CreatingOrderAction includes the action that fires when we enter the creating order state
type CreatingOrderAction struct{}

func (c *CreatingOrderAction) Execute(event sm.EventContext) sm.Event {
	order := event.(*OrderCreationContext)
	_ = order
	return sm.NoOp
}
