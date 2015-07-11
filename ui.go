package parsesearch

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/alecthomas/template"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/sheki/parsesearch/static"
)

type UI struct {
	ClassName     string
	JavascriptKey string
	AppID         string

	http.Handler
}

func NewUI(appID, javascriptKey, className string) (*UI, error) {
	return &UI{
		AppID:         appID,
		JavascriptKey: javascriptKey,
		ClassName:     className,
		Handler: http.FileServer(&assetfs.AssetFS{
			Asset: static.Asset, AssetDir: static.AssetDir}),
	}, nil
}

func (ui *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	ui.Handler.ServeHTTP(rec, r)
	rendered, err := ui.render(rec.Body)
	if err != nil {
		log.Println("error rendering template:", err)
		io.Copy(w, rec.Body)
		return
	}
	w.WriteHeader(rec.Code)
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
