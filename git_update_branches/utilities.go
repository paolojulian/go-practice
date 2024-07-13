package main

import (
	"fmt"
	"log"
)

func displayHeader(number int, title string) {
	fmt.Println()
	fmt.Printf("-- %d. %s", number, title)
}

func displayDescription(description string) {
	fmt.Println()
	fmt.Printf("---- %s", description)
}

func displayError(err error) {
	log.Fatalln(`
=======================================
ERROR: ` + err.Error() + `
=======================================
	`)
}
