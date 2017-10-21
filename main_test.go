package main

import (
	"testing"
)

func TestParsesHTTPS(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"https://github.com/freeCodeCamp/freeCodeCamp.git", "github.com/freeCodeCamp/freeCodeCamp"},
		{"https://github.com/angular/angular.js.git", "github.com/angular/angular.js"},
		{"https://github.com/robbyrussell/oh-my-zsh.git", "github.com/robbyrussell/oh-my-zsh"},
		{"https://github.com/daneden/animate.css.git", "github.com/daneden/animate.css"},
		{"https://gitlab.com/gitlab-org/gitlab-ce.git", "gitlab.com/gitlab-org/gitlab-ce"},
		{"https://gitlab.com/gnachman/iterm2.git", "gitlab.com/gnachman/iterm2"},
		{"https://somerandomdomain.xyz/hackerman/1337.git", "somerandomdomain.xyz/hackerman/1337"},
		{"https://notagitrepo.com", ""},
	}
	for _, test := range tests {
		if got, _ := parseGitRepo(test.input); got != test.want {
			t.Errorf("parseGitRepo(%q) = %v", test.input, got)
		}
	}
}

func TestParsesSSH(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"git@github.com:freeCodeCamp/freeCodeCamp.git", "github.com/freeCodeCamp/freeCodeCamp"},
		{"git@github.com:angular/angular.js.git", "github.com/angular/angular.js"},
		{"git@github.com:robbyrussell/oh-my-zsh.git", "github.com/robbyrussell/oh-my-zsh"},
		{"git@github.com:daneden/animate.css.git", "github.com/daneden/animate.css"},
		{"git@gitlab.com:gitlab-org/gitlab-ce.git", "gitlab.com/gitlab-org/gitlab-ce"},
		{"git@gitlab.com:gnachman/iterm2.git", "gitlab.com/gnachman/iterm2"},
		{"sysadmin@somerandomdomain.xyz:hackerman/1337.git", "somerandomdomain.xyz/hackerman/1337"},
		{"user@notagitrepo.com", ""},
	}
	for _, test := range tests {
		if got, _ := parseGitRepo(test.input); got != test.want {
			t.Errorf("parseGitRepo(%q) = %v", test.input, got)
		}
	}
}

func TestParsesKeybase(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"keybase://private/user/repo", "keybase/user/repo"},
		{"keybase://team/teamname/repo", "keybase/teamname/repo"},
	}
	for _, test := range tests {
		if got, _ := parseGitRepo(test.input); got != test.want {
			t.Errorf("parseGitRepo(%q) = %v", test.input, got)
		}
	}
}
