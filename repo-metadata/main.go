package main

import (
	"context"
	"fmt"

	"credentials"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type repository struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	ForksCount      int    `json:"forks_count"`
	OpenIssues      int    `json:"open_issues_count"`
	StargazersCount int    `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
}

func main() {

	// Get repos for authenticated user
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: credentials.GithubAccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)

	if err != nil {
		fmt.Println(err)
	}

	refinedRepos := make([]*repository, 0)

	for _, t := range repos {
		var repoDetails *repository = new(repository)

		repoDetails.ID = t.GetID()
		repoDetails.Name = t.GetName()
		repoDetails.FullName = t.GetFullName()
		repoDetails.ForksCount = t.GetForksCount()
		repoDetails.OpenIssues = t.GetOpenIssuesCount()
		repoDetails.StargazersCount = t.GetStargazersCount()
		repoDetails.WatchersCount = t.GetWatchersCount()

		refinedRepos = append(refinedRepos, repoDetails)
	}

	for _, repo := range refinedRepos {
		fmt.Println("Name of repo: " + repo.FullName)
		fmt.Println("Open issues: ", repo.OpenIssues)
		fmt.Println("Count of stargazers: ", repo.StargazersCount)
		fmt.Println("Number of forks: ", repo.ForksCount)
		fmt.Println("Number of watchers: ", repo.WatchersCount)
		fmt.Println("===============================")
	}

}
