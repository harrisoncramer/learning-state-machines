package examples

import (
	"fmt"
	"log"

	"github.com/harrisoncramer/learning-state-machines/order_processor"
	"github.com/harrisoncramer/learning-state-machines/sm"
	logger "github.com/harrisoncramer/learning-state-machines/sm/logger"
)

func OrderProcessor() {
	op, err := order_processor.NewOrderProcessor(setLogger)
	if err != nil {
		log.Fatalf("failed to set up order processor: %v", err)
	}

	fmt.Printf("The initial state is: %s\n", op.GetCurrentState())

	err = op.SendEvent(order_processor.CreateOrder, order_processor.OrderCreationContext{
		Items: []string{
			"hat",
			"shoe",
			"can",
		},
	})
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	fmt.Printf("The next state is: %s\n", op.GetCurrentState())

}

// setLogger sets the debug logger on the state machine
var setLogger sm.Option = func(sm *sm.StateMachine) error {
	l, err := logger.GetZapLogger("debug", "")
	if err != nil {
		return err
	}
	sm.SetLogger(l)
	return nil
}
