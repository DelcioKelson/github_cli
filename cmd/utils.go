package cmd

import (
	"fmt"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/spf13/viper"

)

func getClient(ctx context.Context)(*github.Client){

	for {
		if !viper.IsSet("github_token") {
			fmt.Println("Set you github token:")
			var token string
			fmt.Scanln(&token)
			viper.Set("github_token", token)
		}

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: viper.GetViper().GetString("github_token")},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		_, _, err := client.Users.Get(ctx, "")

		if err != nil {
			if _, ok := err.(*github.ErrorResponse); ok {
				fmt.Println("Invalid GitHub token. Please provide a valid token.")
				viper.Set("github_token", nil)
				continue
			}
		}	
		viper.WriteConfig()
		return client

	}

	return nil	
}

func printRepositories(repos []*github.Repository){
	// Iterate over the repositories and print their names
	for _, repo := range repos {
		fmt.Println(repo.GetFullName())
	}
}

func getRepositories(targetUser string, ctx context.Context) ([]*github.Repository){

	client := getClient(ctx)

	// List all repositories for the target user
	repos, _, err := client.Repositories.List(ctx, targetUser, nil)
	if err != nil {
		fmt.Printf("Error retrieving repositories: %v\n", err)
		return nil
	}

	return repos
}

