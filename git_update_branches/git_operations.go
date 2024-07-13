package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func gitFetch() {
	cmd := exec.Command("git", "fetch", "--all")
	_, err := cmd.Output()
	if err != nil {
		displayError(err, "Unable to fetch all branches")
	}
}

func gitSwitchTo(branchName string) {
	// Ensure that the branch is a local branch
	branchToSwitchTo := strings.TrimPrefix(branchName, "origin/")

	cmd := exec.Command("git", "switch", branchToSwitchTo)
	_, err := cmd.Output()
	if err != nil {
		displayError(err, "Unable to switch to branch:", branchToSwitchTo)
	}
}

func gitGetAllBranches() []string {
	cmd := exec.Command("git", "branch", "-a")
	output, err := cmd.Output()
	if err != nil {
		displayError(err, "Unable to get all branches")
	}

	branches := strings.Split(string(output), "\n")

	return branches
}

func gitMerge(branchName string) {
	_, err := exec.Command("git", "merge", branchName).Output()
	if err != nil {
		displayError(err, "Unable to merge branch:", branchName)
	}
}

func gitPullFastForward() {
	output, err := exec.Command("git", "pull", "--ff-only").Output()
	fmt.Println(string(output))
	if err != nil {
		displayError(err, "Unable to pull fast-forward")
	}
}

func gitPush() {
	_, err := exec.Command("git", "push").Output()
	if err != nil {
		displayError(err, "Unable to push")
	}
}
