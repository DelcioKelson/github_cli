package utils

import (

	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)


func getClient()(*github.Client){

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
	
}