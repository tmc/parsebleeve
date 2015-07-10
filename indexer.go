package parsesearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/blevesearch/bleve"
)

type indexer struct {
	index      bleve.Index
	webhookKey string
}

func NewIndexer(webhookKey string) (*indexer, error) {
	path := "contents.bleve"
	i, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		im := bleve.NewIndexMapping()
		i, err = bleve.New(path, im)
	} else if err != nil {
		return nil, err
	}
	return &indexer{
		index:      i,
		webhookKey: webhookKey,
	}, nil
}

func (i *indexer) Index(w http.ResponseWriter, r *http.Request) {
	body := WebhookRequest{}
	buf := &bytes.Buffer{}
	io.Copy(buf, r.Body)
	defer r.Body.Close()
	err := json.NewDecoder(buf).Decode(&body)
	if err != nil {
		writeErr(w, err)
		return
	}
	if r.Header.Get("X-Parse-Webhook-Key") != i.webhookKey {
		writeErr(w, fmt.Errorf("invalid webhook key"))
		return
	}
	//TODO(tmc): guard these casts
	obj := body.Object.(map[string]interface{})
	err = i.index.Index(obj["objectId"].(string), obj)
	if err != nil {
		writeErr(w, err)
		return
	}
	json.NewEncoder(os.Stdout).Encode(body)

	err = json.NewEncoder(w).Encode(Response{
		Success: true,
	})
	if err != nil {
		log.Println("error writing response:", err)
	}
}
func (i *indexer) Search(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		writeErr(w, err)
		return
	}
	q := r.Form.Get("q")
	if q == "" {
		writeErr(w, fmt.Errorf("no term provided"))
		return
	}
	query := bleve.NewMatchQuery(q)
	search := bleve.NewSearchRequest(query)
	searchResults, err := i.index.Search(search)
	if err != nil {
		writeErr(w, err)
		return
	}
	err = json.NewEncoder(w).Encode(Response{
		Success: searchResults,
	})
	if err != nil {
		log.Println("error encoding response:", err)
	}
}
