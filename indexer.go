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
	"github.com/tmc/parse"
)

type Indexer struct {
	index bleve.Index

	// Parse Keys
	webhookKey string
	masterKey  string
	appID      string
}

func NewIndexer(webhookKey, masterKey, appID string) (*Indexer, error) {
	path := "contents.bleve"
	i, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		im := bleve.NewIndexMapping()
		i, err = bleve.New(path, im)
	} else if err != nil {
		return nil, err
	}
	return &Indexer{
		index:      i,
		webhookKey: webhookKey,
		masterKey:  masterKey,
		appID:      appID,
	}, nil
}

func webhookRequest(r *http.Request, webhookKey string) (*WebhookRequest, error) {
	req := &WebhookRequest{}
	buf := &bytes.Buffer{}
	io.Copy(buf, r.Body)
	defer r.Body.Close()
	err := json.NewDecoder(buf).Decode(&req)
	if err != nil {
		return nil, err
	}
	if r.Header.Get("X-Parse-Webhook-Key") != webhookKey {
		return nil, fmt.Errorf("invalid webhook key")
	}
	return req, nil
}

func (i *Indexer) Index(w http.ResponseWriter, r *http.Request) {
	req, err := webhookRequest(r, i.webhookKey)
	if err != nil {
		writeErr(w, err)
		return
	}
	//TODO(tmc): guard these casts
	obj := req.Object.(map[string]interface{})
	err = i.index.Index(obj["objectId"].(string), obj)
	if err != nil {
		writeErr(w, err)
		return
	}
	json.NewEncoder(os.Stdout).Encode(req)

	err = json.NewEncoder(w).Encode(Response{
		Success: true,
	})
	if err != nil {
		log.Println("error writing response:", err)
	}
}

func (i *Indexer) Unindex(w http.ResponseWriter, r *http.Request) {
	req, err := webhookRequest(r, i.webhookKey)
	if err != nil {
		writeErr(w, err)
		return
	}
	//TODO(tmc): guard these casts
	obj := req.Object.(map[string]interface{})
	err = i.index.Delete(obj["objectId"].(string))
	if err != nil {
		writeErr(w, err)
		return
	}
	json.NewEncoder(os.Stdout).Encode(req)

	err = json.NewEncoder(w).Encode(Response{
		Success: true,
	})
	if err != nil {
		log.Println("error writing response:", err)
	}
}

func (i *Indexer) Search(w http.ResponseWriter, r *http.Request) {
	req, err := webhookRequest(r, i.webhookKey)
	if err != nil {
		writeErr(w, err)
		return
	}
	rawq := req.Params["q"]
	if rawq == nil {
		writeErr(w, fmt.Errorf("no term provided"))
		return
	}
	q, ok := rawq.(string)
	if q == "" || !ok {
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
	ids := []string{}
	for _, h := range searchResults.Hits {
		ids = append(ids, h.ID)
	}
	err = json.NewEncoder(w).Encode(Response{
		Success: ids,
	})
	if err != nil {
		log.Println("error encoding response:", err)
	}
}

// fetches all objects and inedexes them again. long running
func (i *Indexer) Reindex(className string) error {
	client, err := parse.NewClient(i.appID, "")
	if err != nil {
		return err
	}
	client.TraceOn(log.New(os.Stderr, "[parse api] ", log.LstdFlags))
	client = client.WithMasterKey(i.masterKey)
	iter, err := client.NewScanner(className, "{}")
	if err != nil {
		return err
	}
	for o := iter.Next(); o != nil; o = iter.Next() {
		obj := o.(map[string]interface{})
		if err := i.index.Index(obj["objectId"].(string), obj); err != nil {
			log.Println("index error:", err)
		}
	}
	return iter.Err()
}
