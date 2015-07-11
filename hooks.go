package parsesearch

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/tmc/parse"
)

func squelchAlreadyExists(err error) error {
	if err, ok := err.(*parse.Error); ok {
		if strings.Contains(err.Message, "already exists") {
			return nil
		}
	}
	return err
}

// RegisterHooks auto-registers the search service with a Parse Application.
func (i *Indexer) RegisterHooks(className string) error {
	c, err := parse.NewClient(i.appID, "")
	c.TraceOn(log.New(os.Stderr, "[parse api] ", log.LstdFlags))
	if err != nil {
		return err
	}
	c = c.WithMasterKey(i.masterKey)
	urlPrefix := os.Getenv("URL")
	if urlPrefix == "" {
		return fmt.Errorf("Skipping registering Webhooks as the 'URL' environment variable is empty.")
	}
	err = squelchAlreadyExists(c.CreateHookFunction(&parse.HookFunction{
		FunctionName: "search",
		URL:          path.Join(urlPrefix, "search"),
	}))
	if err != nil {
		return err
	}
	err = squelchAlreadyExists(c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:   className,
		TriggerName: "afterSave",
		URL:         path.Join(urlPrefix, "index"),
	}))
	if err != nil {
		return err
	}
	err = squelchAlreadyExists(c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:   className,
		TriggerName: "afterDelete",
		URL:         path.Join(urlPrefix, "unindex"),
	}))
	return err
}
