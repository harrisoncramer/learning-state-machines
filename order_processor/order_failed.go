package order_processor

import (
	"fmt"

	"github.com/harrisoncramer/learning-state-machines/sm"
)

type orderFailedAction struct{}

func (a *orderFailedAction) Execute(eventCtx sm.EventContext) (sm.Event, error) {
	order, ok := eventCtx.(OrderCreationContext)
	if !ok {
		return "", fmt.Errorf("%w: order_failed got %T but expected OrderCreationContext", ErrBadExecutionContext, order)
	}
	fmt.Println("Order failed, err:", order.err)
	return sm.NoOp, nil
}
