package main

import (
	"fmt"
	"github.com/davidji99/bitbucket-go/bitbucket"
)

// state (active|archived|unknown)
// last modified
// last commit
// number of stars
// number of forks
// number of watches

type repository struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	ForksCount      int    `json:"forks_count"`
	OpenIssues      int    `json:"open_issues_count"`
	StargazersCount int    `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
	IsPrivate       bool   `json:"private"`
}

func main() {

	// Get repos for authenticated user

	client := bitbucket.NewBasicAuth("aliahsan07@outlook.com", "------")

	result, _, error := client.Repositories.List("aliahsan07")

	if error != nil {
		fmt.Println("error encountered")
	}

	repos := result.Values

	for _, repo := range repos {
		fmt.Println(repo.GetTitle())
		fmt.Println(repo.GetCreatedOn())
		fmt.Println(repo.GetEditedOn())
		fmt.Println(repo.GetState())
		fmt.Println(repo.GetUpdatedOn())
		fmt.Println(repo.GetWatches())
		fmt.Println(repo.GetRepository().GetSlug())

	}

	// repos, _, err := client.Repositories.List(ctx, "", nil)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// refinedRepos := make([]*repository, 0)

	// for _, t := range repos {
	// 	var repoDetails *repository = new(repository)

	// 	repoDetails.ID = t.GetID()
	// 	repoDetails.Name = t.GetName()
	// 	repoDetails.FullName = t.GetFullName()
	// 	repoDetails.ForksCount = t.GetForksCount()
	// 	repoDetails.OpenIssues = t.GetOpenIssuesCount()
	// 	repoDetails.StargazersCount = t.GetStargazersCount()
	// 	repoDetails.WatchersCount = t.GetWatchersCount()
	// 	repoDetails.IsPrivate = t.GetPrivate()
	// 	t.GetArchived()

	// 	refinedRepos = append(refinedRepos, repoDetails)
	// }

	// for _, repo := range refinedRepos {
	// 	fmt.Println("Name of repo: " + repo.FullName)
	// 	fmt.Println("Open issues: ", repo.OpenIssues)
	// 	fmt.Println("Count of stargazers: ", repo.StargazersCount)
	// 	fmt.Println("Number of forks: ", repo.ForksCount)
	// 	fmt.Println("Number of watchers: ", repo.WatchersCount)
	// 	fmt.Println("Is Private: ", repo.IsPrivate)
	// 	fmt.Println("===============================")
	// }

}
