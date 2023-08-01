# Github Analyzer

## Description
This is a webapp and cli application that allows you to analyze a github repository. It uses the Github API to retrieve the data.

You can play with a deployed version here: https://gha.bueti-online.ch

## Installation

### Webapp
To run the webapp locally, checkout the repository and run the following command:
```
export GITHUB_TOKEN=<your github token>
docker compose up --build
```

### CLI
To install the cli app, run the following command:
```
go install github.com/bueti/github-analyzer@latest
```
## Usage
To use the application, run the following command:
```
export GITHUB_TOKEN=<your github token>
github-analyzer -org <organization> -repo <repository>
```
