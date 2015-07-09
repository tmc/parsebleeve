package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/davecgh/go-spew/spew"
)

type index struct {
	index bleve.Index
}

func hello(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var m map[string]interface{}
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Println(m)
	io.WriteString(w, `{"success" : true}`)
}

func main() {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	i, err := NewIndex("content.bleve")
	if err != nil {
		log.Fatalln("error creating index:", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/search", i.searchHandler)
	mux.HandleFunc("/index", i.indexHandler)
	http.ListenAndServe(":"+port, mux)
}

type Response struct {
	Error   interface{} `json:"error,omitempty"`
	Success interface{} `json:"success,omitempty"`
}

type WebhookRequest struct {
	InstallationId string      `json:"installationId,omitempty"`
	Master         bool        `json:"master,omitempty"`
	Object         interface{} `json:"object,omitempty"`
	TriggerName    string      `json:"triggerName,omitempty"`
}

func NewIndex(path string) (*index, error) {
	i, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		im := bleve.NewIndexMapping()
		i, err = bleve.New(path, im)
	} else if err != nil {
		return nil, err
	}
	return &index{
		index: i,
	}, nil
}

func writeErr(w io.Writer, msg error) {
	spew.Dump("writing err", msg)
	err := json.NewEncoder(w).Encode(Response{Error: msg.Error()})
	if err != nil {
		log.Println("error encoding response:", err)
	}
}

func (i *index) indexHandler(w http.ResponseWriter, r *http.Request) {
	body := WebhookRequest{}

	buf := &bytes.Buffer{}
	io.Copy(buf, r.Body)
	defer r.Body.Close()
	err := json.NewDecoder(buf).Decode(&body)
	if err != nil {
		writeErr(w, err)
	}
	obj := body.Object.(map[string]interface{})
	err = i.index.Index(obj["objectId"].(string), obj)
	if err != nil {
		writeErr(w, err)
	}

	err = json.NewEncoder(w).Encode(Response{
		Success: true,
	})
	if err != nil {
		log.Println("error writing response:", err)
	}
}
func (i *index) searchHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		writeErr(w, err)
		return
	}
	q := r.Form.Get("q")
	spew.Dump(r)
	if q == "" {
		writeErr(w, fmt.Errorf("no term provided"))
		return
	}
	query := bleve.NewMatchQuery(q)
	search := bleve.NewSearchRequest(query)
	searchResults, err := i.index.Search(search)
	if err != nil {
		writeErr(w, err)
	}
	err = json.NewEncoder(w).Encode(Response{
		Success: searchResults,
	})
	if err != nil {
		log.Println("error encoding response:", err)
	}
}
