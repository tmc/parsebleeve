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
func (i *Indexer) RegisterHooks(prefix string, className string) error {
	c, err := parse.NewClient(i.appID, "")
	c.TraceOn(log.New(os.Stderr, "[parse api] ", log.LstdFlags))
	if err != nil {
		return err
	}
	c = c.WithMasterKey(i.masterKey)

	err = c.CreateHookFunction(&parse.HookFunction{
		FunctionName: prefix + "search",
		URL:          os.Getenv("URL") + "/search",
	})
	err = squelchAlreadyExists(err)
	if err != nil {
		return err
	}
	err = c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:   className,
		TriggerName: "afterSave",
		URL:         os.Getenv("URL") + "/index",
	})
	err = squelchAlreadyExists(err)
	if err != nil {
		return err
	}
	return nil
}
