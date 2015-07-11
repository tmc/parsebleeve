package parsesearch

import (
	"log"
	"os"
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
	err = squelchAlreadyExists(c.CreateHookFunction(&parse.HookFunction{
		FunctionName: "search",
		URL:          urlPrefix + "/search",
	}))
	if err != nil {
		return err
	}
	err = squelchAlreadyExists(c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:   className,
		TriggerName: "afterSave",
		URL:         urlPrefix + "/index",
	}))
	if err != nil {
		return err
	}
	err = squelchAlreadyExists(c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:   className,
		TriggerName: "afterDelete",
		URL:         urlPrefix + "/unindex",
	}))
	return err
}
