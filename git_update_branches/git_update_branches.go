package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	args, err := getArgs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(args)
}

func getArgs() ([]string, error) {
	argsWithoutProp := os.Args[1:]
	if len(argsWithoutProp) != 1 {
		return []string{}, errors.New("should contain exactly one arg")
	}

	args := argsWithoutProp[0]
	doesMatchFormat := regexp.MustCompile(`^([\w\d-]+)(\/[\w\d-]+)+$`).MatchString(args)
	if !doesMatchFormat {
		return []string{}, errors.New("invalid arg format, should be like 'master/developer/feature/feature-1'")
	}

	return strings.Split(args, "/"), nil
}
