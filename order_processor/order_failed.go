package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type OrderFailedAction struct{}

func (a *OrderFailedAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	order, ok := eventCtx.(*OrderCreationContext)
	if !ok {
		return "", ErrMissingOrderContext
	}
	fmt.Println("Order failed, err:", order.err)
	return sm.NoOp, nil
}
