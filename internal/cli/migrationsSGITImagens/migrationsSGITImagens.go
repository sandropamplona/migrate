package migrationsSGITImagens

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

var __000001_se_teste_imagem_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000001_se_teste_imagem_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_se_teste_imagem_down_sql,
		"000001_SE-TESTE_IMAGEM.down.sql",
	)
}

var __000001_se_teste_imagem_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xd1\x09\x80\x40\x0c\x03\xd0\x6f\x6f\x8a\x7c\xea\x58\xe7\x19\xa5\x10\x5a\xa8\x71\x7f\xdf\x6a\x4e\x13\x9e\xa7\x08\xad\xbb\x5c\xe6\x6b\xee\x63\x8b\x0b\x91\xe6\xc3\x46\x96\x91\x9f\x74\x8c\x3f\x00\x00\xff\xff\x7e\xa3\xf2\x98\x30\x00\x00\x00")

func _000001_se_teste_imagem_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_se_teste_imagem_up_sql,
		"000001_SE-TESTE_IMAGEM.up.sql",
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
	"000001_SE-TESTE_IMAGEM.down.sql": _000001_se_teste_imagem_down_sql,
	"000001_SE-TESTE_IMAGEM.up.sql": _000001_se_teste_imagem_up_sql,
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
	"000001_SE-TESTE_IMAGEM.down.sql": &_bintree_t{_000001_se_teste_imagem_down_sql, map[string]*_bintree_t{
	}},
	"000001_SE-TESTE_IMAGEM.up.sql": &_bintree_t{_000001_se_teste_imagem_up_sql, map[string]*_bintree_t{
	}},
}}
