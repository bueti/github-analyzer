package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	addrFlag := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	addr := ":" + *addrFlag

	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: true,
		ReportCaller:    true,
	})

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/v1/repo/view", repoView)

	logger.Info("Starting server", "addr", *addrFlag)
	err := http.ListenAndServe(addr, mux)
	logger.Fatal(err)
}
