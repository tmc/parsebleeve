package parsesearch

import (
	"bytes"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/sheki/parsesearch/static"
)

type UI struct {
	className     string
	javascriptKey string
	appID         string

	http.Handler
}

func NewUI(appID, javascriptKey, className string) (*UI, error) {
	return &UI{
		appID:         appID,
		javascriptKey: javascriptKey,
		className:     className,
		Handler: http.FileServer(&assetfs.AssetFS{
			Asset: static.Asset, AssetDir: static.AssetDir}),
	}, nil
}

func (ui *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bwr := &bufferedResponseWriter{
		rw:     w,
		Buffer: new(bytes.Buffer),
	}
	ui.Handler.ServeHTTP(bwr, r)
	bwr.WriteString("MODIFIED")
	io.Copy(w, bwr)
}

type bufferedResponseWriter struct {
	*bytes.Buffer
	rw http.ResponseWriter
}

func (w *bufferedResponseWriter) Write(b []byte) (int, error) {
	spew.Dump("Writing", b)
	return w.Buffer.Write(b)
}

func (w *bufferedResponseWriter) Header() Header {
}
