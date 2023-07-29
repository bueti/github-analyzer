package main

import (
	"context"
	"flag"
	"os"

	"github.com/charmbracelet/log"
)

func main() {

	// read GITHUB_TOKEN from environment variable
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Error("GITHUB_TOKEN environment variable is not set")
		os.Exit(1)
	}

	org := flag.String("org", "", "organization name")
	repo := flag.String("repo", "", "repository name")

	flag.Parse()
	if *org == "" {
		log.Error("organization name is not set")
		os.Exit(1)
	}
	if *repo == "" {
		log.Error("repository name is not set")
		os.Exit(1)
	}

	ctx := context.Background()

	printRepositoryInfo(ctx, *org, *repo, token)

}
