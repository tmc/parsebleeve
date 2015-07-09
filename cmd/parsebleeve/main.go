package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/blevesearch/bleve"
)

func hello(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var m map[string]interface{}
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Println(m)
	io.WriteString(w, `{ "success" : true}`)
}

func main() {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	bleeveMain()
	http.ListenAndServe(":"+port, mux)
}

func bleeveMain() {
	// open a new index
	index, err := bleve.Open("content.bleve")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}
	index.Index("1", data)

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
