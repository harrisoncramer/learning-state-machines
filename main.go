package main

import (
	"fmt"
	"os"

	"github.com/harrisoncramer/learning-state-machines/examples"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide exactly one argument.")
		os.Exit(1)
	}
	program := os.Args[1]
	examples.Run(program)
}
