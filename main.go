package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	repo := os.Args[1]

	path, err := parseGitRepo(repo)
	checkErr(err)

	path = os.Getenv("HOME") + "/src/" + path
	cloneRepo(repo, path)
	fmt.Println(path)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseGitRepo(repo string) (string, error) {
	if !strings.HasSuffix(repo, ".git") {
		return "", fmt.Errorf("Unknown Repo Format: %s", repo)
	}

	repo = strings.TrimSuffix(repo, ".git")
	if strings.HasPrefix(repo, "https://") {
		repo = strings.TrimPrefix(repo, "https://")
	} else if strings.Contains(repo, "@") {
		repo = strings.Replace(strings.Split(repo, "@")[1], ":", "/", 1)
	}

	return repo, nil
}

func cloneRepo(repo string, path string) {
	cmd := exec.Command("git", "clone", repo, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkErr(err)
}
