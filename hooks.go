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

	err = squelchAlreadyExists(c.CreateHookFunction(&parse.HookFunction{
		FunctionName: prefix + "search",
		URL:          os.Getenv("URL") + "/search",
	}))
	if err != nil {
		return err
	}
	err = squelchAlreadyExists(c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:   className,
		TriggerName: "afterSave",
		URL:         os.Getenv("URL") + "/index",
	}))
	if err != nil {
		return err
	}
	err = squelchAlreadyExists(c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:   className,
		TriggerName: "afterDelete",
		URL:         os.Getenv("URL") + "/unindex",
	}))
	if err != nil {
		return err
	}
	return nil
}
