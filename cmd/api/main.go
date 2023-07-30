package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	addrFlag := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	addr := ":" + *addrFlag

	infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00d7ff"))
	errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff005f"))
	infoStr := infoStyle.Render("INFO\t")
	errorStr := errorStyle.Render("ERROR\t")
	infoLog := log.New(os.Stdout, fmt.Sprintf("%s ", infoStr), log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, fmt.Sprintf("%s ", errorStr), log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on port %s", *addrFlag)
	err := srv.ListenAndServe()
	app.infoLog.Fatal(err)
}
