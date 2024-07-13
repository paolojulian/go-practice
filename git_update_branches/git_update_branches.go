package main

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

const ARG_SPLITTER string = "/"

var gitOps GitOperations

func main() {
	args, err := getArgs()
	if err != nil {
		displayError(err)
	}

	gitOps = NewGitOps()

	displayHeader(1, "Fetching branches")
	if err := gitOps.Fetch(); err != nil {
		log.Fatal(err)
	}

	displayHeader(2, "Convert args to full branch names")
	branchNames, err := getBranchNames(args)
	if err != nil {
		displayError(err)
	}

	validateBranches(branchNames)

	displayHeader(3, "Updating branches to latest change")
	for _, branchName := range branchNames {
		pullBranch(branchName)
	}

	displayHeader(4, "Merge dependent branches")
	mergeDependentBranches(branchNames)

	displayHeader(5, "Finished")
}

func getArgs() ([]string, error) {
	argsWithoutProp := os.Args[1:]
	if len(argsWithoutProp) != 1 {
		return []string{}, errors.New("should contain exactly one arg")
	}

	args := argsWithoutProp[0]
	doesMatchFormat := regexp.MustCompile(`^([\w\d-]+)(\/[\w\d-]+)+$`).MatchString(args)
	if !doesMatchFormat {
		return []string{}, errors.New("invalid arg format, should be like 'master>developer>feature>feature-1'")
	}

	return strings.Split(args, ARG_SPLITTER), nil
}

func getBranchNames(args []string) ([]string, error) {
	displayDescription("Getting all branch names (git branch -a)")

	branches, err := gitOps.GetBranchNames()
	if err != nil {
		log.Fatal(err)
	}

	fullBranchNames := []string{}

	displayDescription("Mapping args to full branch names")
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

			return strings.TrimSpace(removedRemotes), nil
		}
	}

	return "", errors.New("No branch name matches: " + shortName)
}

func pullBranch(branchName string) {
	branchToUpdate := strings.TrimPrefix(branchName, "origin/")
	displayDescription("Pulling branch: " + branchToUpdate)
	gitOps.Switch(branchToUpdate)
	gitOps.Pull()
}

func mergeDependentBranches(branchNames []string) {
	currentBranch := branchNames[0]
	for index, branchName := range branchNames {
		// We skip the first branch since it's the base branch
		if index == 0 {
			continue
		}
		displayDescription("Merging branch: " + currentBranch + " --> " + branchName)
		gitOps.Switch(branchName)
		gitOps.Merge(currentBranch)
		gitOps.Push()
		currentBranch = branchName
	}
}
