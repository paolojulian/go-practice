package main

import (
	"fmt"
	"log"
	"strings"
)

func displayHeader(number int, title string) {
	fmt.Printf("\n-- %d. %s", number, title)
}

func displayDescription(description string) {
	fmt.Println()
	fmt.Printf("---- %s", description)
}

func displayError(err error, description ...string) {
	fmt.Println()
	fmt.Println("=======================================")
	log.Fatalln(`
ERROR: ` + err.Error() + `
Description: ` + strings.Join(description, " ") + `
=======================================
	`)
}
