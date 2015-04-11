package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _assets_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x53\xcf\x4f\xdc\x3a\x10\x3e\xb3\x7f\x85\x5f\x84\x5e\x4e\xf9\x01\x7a\x3c\x55\x21\x49\x0f\x85\x4a\x1c\x5a\xaa\x42\x0f\x3d\x9a\x78\xb2\xb1\xea\xd8\x91\x3d\xbb\xb0\x45\xfc\xef\x1d\xdb\x5b\x08\xbb\x20\xda\xcb\xee\x78\x7e\x7c\xdf\xcc\x37\x93\xfa\x9f\xb3\xcb\x0f\xd7\xdf\xbf\x9c\xb3\x01\x47\xd5\x2e\x6a\x94\xa8\xa0\x5d\x9a\xd5\x54\x17\xd1\x5e\xd4\x23\x20\x67\xdd\xc0\xad\x03\x6c\xd2\x15\xf6\xd9\xbb\xf4\xb7\x5b\xf3\x11\x9a\x64\x2d\xe1\x76\x32\x16\x13\xd6\x19\x8d\xa0\xb1\x49\x6e\xa5\xc0\xa1\x11\xb0\x96\x1d\x64\xe1\x91\x50\x8d\xc3\x8d\x02\x86\x9b\x89\x8a\x10\xee\xb0\xe8\x9c\x23\x3f\xf2\x1b\x72\xdf\xb3\x9e\xaa\xb3\x9e\x8f\x52\x6d\x2a\x36\x1a\x6d\xdc\xc4\x3b\x38\x65\x0f\x0b\x14\x14\xf6\x15\x19\x57\x72\xa9\x2b\x66\xe5\x72\xc0\x53\x36\x71\x21\xa4\x5e\x56\xac\x64\x47\x27\xd3\x5d\x4c\xcd\x7d\x57\x3b\xf9\x0a\x7a\xdc\x46\x05\x28\x40\x1f\x5f\x4b\x27\x6f\xa4\x92\x48\x6c\x83\x14\x02\x74\xc8\xb0\xd5\x60\xd6\x60\xd9\x6b\xa9\xc1\x56\xa1\xad\xba\x08\x13\xb5\x8b\xfb\x7b\x26\x7b\x96\x7f\x9b\x94\xe1\x82\x3d\x50\x44\xc8\x75\xbb\x38\xa8\x7b\x63\x47\xc6\x3b\x94\x46\x37\x69\x9e\x32\x52\x6d\x30\xa2\x49\x27\xe3\x30\x65\xa0\xbb\x20\x46\x3a\xae\x14\xca\x89\x5b\x2c\x7c\x41\x26\x38\x72\xd2\xf8\xe0\xa0\x96\x7a\x5a\x61\x54\x2c\xed\xa5\x82\x34\x4a\x1e\xed\xbd\x0c\xb7\xba\x19\x25\x7a\x7f\x1d\x80\xde\xec\x60\x0f\xc1\x6b\xf6\xc8\x61\x94\x00\x9b\xb2\x49\xd1\x16\x86\xf0\x78\x74\xfe\x01\x75\x5d\x04\x0d\xea\xc1\x06\x7d\x40\x47\x61\xc2\xb2\x7d\x1a\xda\x00\x82\x43\x5b\x73\x36\x58\xe8\x9b\x24\x7f\xff\x03\x36\xcd\x68\x04\xfc\x6b\xac\xa7\xa3\xba\xfc\x8a\x2e\xeb\xd2\xbf\x72\x1f\x20\x8c\xa4\xf5\x46\x5d\xf0\x96\x8e\x74\x78\x05\x04\xe5\xf8\x32\x88\x0f\x04\x10\xc5\x1d\xd2\x95\x09\xd9\x4b\x10\x6f\xa0\x39\xf9\xf3\x65\x34\x1f\x08\x68\xde\x78\x03\xc4\xab\xfa\x22\x48\xb8\x57\x0f\xe2\x8d\x19\x08\xfd\x79\x8d\xe2\x71\x69\x83\x2c\xbf\x70\x5f\x0d\xfd\x93\x8e\x4f\xfa\x09\x9f\x4e\x3f\x7b\x46\x8c\xb2\x8e\x06\x75\x4d\xe2\xb1\x93\x59\x4b\x79\xd2\xe6\xf9\x96\x4c\x3c\x23\xdb\x6e\xca\x9b\x96\xeb\x25\xb0\x43\x60\x55\xc3\xf2\x33\x69\xcf\x35\x5a\x09\x6e\xb7\x01\xca\x3c\x84\xfc\x53\x5c\xcf\x8c\xfb\x29\x70\x4d\xaa\xe7\x1f\xe9\x2e\x38\xb2\xe4\xb8\x2c\xff\xcf\xca\xa3\xac\x3c\xa6\xaf\xb6\x2a\xff\xab\xca\x93\x64\x56\x18\xe7\xa5\xba\x0b\x47\x94\x81\x6b\x36\xe7\xf6\xf1\xca\x58\x91\xf0\x73\x14\xb4\x48\xda\xe7\xef\xd9\xb8\x61\x50\xe5\x60\x06\x4f\x9e\xb0\x4e\x2a\xb8\x8a\x6b\xfd\x5b\xc2\x1d\xbe\x5d\xba\xad\xae\x7b\x4a\x93\x23\x7e\x14\xbf\x02\x00\x00\xff\xff\x43\x28\x22\x0d\x90\x05\x00\x00")

func assets_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_assets_index_html,
		"assets/index.html",
	)
}

func assets_index_html() (*asset, error) {
	bytes, err := assets_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/index.html", size: 1424, mode: os.FileMode(438), modTime: time.Unix(1428775433, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"assets/index.html": assets_index_html,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{assets_index_html, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

