package main

import (
	"fmt"
	"golang/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	// Request a greeting message.
	// message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
