package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sheki/parsesearch"
)

var parseKeys = []string{"PARSE_APPLICATION_ID", "PARSE_CLASS_NAME", "PARSE_JAVASCRIPT_KEY", "PARSE_MASTER_KEY", "PARSE_WEBHOOK_KEY"}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	keys := map[string]string{}
	for _, key := range parseKeys {
		val := os.Getenv(key)
		if val == "" {
			log.Fatalln("Must provide", key)
		}
		keys[key] = val
	}
	appID := keys["PARSE_APPLICATION_ID"]
	whkey := keys["PARSE_WEBHOOK_KEY"]
	mkey := keys["PARSE_MASTER_KEY"]
	jskey := keys["PARSE_JAVASCRIPT_KEY"]
	className := keys["PARSE_CLASS_NAME"]
	i, err := parsesearch.NewIndexer(whkey, mkey, appID)
	if err != nil {
		log.Fatalln("error creating Indexer:", err)
	}
	if err = i.RegisterHooks(className); err != nil {
		fmt.Println("error creating hooks:", err)
	}
	fmt.Println("start reindex job")
	go i.Reindex(className)
	mux := http.NewServeMux()
	mux.HandleFunc("/search", i.Search)
	mux.HandleFunc("/index", i.Index)
	mux.HandleFunc("/unindex", i.Unindex)
	mux.HandleFunc("/status", i.IndexStatus)
	if ui, err := parsesearch.NewUI(appID, jskey, className); err == nil {
		mux.Handle("/", ui)
	} else {
		fmt.Println("error constructing ui:", err)
	}
	http.ListenAndServe(":"+port, mux)
}
