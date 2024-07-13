package main

import (
	"fmt"
	"log"
)

func displayHeader(number int, title string) {
	fmt.Println()
	fmt.Printf("-- %d. %s", number, title)
	fmt.Println()
}

func displayDescription(description string) {
	fmt.Println()
	fmt.Printf("---- %s", description)
}

func displayError(err error) {
	fmt.Println()
	log.Fatalln(`
=======================================
ERROR: ` + err.Error() + `
=======================================
	`)
}
