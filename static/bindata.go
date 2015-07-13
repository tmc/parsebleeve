package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb2, 0x51,
		0x4c, 0xc9, 0x4f, 0x2e, 0xa9, 0x2c, 0x48, 0x55, 0xc8, 0x28, 0xc9, 0xcd,
		0xb1, 0xe3, 0xb2, 0xc9, 0x48, 0x4d, 0x4c, 0xb1, 0xb3, 0xd1, 0x07, 0x53,
		0x5c, 0x36, 0x49, 0xf9, 0x29, 0x95, 0x76, 0x5c, 0x0a, 0x0a, 0x36, 0x05,
		0x45, 0xa9, 0x20, 0x5a, 0x41, 0xc1, 0xb1, 0xa0, 0x40, 0xc1, 0xd3, 0xc5,
		0x4a, 0xa1, 0xba, 0x5a, 0x0f, 0xc8, 0xf4, 0x74, 0xa9, 0xad, 0x05, 0x0b,
		0x7b, 0x05, 0x2b, 0x64, 0xa7, 0x56, 0x82, 0x85, 0xbd, 0x12, 0xcb, 0x12,
		0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xbc, 0x53, 0x2b, 0xa1, 0xd2, 0xce,
		0x39, 0x89, 0xc5, 0xc5, 0x60, 0x59, 0x30, 0xcb, 0x2f, 0x31, 0x37, 0x15,
		0x2c, 0x63, 0xa3, 0x0f, 0x36, 0xd8, 0x46, 0x1f, 0x62, 0x11, 0x20, 0x00,
		0x00, 0xff, 0xff, 0xac, 0x34, 0x4d, 0xd1, 0x8e, 0x00, 0x00, 0x00,
	},
		"index.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"index.html": index_html,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"index.html": &_bintree_t{index_html, map[string]*_bintree_t{
	}},
}}
