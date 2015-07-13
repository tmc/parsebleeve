package parsesearch

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"html/template"

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
	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}
	contents, err := static.Asset(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found."))
		return
	}
	if err := ui.render(bytes.NewReader(contents), w); err != nil {
		// assume the file is binary and return it verbatim
		io.Copy(w, bytes.NewReader(contents))
	}
}

func (ui *UI) render(r io.Reader, w io.Writer) error {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	tmpl, err := template.New("").Parse(string(content))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, ui)
}
