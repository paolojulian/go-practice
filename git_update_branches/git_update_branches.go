package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"slices"
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
	// Get all branches
	fmt.Println("-- getting all branch names")
	cmd := exec.Command("git", "branch", "-a")
	output, err := cmd.Output()
	if err != nil {
		return []string{}, errors.New("git branch command failed")
	}

	// Get the full branch names of the args
	fmt.Println("-- getting branch names of args:", args)

	fullBranchNames := []string{}
	branches := strings.Split(string(output), "\n")
	// Reverse the branches so we look for the remote branches first
	slices.Reverse(branches)

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
			trimmedSpaces := strings.TrimSpace(branch)
			removedAsterisk := strings.TrimPrefix(trimmedSpaces, "*")
			removedRemotes := strings.TrimPrefix(removedAsterisk, "remotes/")

			return removedRemotes, nil
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
