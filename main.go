package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
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

	// get repository information
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repoInfo, _, err := client.Repositories.Get(ctx, *org, *repo)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	var titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true)
	var nameStyle = lipgloss.NewStyle().
		Width(18).
		Align(lipgloss.Left).
		Foreground(lipgloss.Color("#FFFDF5"))
	var detailStyle = lipgloss.NewStyle().
		Align(lipgloss.Left).
		Bold(true)

	name := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Name: "), detailStyle.Render(repoInfo.GetName()))
	description := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Description: "), detailStyle.Render(repoInfo.GetDescription()))
	stars := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Stars: "), detailStyle.Render(strconv.Itoa(repoInfo.GetStargazersCount())))
	language := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Language: "), detailStyle.Render(repoInfo.GetLanguage()))
	forks := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Forks: "), detailStyle.Render(strconv.Itoa(repoInfo.GetForksCount())))
	openIssues := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Open Issues: "), detailStyle.Render(strconv.Itoa(repoInfo.GetOpenIssuesCount())))
	subscribers := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Subscribers: "), detailStyle.Render(strconv.Itoa(repoInfo.GetSubscribersCount())))
	network := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Network: "), detailStyle.Render(strconv.Itoa(repoInfo.GetNetworkCount())))
	watchers := lipgloss.JoinHorizontal(lipgloss.Top, nameStyle.Render("Watchers: "), detailStyle.Render(strconv.Itoa(repoInfo.GetWatchersCount())))

	fmt.Println(titleStyle.Render("Repository Information"))
	fmt.Println(name)
	fmt.Println(description)
	fmt.Println(stars)
	fmt.Println(language)
	fmt.Println(forks)
	fmt.Println(openIssues)
	fmt.Println(subscribers)
	fmt.Println(network)
	fmt.Println(watchers)

}
