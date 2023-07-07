package cmd

import (
	"testing"
	"context"
	"github.com/spf13/viper"
	"os"

)


func TestGetFollowedRepositories(t *testing.T)  {
	

	token := os.Getenv("TOKEN")
	viper.Set("github_token", token)

	// Scenario 1: Valid Scenario


	t.Run("Valid Scenario", func(t *testing.T) {
		targetUser := "johnDoe"
		ctx := context.Background()

		if res := len(getFollowedRepositories(targetUser,ctx)); res == 0{
			t.Errorf("Expected different instead of %d", res)
		}
	})
}