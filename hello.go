package main

import (
	"fmt"
	"github.com/GregorSondermeier/greetings-go"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gregor")
	fmt.Println(message)
}
