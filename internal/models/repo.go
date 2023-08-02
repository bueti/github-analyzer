package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/go-github/v53/github"
)

var RepoNotFound = errors.New("repo not found")
var NotAuthorized = errors.New("not authorized")

type Repo struct {
	Org         string
	Name        string
	Description string
	Stars       int
	Language    string
	Forks       int
	OpenIssues  int
	Subscribers int
	CreatedAt   time.Time
}

type RepoModel struct {
	Client *github.Client
	Ctx    context.Context
}

func (r *RepoModel) Get(org, name string) (*Repo, error) {
	repo := &Repo{}
	repoInfo, resp, _ := r.Client.Repositories.Get(r.Ctx, org, name)
	log.Print(resp.StatusCode)
	if resp.StatusCode == 404 {
		return nil, RepoNotFound
	}
	if resp.StatusCode == 401 {
		return nil, NotAuthorized
	}

	repo.Org = org
	repo.Name = repoInfo.GetName()
	repo.Description = repoInfo.GetDescription()
	repo.Stars = repoInfo.GetStargazersCount()
	repo.Language = repoInfo.GetLanguage()
	repo.Forks = repoInfo.GetForksCount()
	repo.OpenIssues = repoInfo.GetOpenIssuesCount()
	repo.Subscribers = repoInfo.GetSubscribersCount()
	repo.CreatedAt = repoInfo.GetCreatedAt().UTC()

	return repo, nil
}
