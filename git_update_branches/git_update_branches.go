package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	args, err := getArgs()
	if err != nil {
		displayError(err)
	}

	gitFetch()

	branchNames, err := getBranchNames(args)
	if err != nil {
		displayError(err)
	}

	fmt.Println(branchNames)
}

func gitFetch() {
	fmt.Println("-- fetching branches")
	cmd := exec.Command("git", "fetch", "--all")
	_, err := cmd.Output()
	if err != nil {
		displayError(err)
	}
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

func getBranchNames(args []string) ([]string, error) {
	fmt.Println("-- getting remote branch names")
	cmd := exec.Command("git", "branch", "-r")
	output, err := cmd.Output()

	if err != nil {
		return []string{}, errors.New("git branch command failed")
	}

	branches := strings.Split(string(output), "\n")
	fmt.Println("-- getting branch names of args:", args)

	fullBranchNames := []string{}
	for _, arg := range args {
		fullBranchName, err := getFullBranchName(arg, branches)
		if err != nil {
			displayError(err)
		}
		fullBranchNames = append(fullBranchNames, fullBranchName)
	}

	return fullBranchNames, nil
}

func getFullBranchName(shortName string, branches []string) (string, error) {
	for _, branch := range branches {
		if strings.Contains(branch, shortName) {
			return strings.TrimSpace(strings.TrimPrefix(branch, "*")), nil
		}
	}

	return "", errors.New("branch (" + shortName + ") not found")
}

func displayError(err error) {
	log.Fatal(`
=======================================
ERROR: ` + err.Error() + `
=======================================
	`)
}
