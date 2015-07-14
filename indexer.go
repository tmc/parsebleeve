package parsesearch

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/blevesearch/bleve"
	"github.com/tmc/parse"
)

// Indexer manages the search index for a Parse app.
type Indexer struct {
	index bleve.Index

	// Parse app keys/appID
	webhookKey string
	masterKey  string
	appID      string
	status     *indexStatusList
}

// NewIndexer prepares a new Indexer given the necessary Parse App credentials.
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
		status:     &indexStatusList{},
	}, nil
}

// Index is an http.HandlerFunc that accepts a parse afterSave webhook request.
//
// It adds or updates the provided objet in the search index.
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

// Unindex is an http.HandlerFunc that accepts a parse afterDelete webhook request.
//
// It removes the provided object from the index.
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

// Search is an http.HandlerFunc that accepts a Parse Cloud Code Webhook request.
//
// The expected query parameter is 'q'
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
	query := bleve.NewQueryStringQuery(q)
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

// Reindex fetches all objects for a class and indexes them.
// Could be long-running.
func (i *Indexer) Reindex(className string) error {
	//TODO(tmc): add counters/progress reporting
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
	status := &indexStatus{class: className}

	i.status.register(status)
	defer i.status.unregister(status)

	for o := iter.Next(); o != nil; o = iter.Next() {
		obj := o.(map[string]interface{})
		if err := i.index.Index(obj["objectId"].(string), obj); err != nil {
			log.Println("index error:", err)
		}

		log.Println("index incremented")
		status.incr()
	}
	return iter.Err()
}

// IndexStatus lists the status of all current Reindexing processes.
func (i *Indexer) IndexStatus(w http.ResponseWriter, r *http.Request) {
	i.status.writeIndexStatus(w)
}

type indexStatus struct {
	class string
	count uint32
}

func (i *indexStatus) incr() {
	atomic.AddUint32(&i.count, 1)
}

func (i *indexStatus) String() string {
	return fmt.Sprintf("indexed %d objects for class %s", i.count, i.class)
}

type indexStatusList struct {
	inProgress *list.List
	lock       sync.Mutex
	once       sync.Once
}

func (i *indexStatusList) register(status *indexStatus) {
	i.once.Do(func() {
		i.inProgress = list.New()
	})
	i.lock.Lock()
	defer i.lock.Unlock()
	i.inProgress.PushBack(status)
}

func (i *indexStatusList) unregister(status *indexStatus) {
	i.lock.Lock()
	defer i.lock.Unlock()
	var found *list.Element
	for e := i.inProgress.Front(); e != nil; e = e.Next() {
		if reflect.DeepEqual(e.Value, status) {
			found = e
			break
		}
	}
	if found != nil {
		i.inProgress.Remove(found)
	}
}

func (i *indexStatusList) writeIndexStatus(w io.Writer) {
	i.lock.Lock()
	defer i.lock.Unlock()
	for e := i.inProgress.Front(); e != nil; e = e.Next() {
		m := e.Value.(*indexStatus).String()
		fmt.Fprintf(w, "%s\n", m)
	}
}
