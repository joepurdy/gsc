package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s [git repo]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nParameters:\n\t[git repo] - valid git repository url. Accepts HTTPS or SSH protocols\n")
}

var parseHelp = func() bool {
	switch os.Args[1] {
	case "--help":
		return true
	case "-help":
		return true
	case "-h":
		return true
	default:
		return false
	}
}

func main() {
	if len(os.Args) < 2 || parseHelp() {
		usage()
		return
	}
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
