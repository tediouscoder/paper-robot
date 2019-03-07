package github

import (
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	client *github.Client
)

func init() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(nil, ts)

	client = github.NewClient(tc)
}
