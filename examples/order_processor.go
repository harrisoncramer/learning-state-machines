package examples

import (
	"log"

	"github.com/harrisoncramer/learning-state-machines/order_processor"
	"github.com/harrisoncramer/learning-state-machines/sm"
	logger "github.com/harrisoncramer/learning-state-machines/sm/logger"
)

func OrderProcessor() {

	op, err := order_processor.NewOrderProcessor(func(sm *sm.StateMachine) error {
		l, err := logger.GetZapLogger("debug", "")
		if err != nil {
			return err
		}
		sm.SetLogger(l)
		return nil
	})
	if err != nil {
		log.Fatalf("failed to set up order processor: %v", err)
	}

	err = op.SendEvent(order_processor.PlaceOrder, nil)
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	op.GetCurrentState()

}
