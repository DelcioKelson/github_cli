/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// getReposCmd represents the getRepos command
var getReposCmd = &cobra.Command{
	Use:   "getRepos",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepos(args[0])

	},
}

func init() {
	rootCmd.AddCommand(getReposCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getReposCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getReposCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func getRepos(targetUser string){

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Replace "TARGET_USER" with the username of the target user

	// List all repositories for the target user
	repos, _, err := client.Repositories.List(ctx, targetUser, nil)
	if err != nil {
		fmt.Printf("Error retrieving repositories: %v\n", err)
		return
	}

	for _, repo := range repos {
		fmt.Println(repo.GetFullName())
	}
}