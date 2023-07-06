package utils

import (

	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)


func getClient()(*github.Client){

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "github Token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
	
}