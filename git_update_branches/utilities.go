package main

import (
	"fmt"
	"log"
)

func displayHeader(number int, title string) {
	fmt.Printf("-- %s. %s", string(number), title)
}

func displayDescription(description string) {
	fmt.Printf("---- %s", description)
}

func displayError(err error) {
	log.Fatal(`
=======================================
ERROR: ` + err.Error() + `
=======================================
	`)
}
