// Code generated by go-bindata.
// sources:
// assets/col16.txt
// assets/col256.txt
// assets/defs.txt
// assets/help.txt
// DO NOT EDIT!

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

var _assetsCol16Txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8d\x5d\x0a\x83\x30\x10\x84\xdf\xe7\x14\x7b\x81\x82\xbf\x2b\xec\x6d\x56\x5d\xa2\x98\xa6\x10\x52\x24\xb7\x2f\xb5\x5b\x1f\xbf\x99\x61\xbe\x39\xea\x72\x08\x35\xc8\xb6\x0a\xb5\x08\xd9\x2c\x09\x75\xa8\x16\xe3\xeb\x14\xea\x31\xc7\xb7\x09\x0d\x78\x6a\xb0\x54\x54\x68\xc4\x52\x35\x09\x31\xe2\x1e\xb6\xf2\x08\x59\xab\xd0\x84\x55\xf3\xe1\xc0\x8d\x77\xd7\x2d\xb7\xf7\xf2\x7a\xe7\xce\xf9\x2f\xe1\xde\x83\x9f\x8b\x07\xc7\x5b\xc9\xa3\x27\x6e\x66\x9c\xdb\x5e\xbe\xd3\xe9\x13\x00\x00\xff\xff\xd8\xbf\x18\xcb\xc2\x00\x00\x00")

func assetsCol16TxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsCol16Txt,
		"assets/col16.txt",
	)
}

func assetsCol16Txt() (*asset, error) {
	bytes, err := assetsCol16TxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/col16.txt", size: 194, mode: os.FileMode(436), modTime: time.Unix(1596309230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCol256Txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x95\x51\x6f\xe3\x36\x0c\xc7\xdf\xf9\x69\x2a\xcb\x6d\xd2\xc7\x35\x87\xf5\x86\x5d\xd6\xa2\x09\x76\xd8\xa3\x50\x73\xb1\x16\x45\x6a\x15\xbb\x81\xbf\xfd\x20\x91\x92\x25\x37\xe8\xbd\xd8\x64\x42\xfe\x2d\xff\x7f\xa2\xfc\x60\xd4\xeb\x11\xb6\xca\x3b\x67\xe1\xd1\x23\x5a\x78\x32\xfa\x03\xe1\x2f\xf5\x31\xc1\xf3\xe8\xdf\x0c\xc2\x1e\x95\x81\x9d\x36\x1f\xe8\x43\xcd\x04\x2f\xd8\xc1\x0f\x7d\x42\xf8\x07\x8d\x71\x17\x78\x30\x23\xc2\xef\xe3\x6b\x7f\xd6\x0a\x7e\x7b\x1f\x15\xfc\xec\xf5\x80\xb1\xf8\x26\x4a\xc5\x8a\x6f\xca\x1f\x63\x10\x2e\xb2\xb8\x8a\xf8\x17\x3d\xfe\x1b\xe2\xdb\xee\x18\x1b\xda\xaf\x12\xd7\x1d\xd0\x93\xc4\x1c\x37\xf4\x0e\x2d\xec\xde\xbc\xb6\x07\x4e\xf6\xa3\x7f\x1f\x9d\x3e\xd7\x1a\x72\x91\x64\x11\x41\x22\xb2\x14\x91\x71\x85\x9b\x49\x59\xf8\xa1\x0f\xfd\xb0\x43\xf5\x69\xb9\x4d\x99\x5c\x17\x29\x92\x06\x82\x1a\xe9\xe6\x05\xce\x4b\xe5\x37\x11\x75\x4b\x91\x08\xd8\x62\xa7\xc7\x53\xf1\x53\x54\x24\x5d\x72\x34\x70\x0a\x6b\x7a\xd6\xf6\xd8\x32\xcd\xe5\x9d\x08\xfc\xad\x9d\xc1\x01\x9e\xbc\xb2\x07\x6c\x23\x38\xb9\xe2\x27\xa4\xfa\x9d\x51\x03\x92\x59\x45\xf8\xe2\x26\x65\xe8\x8d\x37\xbd\xf2\x83\xc7\x31\x1a\xad\xfc\x31\x99\xd4\xc2\xb3\x32\x58\x50\xd8\x0d\x88\xb1\x67\x8e\x24\x6c\x9c\xb7\xff\x1a\x77\x21\x08\x85\x98\x5c\x88\x6d\x54\x87\x03\xd5\xe4\x28\x63\xcc\x82\xa2\x52\x08\x0b\x48\x0c\x58\x48\xc6\x9d\x7a\x52\x5e\x5b\x94\xfc\xaa\x33\x88\xab\x3a\x4d\x6e\x9e\x23\x51\x44\xb3\x20\xf9\x1f\x6d\x7a\xf4\x6a\x6a\xae\xe0\x08\xbf\x6c\xd5\x01\xed\xa0\x3e\xc5\x8c\x83\xe7\x2f\x51\x89\x3b\x8f\x59\x9a\xf1\xd4\x56\x7c\xe4\x22\xcb\x88\x04\x4f\x69\x0b\x3f\x7b\x54\x03\xc1\xbd\x95\xbc\x8f\x69\x85\x38\x55\xdd\xc5\x7f\xd1\xdd\x24\x10\x96\x16\x8f\x87\x62\x26\xf2\x24\x50\x4f\xe2\x50\x67\x79\x42\x4a\x2f\x97\x6a\x05\xa3\x52\x58\xd6\x56\x66\xb1\x12\x8c\xa0\xe7\x15\x0b\xa1\x30\x4b\x7e\xc1\x46\x84\xe3\x4c\x16\x60\xc8\x09\x42\x10\x90\x31\x18\x79\x85\x4c\x7c\x85\x48\x47\xc2\x1f\xb6\xd3\xca\x86\x86\xef\x2e\x52\x4a\x44\x9e\xfc\x6b\xaf\xbb\x3a\xab\xec\x26\x2b\x1e\x9d\xe9\xd0\x7a\xd7\xb1\x77\xca\x9c\x9c\x0d\xf3\x75\x9e\x1e\xbc\xbb\xc4\xe3\x79\xba\x93\x8b\xce\x32\x13\x10\x34\x68\x9d\x7f\xf6\xea\xa8\xc3\xd1\xab\xfe\x73\xf1\x30\x96\xd4\x7f\xcf\xea\xf3\xdc\xd5\x39\xb3\x96\x5f\xb2\xae\xb3\x86\x24\xe8\x38\x2b\xb9\xf3\x19\xc8\x1f\x89\x5a\xb0\x29\xd9\xd4\x6a\x65\x26\xea\xa3\x63\xc1\xaa\x8c\x32\xa5\x65\xd0\xfc\x8a\x12\x07\x0d\x54\x68\x28\x11\x90\x1a\x2b\x28\x79\x12\x25\xf0\xd5\x8c\x27\x09\xbc\x39\x88\x02\x6d\xc3\x04\x55\xc2\x5e\x59\xd8\xea\xf3\x30\xbd\xb8\x70\x22\xed\x7b\x7d\x1e\xc2\xa0\x86\xd6\x26\xdb\x1e\xb9\x2d\x9b\xd9\xe1\x54\x13\x40\xae\xdb\x05\xb8\x34\xe6\xcb\xb9\x12\x57\xf3\xd9\xdf\xef\xce\xe2\xd4\xe1\xa5\xa0\x18\x5d\x16\xd9\xdb\x26\x47\xa2\x88\xb2\xbb\x1c\x24\xa7\x62\x6b\x76\xb9\x0a\xd9\xe8\x7c\xaf\x8d\x9e\x29\x01\xd9\xcc\x43\xbd\x71\x5e\x99\xb8\x0d\xf2\x48\x0a\x46\x95\x90\xa5\x67\x0b\xd8\x29\xdb\xf1\xb8\x14\xc0\xc4\x0c\x4c\x00\x5f\xcd\x78\xa2\x79\x11\xd7\xcd\x2e\xf2\x62\x8a\xc4\x8c\x50\x24\x84\xc9\xf9\xa5\x90\x20\x98\x82\xce\x5d\x11\x3f\x70\x67\x6d\x8e\x71\x2c\x26\x71\x73\x43\x9f\xd9\x78\x5d\xd1\x6f\xfc\xd7\x2d\xdd\xee\xe3\xad\xa1\x8a\x86\x4a\x24\x77\x51\x89\xa4\x92\xb6\xa1\xdb\x1d\x9d\xed\x54\x72\xcb\x27\xfd\x9a\x06\x9f\x4a\xee\xa8\x64\x45\x25\x2b\x2a\x59\x51\xc9\x9a\x4a\xd6\x24\xbd\x26\xe9\x7b\xf9\x7f\x00\x00\x00\xff\xff\xcf\x8b\x8f\x21\x23\x0a\x00\x00")

func assetsCol256TxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsCol256Txt,
		"assets/col256.txt",
	)
}

func assetsCol256Txt() (*asset, error) {
	bytes, err := assetsCol256TxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/col256.txt", size: 2595, mode: os.FileMode(436), modTime: time.Unix(1596280880, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDefsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc6\xc1\x09\xc2\x50\x10\x04\xd0\xfb\x54\x31\x25\xc4\xa0\xa0\xd3\x8d\x61\x47\xfc\xb8\x6e\x60\x93\x78\xb0\x7a\xf1\xf6\xbe\xee\x55\x9c\xb0\xac\x19\xe2\x09\x31\xde\xe2\x8c\xa3\xc2\x9d\xa3\x2c\x9e\xb1\xe4\xa8\x97\x78\x41\xfb\xe3\xde\xfc\xe7\x73\x44\xb8\xc4\x2b\xc2\x8f\xfb\x91\xbb\x78\x43\x7b\xf3\x2e\xce\x13\x7e\x01\x00\x00\xff\xff\x35\x3b\x6c\x07\x57\x00\x00\x00")

func assetsDefsTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsDefsTxt,
		"assets/defs.txt",
	)
}

func assetsDefsTxt() (*asset, error) {
	bytes, err := assetsDefsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/defs.txt", size: 87, mode: os.FileMode(436), modTime: time.Unix(1596311296, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsHelpTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x4d\x6f\xd3\x4c\x10\xbe\xef\xaf\x98\xc3\x7b\xec\x8b\x00\x89\x1e\x7c\xa2\x1f\x71\x55\x89\x24\x55\x5a\x09\x50\x53\xa1\xb1\x3d\xeb\x5d\xb1\xde\xb1\x66\xc7\x24\xe1\xd7\xa3\xb5\x49\x9a\x94\x0a\x8e\x9e\x79\xbe\xe6\x59\x5f\x62\x72\x70\x27\xdc\xf5\x0a\x2d\x45\x12\x54\x16\x93\x76\x51\x71\x5b\x40\x85\xc9\xfd\xdf\x4f\xdb\xc7\xab\xe5\xa2\xbc\xbd\xf9\x56\xde\x7e\x9a\x3d\x19\xf3\xd9\xab\xe3\x41\x01\xa5\x1d\x3a\x8a\x9a\x20\x39\xde\x24\x50\xe7\x13\x38\x0a\xbd\x39\xc2\x83\xcf\x0b\x02\xeb\x03\x45\xec\x08\x36\x5e\xdd\x38\xa9\x39\x5a\xdf\x0e\x82\xea\x39\x82\x32\x60\xdf\x87\x9d\x31\x57\x27\xf3\xd2\x07\x32\x33\xac\x1d\x04\x1f\x29\xab\x21\xd4\xdc\x75\x18\x9b\x67\x29\xcb\x21\xf0\xc6\xc7\x16\x7e\xc7\x37\x57\xf3\xeb\xe2\x62\x75\x73\x9f\xe5\x46\x70\x32\x8b\xe5\x1d\x00\xc0\x82\x61\xd9\xd3\xa4\xfe\x06\xbe\xf2\x00\x35\xc6\x1c\x26\xf9\x86\x04\x70\x6f\x40\x51\xcd\xc3\x97\x87\x4c\xb9\x68\x9a\xe9\x08\xa5\xad\x1e\x99\x4d\x57\x04\x8e\x30\x44\xf5\x61\xfc\xa6\xd8\x00\xdb\x31\xac\x29\x57\xf3\x03\x1d\xc1\xb2\x74\xa8\xfb\xf0\x29\x3b\xc5\xa1\xab\x48\xd2\x19\x34\x64\x7d\xf4\x39\x52\x02\x16\x28\x57\xf3\x9c\x1e\x36\xce\xd7\x6e\xcc\x57\xd1\x1e\x9c\xf7\x47\x70\x73\xb9\xba\x3e\x32\xa9\x84\xb0\xa9\x65\xe8\x7a\xa8\x1d\x0a\xd6\x4a\x32\xf5\x84\x90\x7a\xac\x09\xbc\x85\xc8\x30\x5b\x5c\x1f\xde\x2f\x97\xda\x0b\xff\xf0\x0d\x35\xc6\x94\x53\xcc\x43\x6d\x89\x34\xeb\xdf\x93\x4e\x1d\xec\x91\xcf\x74\x4c\xe0\x93\x91\x34\xe2\x56\x94\xf6\xc8\xe9\x60\xcd\x4d\x59\x96\xd7\xc9\xc6\xb6\x70\x22\x6f\x59\xa8\x15\x1e\x62\x03\x43\xca\x54\x84\x77\xe7\x63\xcb\x02\x3d\x06\x52\x25\x53\xbd\xe0\x54\x58\x7f\xff\x17\xc7\xb6\xef\x3f\x9c\xff\xdd\x27\x03\x5e\x1a\x9d\x92\x5e\x31\xfa\x93\x64\x66\x5b\xec\xfa\x30\x3e\x7f\x51\x71\x68\xce\x6c\x5b\xd4\x3b\x8c\xf9\x77\x2a\x1e\xd7\xfa\x04\x27\xab\x56\x88\xa6\xdd\x7a\xf8\xb8\x76\xe3\x4e\x92\x16\x6f\xc7\x59\x71\x82\xad\xc2\x40\x13\x74\x73\xc0\xfd\x24\xe1\x71\xf6\x1f\x98\x5f\x01\x00\x00\xff\xff\x3d\x19\x31\x7a\xd4\x03\x00\x00")

func assetsHelpTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsHelpTxt,
		"assets/help.txt",
	)
}

func assetsHelpTxt() (*asset, error) {
	bytes, err := assetsHelpTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/help.txt", size: 980, mode: os.FileMode(436), modTime: time.Unix(1596365644, 0)}
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
	"assets/col16.txt": assetsCol16Txt,
	"assets/col256.txt": assetsCol256Txt,
	"assets/defs.txt": assetsDefsTxt,
	"assets/help.txt": assetsHelpTxt,
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
		"col16.txt": &bintree{assetsCol16Txt, map[string]*bintree{}},
		"col256.txt": &bintree{assetsCol256Txt, map[string]*bintree{}},
		"defs.txt": &bintree{assetsDefsTxt, map[string]*bintree{}},
		"help.txt": &bintree{assetsHelpTxt, map[string]*bintree{}},
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

