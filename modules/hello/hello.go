package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	attendeeNames := []string{
		"Scarlett",
		"Valeria",
		"Iwia",
	}

	// Get a greeting message and print it
	messages, err := greetings.Hellos(attendeeNames)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
