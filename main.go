/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/DelcioKelson/github_cli/cmd"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
        	port = "8080"
   	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	
	cmd.Execute()
}
