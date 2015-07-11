package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sheki/parsesearch"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		whkey     = os.Getenv("PARSE_WEBHOOK_KEY")
		mkey      = os.Getenv("PARSE_MASTER_KEY")
		appid     = os.Getenv("PARSE_APPLICATION_ID")
		className = os.Getenv("PARSE_CLASS_NAME")
	)
	if whkey == "" || mkey == "" || appid == "" || className == "" {
		log.Fatalln("Must provide PARSE_WEBHOOK_KEY, PARSE_MASTER_KEY, PARSE_APPLICATION_ID, and PARSE_CLASS_NAME")
	}
	if port == "" {
		port = "8000"
	}
	i, err := parsesearch.NewIndexer(whkey, mkey, appid)
	if err != nil {
		log.Fatalln("error creating Indexer:", err)
	}
	if err = i.RegisterHooks(className); err != nil {
		fmt.Println("error creating hooks:", err)
	}
	go i.Reindex(className)
	mux := http.NewServeMux()
	mux.HandleFunc("/search", i.Search)
	mux.HandleFunc("/index", i.Index)
	mux.HandleFunc("/unindex", i.Unindex)
	if ui, err := parsesearch.NewUI(appid, "JAVSCRIPT KEY", className); err == nil {
		mux.Handle("/", ui)
	} else {
		fmt.Println("error constructing ui:", err)
	}
	http.ListenAndServe(":"+port, mux)
}
