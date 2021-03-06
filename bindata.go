// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/index.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xdb\x30\x0c\x3d\xd7\x5f\xc1\x19\xc1\x7c\x8a\x9c\x16\xeb\x30\xb8\xb2\x77\x58\x3b\xa0\x87\xad\xc3\xda\x1d\x76\x54\x2d\x3a\x16\x26\x4b\x86\xc4\xa4\xcd\x8c\xfc\xfb\x20\x39\x6b\xdd\xa4\x45\xb7\xdb\xa3\x48\xbe\x47\x3e\xda\xfc\xcd\xf9\xd5\xa7\x9b\x9f\xdf\x2e\xa0\xa5\x4e\x57\x09\x27\x45\x1a\xab\xa5\x5d\xf5\x3c\x1f\x71\xc2\x3b\x24\x01\x75\x2b\x9c\x47\x2a\xb3\x15\x35\xf3\x0f\xd9\xdf\x67\x23\x3a\x2c\xd3\xb5\xc2\xbb\xde\x3a\x4a\xa1\xb6\x86\xd0\x50\x99\xde\x29\x49\x6d\x29\x71\xad\x6a\x9c\xc7\x20\xad\x12\xee\x69\xa3\x11\x68\xd3\x63\x99\x12\xde\x53\x5e\x7b\x9f\x56\x09\x89\x5b\x8d\x30\x40\x63\x0d\xcd\x1b\xd1\x29\xbd\x29\xa0\xb3\xc6\xfa\x5e\xd4\x78\x06\xdb\x84\x24\x0c\x10\x3a\xe6\x42\xab\xa5\x29\xc0\xa9\x65\x4b\x67\xd0\x0b\x29\x95\x59\x16\xb0\x80\xe3\xd3\xfe\x7e\x2c\x65\x61\xaa\xbd\x7a\x8d\x0d\xed\xb2\x12\x35\x52\xc8\xaf\x95\x57\xb7\x4a\x2b\xda\x14\xd0\x2a\x29\xd1\xc4\x0a\x57\xb4\x76\x8d\x0e\x5e\x2a\x8d\x58\xc7\xb1\x78\x1e\x37\xaa\x92\x61\x00\xd5\x00\xfb\xd1\x6b\x2b\x24\x6c\xb7\x09\x97\x6a\x5d\x25\x47\xbc\xb1\xae\x03\x51\x93\xb2\xa6\xcc\x58\x06\x1d\x52\x6b\x65\x99\xf5\xd6\x53\x06\x68\xea\x68\x46\xd6\xad\x34\xa9\x5e\x38\xca\x43\xc3\x5c\x0a\x12\x59\x95\x1c\x1d\x71\x65\xfa\x15\x8d\x8e\x65\x8d\xd2\x98\x8d\x96\x8f\xf8\xa0\xc2\xaf\x6e\x3b\x45\xe1\x9d\x47\xa2\x57\x27\x38\x60\x08\x9e\x3d\x68\x58\x2d\xd1\x65\xd0\x6b\x51\x63\x1b\x83\x87\xc7\x7f\x90\xe6\x79\xf4\x80\xb7\x2e\xfa\x83\x66\x34\x26\x1e\x3b\x94\x91\x8b\x24\xd4\x56\x5c\x40\xeb\xb0\x29\x53\xf6\xf1\x17\x6e\xca\xce\x4a\x7c\x6b\x5d\x90\x1b\x06\x60\xd7\xd6\xd1\x55\x88\x58\x48\xc0\x76\x9b\x56\x01\xf0\x5c\x54\x3c\xa7\xf6\x05\x12\x52\xdd\xf3\x24\x21\x11\x49\xb4\xf0\x04\x9d\x95\xaa\x51\x28\x5f\x61\xf3\xea\xf7\xf3\x6c\x21\x11\xd9\x02\x78\x85\x24\xb8\xfa\x2c\x49\xfc\x5e\x03\x49\x00\x13\x12\x9e\x47\x8f\xc6\x8f\xcb\x58\x02\x76\xe9\xbf\x5b\x4b\xc1\xc7\x47\xff\x64\x28\x97\xd5\x21\x18\xb3\x50\x6b\xe1\x7d\x99\x06\xee\x74\x32\x12\x4b\x2b\xc6\x76\x62\xf2\x89\xd8\xee\x52\x01\x3a\x61\x96\x08\x33\x84\xa2\x04\x76\xae\xdc\x85\x21\xa7\xd0\xef\x0f\x30\x0c\x30\x43\xf6\x65\x3c\xcf\x44\xfb\x31\x71\xa3\x3a\x64\x9f\xad\xeb\x04\x41\x7a\xb2\x58\xbc\x9f\x2f\x8e\xe7\x8b\x13\x38\x3e\x2d\x16\xef\x8a\xc5\x69\x3a\x69\x1c\xf7\x9d\x21\xbb\xf4\xe7\xca\x45\xad\xc9\x9e\xbb\xe0\x85\xb5\x46\xc1\xaf\xa3\xa1\x79\x5a\x3d\x8d\x27\xeb\xc6\x45\xb5\xc7\x09\xfd\x30\x40\x3c\xe7\x0c\xd9\xf5\x78\xd6\xff\x15\xdc\xd3\xdb\x97\xdb\xf9\x7a\xe0\x34\xcf\x77\x3f\xc5\x9f\x00\x00\x00\xff\xff\x43\x28\x22\x0d\x90\x05\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 1424, mode: os.FileMode(436), modTime: time.Unix(1528887806, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"assets/index.html": assetsIndexHtml,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

