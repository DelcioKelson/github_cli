/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"context"

	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// followsRepoCmd represents the followsRepo command
var followsRepoCmd = &cobra.Command{
	Use:   "followsRepo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getFollowedRepositories(args[0])
	},
}

func init() {
	rootCmd.AddCommand(followsRepoCmd)

}


func getFollowedRepositories(targetUser string){

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "github token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)


	repos, _, err := client.Activity.ListWatched(ctx, targetUser, nil)
	if err != nil {
		fmt.Printf("Error retrieving followed repositories: %v\n", err)
		return
	}

	// Iterate over the repositories and print their names
	for _, repo := range repos {
		fmt.Println(repo.GetFullName())
	}

}