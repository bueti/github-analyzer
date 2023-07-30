package models

import (
	"context"

	"github.com/google/go-github/v53/github"
)

type Repo struct {
	Name        string
	Description string
	Stars       int
	Language    string
	Forks       int
	OpenIssues  int
	Subscribers int
	Watchers    int
}

type RepoModel struct {
	Client *github.Client
	Ctx    context.Context
}

func (r *RepoModel) Get(org, name string) (*Repo, error) {
	repo := &Repo{}
	repoInfo, _, err := r.Client.Repositories.Get(r.Ctx, org, name)
	if err != nil {
		return nil, err
	}
	repo.Name = repoInfo.GetName()
	repo.Description = repoInfo.GetDescription()
	repo.Stars = repoInfo.GetStargazersCount()
	repo.Language = repoInfo.GetLanguage()
	repo.Forks = repoInfo.GetForksCount()
	repo.OpenIssues = repoInfo.GetOpenIssuesCount()
	repo.Subscribers = repoInfo.GetSubscribersCount()
	repo.Watchers = repoInfo.GetWatchersCount()

	return repo, nil
}
