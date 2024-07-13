package main

import (
	"fmt"
	"log"
	"strings"
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

func displayError(err error, description ...string) {
	fmt.Println()
	fmt.Println("=======================================")
	log.Fatalln(`
ERROR: ` + err.Error() + `
Description: ` + strings.Join(description, " ") + `
=======================================
	`)
}
