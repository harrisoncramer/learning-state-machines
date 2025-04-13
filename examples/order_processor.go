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
	ctx := context.WithValue(context.Background(), "order", &order_processor.OrderCreationContext{
		Items: []string{"hat", "shoe"},
	})
	err = op.SendEvent(ctx, order_processor.CreateOrder)
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	fmt.Printf("The next state is: %s\n", op.GetCurrentState())

}
