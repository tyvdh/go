// Code generated by go-bindata. DO NOT EDIT.
// sources:
// graphiql.html
// schema.gql

package static

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataGraphiqlhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4f\x6f\x13\x3f\x10\x3d\x6f\x3e\x85\x7f\x96\x7e\xd2\x46\x2a\x76" +
		"\x52\x24\x0e\x9b\x4d\x0e\xd0\xaa\x02\x15\x4a\x81\x0b\x47\xd7\x9e\x5d\x3b\x78\xed\xed\xd8\x9b\x36\xaa\xf2\xdd\x91" +
		"\xf7\x4f\x28\x7f\x2a\x21\x04\x17\xaf\x3d\x7e\xf3\xde\xf3\xcc\x68\xcb\xff\xce\xae\x5e\x7d\xfa\xfc\xfe\x9c\xe8\xd8" +
		"\xd8\xcd\xac\x1c\x3e\x59\xa9\x41\xa8\xcd\x2c\xcb\x4a\x6b\xdc\x17\x82\x60\xd7\x34\xc4\xbd\x85\xa0\x01\x22\x25\x1a" +
		"\xa1\x5a\x53\x1d\x63\x1b\x0a\xce\xa5\x72\xdb\xc0\xa4\xf5\x9d\xaa\xac\x40\x60\xd2\x37\x5c\x6c\xc5\x3d\xb7\xe6\x26" +
		"\xf0\x1a\x45\xab\xcd\xad\xe5\x0b\xb6\x5c\xb2\xe5\xf2\x18\x60\x32\x04\xca\x7b\x99\x20\xd1\xb4\x91\x04\x94\xbf\x4d" +
		"\x5b\x41\x94\x9a\x9f\xb2\x05\x7b\x3e\xec\x59\x63\x1c\xdb\x06\xba\x29\xf9\x40\xf7\xa7\xcc\x08\x42\x46\xbe\x7c\xc1" +
		"\x4e\xd9\x82\x77\x8d\x1a\x02\xac\x45\xaf\x3a\x19\x8d\x77\x7f\x57\xe9\x99\xf2\xcd\x4f\x6a\x29\xf8\x2f\x14\x9f\x6e" +
		"\xc6\x2f\x14\x4a\x3e\xce\x41\x79\xe3\xd5\x9e\xf4\x13\xb0\xa6\x77\x46\x45\x5d\x90\xe5\x62\xf1\xff\x8a\x68\x30\xb5" +
		"\x8e\xd3\xa9\x11\x58\x1b\x57\x90\xc5\x8a\xf8\x1d\x60\x65\xfd\x5d\x41\xb4\x51\x0a\xdc\x8a\xf6\x96\x95\xd9\x11\xa3" +
		"\xd6\x74\x92\xa5\x13\xeb\x23\xa2\x9d\x5e\xd1\xcd\xa5\x17\xca\xb8\x9a\x31\x56\x72\x65\x76\x8f\xde\x9b\xb6\x59\xd5" +
		"\xb9\xbe\x30\xa4\x6f\xfd\xc5\xf5\x65\xde\x0a\x14\x4d\x98\x93\x87\x74\x9d\x21\xc4\x0e\xc7\xdb\x9c\x0e\xaf\xbc\xb5" +
		"\xf4\x64\xbc\xce\x1a\x88\xda\xab\x82\xd0\xd6\x87\x48\x4f\x86\x60\x7a\x65\x41\xde\x7c\xbc\x7a\xc7\x42\x44\xe3\x6a" +
		"\x53\xed\x27\xde\x11\x22\x11\x14\xb8\x68\x84\x0d\x05\xa1\xc6\x49\xdb\x29\x18\xf3\x0f\x73\x16\x35\xb8\xfc\xe8\x2d" +
		"\x47\x08\xed\xe4\x68\xb2\x94\x62\x2c\xc2\x7d\xcc\xe7\xab\x27\xd2\x92\x8f\x63\x5a\xc4\xfd\xb4\x9d\x28\x7a\x87\xad" +
		"\xc0\x00\x03\x74\xe0\xc9\x0e\x44\x8a\x28\x35\xc9\x01\xd1\xe3\xfc\xc7\xac\x04\x9d\x90\xa3\x70\x7f\x3c\xcc\xd2\xfa" +
		"\x21\x4d\xdd\xd9\xd5\x5b\x86\xe0\x14\x60\xde\x23\xfa\x20\x93\x08\x22\xc2\xb9\x85\x06\x5c\xcc\x2f\xfa\xce\x5d\x5f" +
		"\x9e\x90\x87\xbe\xba\x80\xc5\xb1\x09\x87\xb1\x4c\xca\xcb\x2e\x81\x59\x0d\x71\xcc\x7b\xb9\x7f\xad\xf2\x6f\x6d\x9f" +
		"\x27\x5c\x5a\xbe\x1b\xb7\x64\x71\x33\x2b\xf9\xf0\x1b\xfa\x1a\x00\x00\xff\xff\xdb\x8e\x2c\x18\x9e\x04\x00\x00")

func bindataGraphiqlhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataGraphiqlhtml,
		"graphiql.html",
	)
}

func bindataGraphiqlhtml() (*asset, error) {
	bytes, err := bindataGraphiqlhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "graphiql.html",
		size:        1182,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1556028957, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataSchemagql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x82\xbd\x24\x17\x1f\x8a\x9e" +
		"\x84\xb6\x80\x93\xb4\x68\xd0\xec\xa2\x5d\x67\x8b\x02\x41\x51\x8c\xc5\xb1\x44\x84\x1f\xda\xe1\xd0\x59\xa3\xc8\x7f" +
		"\x2f\x48\xc9\x59\xca\x4a\xdd\x73\x4f\x12\xdf\xcc\x1b\x72\x1e\xdf\x30\xb4\x3d\x59\x84\xbf\xeb\xea\x73\x24\x3e\x34" +
		"\x50\xfd\x96\xbe\xf5\x4b\x5d\xcb\x61\x20\xc8\xab\x14\x7e\x07\x4c\xc2\x9a\xf6\x04\x68\x0c\xec\xd1\x68\x85\x42\x0a" +
		"\x30\x04\x92\x00\xde\x81\xf4\x04\x1b\x21\x63\x90\xc1\x91\x3c\x7b\x7e\x5a\xd5\xd5\x18\x6f\xe0\x71\x9d\x7e\x2e\xfe" +
		"\xbc\xa8\xcf\x14\xd3\x21\x44\xe2\x33\xd5\xa6\x84\x06\x1e\xef\xf2\xdf\xa2\x9e\x30\x2a\x82\x20\x28\x01\x76\xec\x6d" +
		"\xae\x63\x30\x08\x7c\xe7\xa2\xfd\xd9\x47\x0e\xeb\xce\xff\x00\x7d\xfa\x4b\xcc\x4b\x45\x3b\x8c\x46\xe0\x7b\xf8\xe6" +
		"\xdb\x11\xbe\x5a\x81\x1f\x44\x7b\x87\xc6\x1c\x60\x60\xbf\xd7\x8a\xa0\xf5\xd1\x09\x31\xa0\x53\x89\xb7\xc5\x40\x63" +
		"\xf3\xa0\xdd\xce\xc3\xce\x33\xec\xb4\x11\x62\xed\xba\x55\x5d\x59\xe4\x27\x92\x70\x59\x57\x55\x4a\xcd\xdd\xdf\x78" +
		"\x45\x0d\x6c\x24\xa5\x94\xf8\xd8\x4b\x11\x99\xf6\x7a\x8b\x54\x86\x16\xbc\xa2\xc5\x06\xee\x9c\xd4\xd5\x55\x03\x8f" +
		"\xef\xf3\x51\x16\xca\x77\x1d\x53\x97\x65\x9f\x89\xe6\xf9\x5f\x34\x4b\xec\xac\xcf\x9b\xf2\x24\x8e\x43\x4b\xe0\x77" +
		"\xf9\x7f\xac\x39\xa0\x66\xb8\xa4\x55\x52\xe4\x1d\xfc\x71\xff\xfe\xaf\xeb\x87\x9b\xab\xb9\x58\xc0\x14\xa2\x91\xb0" +
		"\xaa\x2b\xd1\xed\x13\x71\xd2\x2c\x11\x3f\xa0\xa5\xff\x6c\x6e\xfd\xda\xc6\x6b\x9b\x2f\x75\x1d\x5a\x4c\xc6\xb9\xd6" +
		"\x5d\x4a\x9c\x56\x0f\xda\xd2\xe4\xeb\x2c\x5f\xf2\x75\x5b\xa8\x7b\x71\xf4\xd7\xba\xcd\x2a\x17\x78\x22\x15\x4b\x17" +
		"\xed\x94\x13\xf2\x51\x2e\xea\x0a\xa3\xf4\x1f\xe9\x73\xd4\x4c\xaa\x81\x6b\xef\x0d\xa1\x7b\xc5\xf7\xbe\xc5\xad\xa1" +
		"\x59\xc0\x8e\x7b\xfc\x64\x3c\xe6\x02\xe3\x65\x3b\x61\x6f\x0c\xa9\xeb\xc3\xad\xb7\xa8\xdd\x8c\xe2\xda\xde\x2f\x5d" +
		"\x31\x8f\x3c\xcc\x8f\xaa\x43\x46\xd7\x39\x61\x7e\x34\xa5\xc3\x60\xf0\x70\x4b\xad\xb6\x68\x42\x33\xc9\x95\xfa\x2b" +
		"\x94\x4f\x89\x14\xda\x62\xd9\x7a\xa7\x74\x32\x40\x28\xc0\x9d\xfe\x42\xea\x43\xb4\xdb\x64\xc8\xd7\x42\x16\xbf\x2c" +
		"\x30\x1d\x3e\x39\xa3\xad\x96\xf9\x69\x98\x14\xd9\xec\xab\x3b\x17\x84\x63\x7b\xba\x43\xeb\x8d\x41\x21\x46\xb3\x56" +
		"\x8a\x29\x04\x3a\x1b\xdd\xe8\xce\xa1\x44\x3e\xc9\x8a\x2e\xf9\xbf\xc4\x92\xef\x63\x58\x98\xe0\xee\x76\xba\xda\xe3" +
		"\x5b\x38\xfa\x2b\x99\x26\x7b\xfb\x57\xd4\x5c\x90\xde\x1c\xf2\x12\x9f\x0f\xeb\xf1\x2c\x6f\x0c\xf9\x49\x68\xc1\x4b" +
		"\x15\x7f\xf7\x26\xa6\x2b\x3a\x9a\x67\x22\x9c\xc2\xf9\xa0\x37\xa3\xcf\x46\xf1\xfd\x40\xee\x6b\xdc\xf8\xe7\xaf\x8b" +
		"\x5e\x77\x7d\x51\xb1\x47\xd7\x95\x3b\x18\x1f\x8a\xa5\x4e\xdb\xed\xd1\x6c\x04\x59\x9a\x3c\x5a\xd9\x04\x1c\xe4\x9e" +
		"\x54\x47\x7c\x93\xf2\x13\x7c\x0c\x1e\x65\x3c\x1d\xd8\x73\x82\xfe\x8f\xdb\x1c\xaf\x2d\x35\x37\xc4\xad\xd1\xed\x2f" +
		"\x74\x28\x1f\x90\xf9\x80\x45\x36\xe5\x63\xe3\xad\xf9\xf4\xf1\xbe\x1c\x2e\x52\xc4\x98\x06\x62\x43\xbc\x9f\xb9\x21" +
		"\xbd\x2f\x0b\x50\x18\x5d\xd8\x11\x2f\x02\xcf\xb4\x5d\x47\xe9\x7f\x74\x6a\xf0\x7a\xf6\xc2\x29\x1a\x7c\xd0\xb2\x60" +
		"\x78\xee\x1e\x9e\xb5\x48\x09\xbe\xd4\xff\x04\x00\x00\xff\xff\x75\xd6\x43\x72\x38\x08\x00\x00")

func bindataSchemagqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataSchemagql,
		"schema.gql",
	)
}

func bindataSchemagql() (*asset, error) {
	bytes, err := bindataSchemagqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "schema.gql",
		size:        2104,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1556055697, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"graphiql.html": bindataGraphiqlhtml,
	"schema.gql":    bindataSchemagql,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"graphiql.html": {Func: bindataGraphiqlhtml, Children: map[string]*bintree{}},
	"schema.gql":    {Func: bindataSchemagql, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
