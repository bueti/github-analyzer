package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/v1/repo/view", repoView)

	log.Info("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
