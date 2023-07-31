package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/bueti/github-analyzer/internal/models"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	repos         *models.RepoModel
	templateCache map[string]*template.Template
}

func main() {
	addrFlag := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	addr := ":" + *addrFlag

	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()

	infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00d7ff"))
	errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff005f"))
	infoStr := infoStyle.Render("INFO\t")
	errorStr := errorStyle.Render("ERROR\t")
	infoLog := log.New(os.Stdout, fmt.Sprintf("%s ", infoStr), log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, fmt.Sprintf("%s ", errorStr), log.Ldate|log.Ltime|log.Lshortfile)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	// Initialize a new template cache...
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		repos: &models.RepoModel{
			Client: github.NewClient(tc),
			Ctx:    ctx,
		},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on port %s", *addrFlag)
	err = srv.ListenAndServe()
	app.infoLog.Fatal(err)
}
