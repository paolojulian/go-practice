package main

import (
	"os/exec"
	"strings"
)

func gitFetch() {
	cmd := exec.Command("git", "fetch", "--all")
	_, err := cmd.Output()
	if err != nil {
		displayError(err)
	}
}

func gitSwitchTo(branchName string) {
	// Ensure that the branch is a local branch
	branchToSwitchTo := strings.TrimPrefix(branchName, "origin/")

	_, err := exec.Command("git", "switch", branchToSwitchTo).Output()
	if err != nil {
		displayError(err)
	}
}

func gitGetAllBranches() []string {
	cmd := exec.Command("git", "branch", "-a")
	output, err := cmd.Output()
	if err != nil {
		displayError(err)
	}

	branches := strings.Split(string(output), "\n")

	return branches
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
