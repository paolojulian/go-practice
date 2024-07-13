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
		fmt.Println()
		fmt.Println(cmd)
		displayError(err, "Unable to fetch all branches")
	}
}

func gitSwitchTo(branchName string) {
	// Ensure that the branch is a local branch
	branchToSwitchTo := strings.TrimPrefix(branchName, "origin/")

	cmd := exec.Command("git", "switch", branchToSwitchTo)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println()
		fmt.Println(cmd)
		displayError(err, "Unable to switch to branch:", branchToSwitchTo)
	}
}

func gitGetAllBranches() []string {
	cmd := exec.Command("git", "branch", "-a")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println()
		fmt.Println(cmd)
		displayError(err, "Unable to get all branches")
	}

	branches := strings.Split(string(output), "\n")

	return branches
}

func gitMerge(branchName string) {
	cmd := exec.Command("git", "merge", branchName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println()
		fmt.Println(cmd)
		displayError(err, "Unable to merge branch:", branchName)
	}
}

func gitPullFastForward() {
	cmd := exec.Command("git", "pull", "--ff-only")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println()
		fmt.Println(cmd)
		displayError(err, "Unable to pull fast-forward")
	}
}

func gitPush() {
	cmd := exec.Command("git", "push")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println()
		fmt.Println(cmd)
		displayError(err, "Unable to push")
	}
}
