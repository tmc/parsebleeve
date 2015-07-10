package parsesearch

import (
	"log"
	"os"

	"github.com/tmc/parse"
)

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
	if err != nil {
		return err
	}
	err = c.CreateTriggerFunction(&parse.TriggerFunction{
		ClassName:    className,
		FunctionName: "afterSave",
		URL:          os.Getenv("URL") + "/index",
	})
	if err != nil {
		return err
	}
	return nil
}
