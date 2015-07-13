package parsesearch

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"text/template"

	"github.com/sheki/parsesearch/static"
)

type UI struct {
	ClassName     string
	JavascriptKey string
	AppID         string
}

func NewUI(appID, javascriptKey, className string) (*UI, error) {
	return &UI{
		AppID:         appID,
		JavascriptKey: javascriptKey,
		ClassName:     className,
	}, nil
}

func (ui *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contents, err := static.Asset(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found."))
		return
	}
	rendered, err := ui.render(bytes.NewReader(contents))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error rendering template:", err)
		return
	}
	io.Copy(w, rendered)
}

func (ui *UI) render(r io.Reader) (io.Reader, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("").Parse(string(content))
	if err != nil {
		return nil, err
	}
	output := new(bytes.Buffer)
	if err := tmpl.Execute(output, ui); err != nil {
		return nil, err
	}
	return output, nil
}
