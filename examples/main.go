package examples

import (
	"log"
)

func Run(program string) {

	switch program {
	case "light_switch":
		LightSwitch()
	case "order_processor":
		OrderProcessor()
	default:
		log.Fatalf("%s is not a valid program", program)
	}
}
