package main

import (
	"fmt"
	"log"
)

func displayHeader(number int, title string) {
	fmt.Printf("-- %d. %s", number, title)
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
