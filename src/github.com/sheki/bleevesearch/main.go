package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", hello)
	// http.ListenAndServe(":8000", mux)
	bleeveMain()
}

func bleeveMain() {

	// open a new index
	index, err := bleve.Open("example.bleve")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

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
