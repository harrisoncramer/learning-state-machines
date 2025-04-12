package examples

import (
	"fmt"
	"log"

	light_switch "github.com/harrisoncramer/learning-state-machines/light_switch"
)

func LightSwitch() {
	ls := light_switch.NewLightSwitch()

	fmt.Printf("The initial state is: %s\n", ls.GetCurrentState())

	err := ls.SendEvent(light_switch.SwitchOn, nil) // No additional context is needed
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The current state is: %s\n", ls.GetCurrentState())

	err = ls.SendEvent(light_switch.SwitchOff, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The current state is: %s\n", ls.GetCurrentState())
}
