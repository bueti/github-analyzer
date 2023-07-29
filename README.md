# Github Analyzer

## Description
This is a simple cli application that allows you to analyze a github repository. It uses the Github API to retrieve the data.

## Installation
To install the application, run the following command:
```
go install github.com/bueti/github-analyzer@latest
```
## Usage
To use the application, run the following command:
```
export GITHUB_TOKEN=<your github token>
github-analyzer -org <organization> -repo <repository>
```
