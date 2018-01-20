# gsc

Handy little command for cloning git repositories to a standard location from anywhere. Basically solves a first world problem.

## problem

Go has `go get` for downloading go packages and their dependancies from version control systems. It works great for go projects. 10/10 would `go get` again.

My first world problem is that I'm too lazy to type out full filepaths when cloning non-go repos. I've become so accustomed to the GOPATH way of structuring source code that I want to clone all repos to a path that conforms to my GOPATH regardless of whether or not the repo is a go project. IE: `$HOME/src/github.com/joepurdy/not-a-go-package`

## solution

`gsc git@github.com:joepurdy/not-a-go-package`

First world problem solved with <100 lines of go. Now I can pass standard git repo paths to my `gsc` command and clone them to a filepath conforming to my GOPATH. 