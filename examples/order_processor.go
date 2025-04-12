package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/harrisoncramer/learning-state-machines/order_processor"
)

func OrderProcessor() {
	op, err := order_processor.NewOrderProcessor()
	if err != nil {
		log.Fatalf("failed to set up order processor: %v", err)
	}

	fmt.Printf("The initial state is: %s\n", op.GetCurrentState())
	err = op.SendEvent(context.Background(), order_processor.CreateOrder, &order_processor.OrderCreationContext{
		Items: []string{"hat", "shoe"},
	})
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	fmt.Printf("The next state is: %s\n", op.GetCurrentState())

}
