package examples

import (
	"context"
	"fmt"
	"log"

	light_switch "github.com/harrisoncramer/learning-state-machines/light_switch"
)

func LightSwitch() {
	ls, err := light_switch.NewLightSwitch()
	if err != nil {
		log.Fatalf("failed to set up light switch: %v", err)
	}

	fmt.Printf("The initial state is: %s\n", ls.GetCurrentState())

	err = ls.SendEvent(context.Background(), light_switch.SwitchOn, nil) // No additional context is needed
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	fmt.Printf("The current state is: %s\n", ls.GetCurrentState())

	err = ls.SendEvent(context.Background(), light_switch.SwitchOff, nil)
	if err != nil {
		log.Fatalf("Failed to send event: %v", err)
	}

	fmt.Printf("The current state is: %s\n", ls.GetCurrentState())
}
