package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sheki/parsesearch"
)

func main() {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	i, err := parsesearch.NewIndexer(os.Getenv("PARSE_WEBHOOK_KEY"))
	if err != nil {
		log.Fatalln("error creating Indexer:", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/search", i.Search)
	mux.HandleFunc("/index", i.Index)
	http.ListenAndServe(":"+port, mux)
}
