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

	fmt.Println("-- 1. fetching branches")
	gitFetch()

	fmt.Println("-- 2. converting args to full branch names")
	branchNames, err := getBranchNames(args)
	if err != nil {
		displayError(err)
	}

	fmt.Println("-- 3. updating branches to latest change")
	for _, branchName := range branchNames {
		pullBranch(branchName)
	}

	fmt.Println("-- 4. merge dependent branches")
	mergeDependentBranches(branchNames)

	fmt.Println("-- 5. Done")
}

func gitFetch() {
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
	fmt.Println("---- git branch -a")
	// Get all branches
	cmd := exec.Command("git", "branch", "-a")
	output, err := cmd.Output()
	if err != nil {
		return []string{}, errors.New("'git branch -a' command failed")
	}

	// Get the full branch names of the args
	fullBranchNames := []string{}
	branches := strings.Split(string(output), "\n")
	// Reverse the branches so we look for the remote branches first
	slices.Reverse(branches)

	fmt.Println("---- mapping args to full branch names")
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

func pullBranch(branchName string) {
	branchToUpdate := strings.TrimPrefix(branchName, "origin/")
	fmt.Println("---- pulling branch:", branchToUpdate)
	gitSwitchTo(branchToUpdate)
	gitPullFastForward()
}

func mergeDependentBranches(branchNames []string) {
	currentBranch := branchNames[0]
	for index, branchName := range branchNames {
		// We skip the first branch since it's the base branch
		if index == 0 {
			continue
		}
		fmt.Println("---- merging branch:", currentBranch, "-->", branchName)
		gitSwitchTo(branchName)
		gitMerge(currentBranch)
		gitPush()
		currentBranch = branchName
	}
}

func gitSwitchTo(branchName string) {
	_, err := exec.Command("git", "switch", branchName).Output()
	if err != nil {
		displayError(err)
	}
}

func gitMerge(branchName string) {
	_, err := exec.Command("git", "merge", branchName).Output()
	if err != nil {
		displayError(err)
	}
}

func gitPullFastForward() {
	_, err := exec.Command("git", "pull", "--ff-only").Output()
	if err != nil {
		displayError(err)
	}
}

func gitPush() {
	_, err := exec.Command("git", "push").Output()
	if err != nil {
		displayError(err)
	}
}

func displayError(err error) {
	log.Fatal(`
=======================================
ERROR: ` + err.Error() + `
=======================================
	`)
}
