package main

import (
	"fmt"
	"gocode.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Rachit")
	fmt.Println(message)
}
