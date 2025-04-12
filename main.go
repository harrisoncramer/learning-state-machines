package main

import (
	"fmt"
	"log"

	lightswitch "github.com/harrisoncramer/learning-state-machines/light_switch"
)

func main() {
	ls := lightswitch.NewLightSwitch()

	fmt.Printf("The initial state is: %s\n", ls.GetCurrentState())

	err := ls.SendEvent(lightswitch.SwitchOn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The current state is: %s\n", ls.GetCurrentState())

	err = ls.SendEvent(lightswitch.SwitchOff)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The current state is: %s\n", ls.GetCurrentState())
}
