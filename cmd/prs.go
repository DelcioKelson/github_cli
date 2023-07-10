/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sort"
	"context"
	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
)


// getPrsCmd represents the getPrs command
var getPrsCmd = &cobra.Command{
	Use:   "prs",
	Short: "Get a list of pull requests for the target user",
	Long: `The "prs" command allows you to retrieve a list of 5 
	latest pull requests associated with the target user from 
	GitHub. It provides detailed information about each pull
	request,including the title, repository name, creation date, and status.
	
Usage:
github_cli prs [target user]
	
Example:
github_cli prs delciokelson`,

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		repos := getRepositories(args[0],ctx)
		prs := getPullRequests(repos, args[0], ctx)
		print5PullRequests(prs)
	},
}

func init() {
	rootCmd.AddCommand(getPrsCmd)
}

func getPullRequests(repos []*github.Repository, targetUser string, ctx context.Context) []*github.PullRequest{

	client := getClient(ctx)

	var allPullRequests []*github.PullRequest

	// Iterate over the repositories
	for _, repo := range repos {
		// Get pull requests for each repository
		pullRequests, _, err := client.PullRequests.List(ctx, targetUser, *repo.Name, nil)
		if err != nil {
			fmt.Printf("Error retrieving pull requests for repository %s: %v\n", *repo.Name, err)
			continue
		}

		allPullRequests = append(allPullRequests, pullRequests...)
	}


	// sort pull requests by their creation data
	sort.Slice(allPullRequests, func(i, j int) bool {
		return allPullRequests[i].CreatedAt.After(*allPullRequests[j].CreatedAt)
		})

	return allPullRequests
}


