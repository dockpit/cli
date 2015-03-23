// +build !debug

package uibin

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

var _server_ui_add_dep_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x51\xc1\x6a\xeb\x30\x10\x3c\xdb\x5f\xb1\x88\x40\x4e\xb1\xe1\x1d\x83\x6c\x78\xd0\x43\x7b\xe9\xa5\x5f\x20\x47\x9b\xc4\x20\x5b\xc2\x5a\x37\x0d\x46\xff\xde\x5d\x0b\xbb\xed\x45\x87\x9d\xd1\xcc\xec\xec\xb2\x00\xe1\x10\x9c\x21\x04\x75\x47\x63\x15\xa4\x54\xea\xfb\xbf\xf6\xbf\xb5\xf0\x82\x01\x47\x8b\xe3\xe5\x09\xe4\xe1\xc8\xe4\xea\x2d\x7a\x26\xf7\x7e\xac\xde\xcd\x80\x4c\x3e\xea\x9a\xd9\xa5\xbe\xfa\x69\x00\xa6\xd2\x33\x60\xa3\x4c\x08\xae\xbf\xac\xc4\xfa\xeb\xf4\x78\x3c\x4e\x82\x9f\xe6\xc9\x31\xc5\x5b\x64\x9f\x01\xe9\xee\x6d\xa3\x82\x8f\xa4\xda\xb2\xd0\x11\x1d\x5e\x08\x46\xd6\x6d\x94\xc5\x20\x43\x00\x36\x9d\xcc\x78\x43\xa8\x38\x4d\x94\x74\x32\x2d\x78\x7c\x60\x0e\x9c\x1b\xa8\xf6\x21\x68\x1f\xe8\x36\xf9\x39\x80\x33\x1d\xba\x46\x2d\xcb\x9a\x33\xa5\x73\x56\x2b\x7e\xe4\x3e\x88\x97\x5e\x05\x65\x5e\x16\xfc\xca\x77\x4e\x0c\x9f\xc6\xcd\x28\x9f\xc5\x62\x13\x38\xef\x5a\x4a\x42\xf5\x57\x38\xfc\x6a\xe3\xd5\x44\xce\x97\x33\xa5\x64\xfb\x68\x3a\x87\x76\x59\xb8\xbf\x94\x5a\x8e\xb6\x1a\x14\xbb\x46\xf6\xab\xb3\x61\xbb\xef\xc4\xf4\x2d\xd1\x0a\xae\xcb\x6c\x3d\x6c\xa0\xae\x73\x55\x52\x5a\x37\x13\x71\xe2\xdc\x7a\x9c\xbb\xa1\xe7\x32\x33\xac\xeb\x0c\xf2\x71\x6a\x69\x9f\x5d\xfe\x9c\xfb\xea\x3d\xc9\xb9\xbf\x03\x00\x00\xff\xff\x45\x08\x81\x97\x04\x02\x00\x00")

func server_ui_add_dep_html_bytes() ([]byte, error) {
	return bindata_read(
		_server_ui_add_dep_html,
		"server/ui/add_dep.html",
	)
}

func server_ui_add_dep_html() (*asset, error) {
	bytes, err := server_ui_add_dep_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server/ui/add_dep.html", size: 516, mode: os.FileMode(420), modTime: time.Unix(1427018687, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_ui_add_state_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x90\xc1\x6a\xc3\x30\x0c\x86\xcf\xf5\x53\x18\x5f\x7a\x4a\x0d\x3b\x3b\x81\xb1\x9d\x77\xd9\x03\x0c\x37\x56\x89\x21\xb6\x44\x22\x93\x8e\xd2\x77\x9f\xb4\xe4\xd2\x8b\x91\xa5\xef\xd7\x2f\xe9\xf1\xb0\x0c\x85\xe6\xc8\x60\xdd\x04\x31\x39\xfb\x7c\x9a\x30\xbd\x0d\xef\x29\xd9\x6f\xd6\x3c\xa3\x3d\x0b\x77\xf9\x04\xba\x7c\xc5\x02\x42\x9c\x83\x17\xc4\x84\x1b\x2e\xc5\x42\x1d\xf9\x97\xa0\x77\x91\x68\xce\x63\xe4\x8c\xd5\xdf\xbb\x6d\xdb\x3a\xad\x77\x6d\x99\x05\xc1\x04\xd2\xbc\x00\x4f\x98\x7a\x47\xb8\xb2\x1b\xcc\x29\xe4\x4a\x8d\xed\xae\x67\xb8\xb3\xb3\x59\xca\x09\xe8\xa7\x8a\x95\xb3\xfa\xf6\xee\x3f\xf6\xca\xaf\x30\xc3\xc8\x47\x5a\x65\xda\xe5\x14\x90\xd4\x75\xf8\x68\x2b\x63\x09\xfe\xf8\x0a\xef\x77\x81\x86\xd7\xc6\x8c\xf5\xf0\x5a\xdb\xb5\x64\x19\x61\x5c\x40\x76\x0c\x7e\x2f\xca\x4a\x5e\x67\x1e\x8c\x79\xb9\xcc\x0d\x91\xf5\x32\x7f\x01\x00\x00\xff\xff\xce\x3a\xfb\x93\x2f\x01\x00\x00")

func server_ui_add_state_html_bytes() ([]byte, error) {
	return bindata_read(
		_server_ui_add_state_html,
		"server/ui/add_state.html",
	)
}

func server_ui_add_state_html() (*asset, error) {
	bytes, err := server_ui_add_state_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server/ui/add_state.html", size: 303, mode: os.FileMode(420), modTime: time.Unix(1426946172, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_ui_edit_state_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x91\x41\x6e\xf3\x20\x10\x85\xd7\xe1\x14\x08\xfd\x52\x56\x09\xd2\xbf\x8c\x30\x8b\x2a\x5d\x74\xd3\x2e\x72\x80\x8a\x86\x49\x8d\x8a\x01\x39\x43\x9c\xc8\xe2\xee\x1d\x5c\xdb\x55\x37\x96\xf9\xde\xbc\xe1\x0d\x33\x8e\x1c\xa1\x4b\xde\x20\x70\xd1\x82\xb1\x82\x97\xc2\x54\xfb\x5f\x3f\x5b\x87\xfc\x84\x55\xd8\x52\xd5\xfe\x08\x69\xff\x6a\x3a\x20\x7d\x7b\x38\x4c\x68\x52\x57\xa8\x24\xb9\x98\xba\xc4\xbe\xe3\x10\xce\xf8\x48\xd0\x08\x93\x92\x77\x67\x83\x2e\x06\x79\xdf\x0d\xc3\xb0\xab\xfa\x2e\xf7\x9e\x4a\xa2\x05\xba\xaf\x03\x6c\xa3\x6d\x44\x8a\x57\x14\x9a\x6d\x94\x0b\x29\x23\xff\xf1\x23\xdc\x51\xf0\x40\x57\x34\xc2\x75\xe6\x13\xde\xeb\xbf\xe0\x37\xe3\x33\xa1\x71\x9c\x43\xbc\x54\xad\x26\x29\x45\xf0\x9e\xe6\x88\xc1\x3f\x1a\x31\xf5\xab\x3d\x0c\xb1\xb9\x8d\x8d\xe7\x2f\xe8\x2f\xce\x83\xd0\xab\xff\xb8\xc2\x52\x94\x5c\x1c\x7a\xc3\xc8\xff\x91\x11\x63\xd0\x4f\xd9\x79\xab\xe4\x7c\x5a\xf9\x12\xe5\x01\xd7\x25\xe8\xd5\xdc\xa8\xf7\x89\xbe\xbf\xe5\x4a\xd6\xc1\x35\x63\xf4\x70\x83\xc3\x96\xef\xdf\x32\xd6\x39\xeb\x73\xa7\x1e\x34\x09\xff\x66\x56\x91\x9c\x19\x04\x4b\x47\xf6\x67\x4f\x97\x18\xb1\xee\xe9\x3b\x00\x00\xff\xff\xe6\x27\xa1\x45\xbd\x01\x00\x00")

func server_ui_edit_state_html_bytes() ([]byte, error) {
	return bindata_read(
		_server_ui_edit_state_html,
		"server/ui/edit_state.html",
	)
}

func server_ui_edit_state_html() (*asset, error) {
	bytes, err := server_ui_edit_state_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server/ui/edit_state.html", size: 445, mode: os.FileMode(420), modTime: time.Unix(1426946172, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_ui_foot_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\xd1\x4f\xdb\x3e\x10\x7e\xfe\xf1\x57\x04\x3f\xa0\x44\x45\x49\xfb\xdb\xd3\x08\x61\x6a\xa1\x53\xd1\x58\x87\x28\x62\x1b\x55\x85\x8c\x73\x49\x0d\x8e\x1d\xd9\x97\x76\x59\xc9\xff\x3e\xa7\x0d\x6d\xa0\xdb\x93\xef\xbe\xb3\xef\xee\xfb\xee\x7c\x70\x6a\x98\xe6\x39\x3a\x58\xe6\x10\x11\x84\x5f\x18\x3c\xd1\x05\xdd\xa0\xe4\xec\xc0\x71\x0e\x93\x42\x32\xe4\x4a\xba\xde\x6a\x41\xb5\x43\x25\x15\x25\x72\x66\xa2\x25\x97\xb1\x5a\xfa\xff\x06\x5e\x5e\xa6\xb3\x90\x27\xee\xe1\x16\xf1\xb9\xe4\xc8\xa9\xe0\xbf\xc1\xb3\x81\x36\xbe\x50\xcf\x10\x7b\x4d\x0a\xa6\xa4\x51\x02\x8e\x8e\x1a\xc3\x07\xad\x95\x7e\xe7\xba\x64\x02\x69\x06\x12\x1d\x23\x79\x9e\x03\x3a\x5c\x32\x51\xc4\x10\x3b\xb8\xe4\x0c\x7c\xe2\x85\x20\x0c\xac\xf6\xca\x44\x87\xdd\x70\x07\x66\x80\x73\x15\x9b\x68\x4a\x50\x53\xf6\x3c\x29\x1e\x33\x8e\xe4\x78\xe3\x9d\x0b\xce\x9e\x5f\x9d\x2b\x2e\xb7\xf6\x67\xa5\x33\x6b\xe7\x34\x85\x05\x87\xa5\x35\x79\x6c\x7b\xe1\x49\x69\xcd\x54\xab\x22\x7f\xbd\x69\x4f\x0d\x34\xae\x71\xcb\x9c\x9a\xe6\x95\x3d\x94\x64\xeb\x23\x49\xd6\x0e\x99\xb5\xba\x4a\x28\x43\xa5\xcb\x68\x2b\x3f\x7a\x2b\x0d\x58\x68\xe9\xbc\x9b\x08\x44\x7d\xad\x69\xe9\xe7\x5a\xa1\xaa\x07\xe9\x1b\x51\xd3\x67\x54\x08\x97\xea\xb4\xa8\x35\x32\x56\x0b\xbf\x90\x66\xce\x13\xb4\xa9\x5a\x85\xf2\xc2\xcc\x5d\xf0\xc2\x26\xf9\x36\x50\x55\x61\x62\x45\xae\x2b\x60\xd4\x0d\xf1\x74\x4f\x31\x5f\x80\x4c\x71\x1e\x62\xa7\xf3\xda\xc9\xde\x9d\x29\xb6\x48\x4d\x61\x16\xed\x31\xb4\xb5\xab\x1d\x28\x14\x8d\xdf\x70\xde\xe4\x8d\x15\x5b\xd3\xf0\x99\x95\x12\x61\x28\xa0\xf6\x5c\xd2\x2c\x6a\x4d\xee\xef\x2b\x6c\x03\xd4\x94\x92\xd5\x23\xb7\xc2\x68\x16\xb9\x64\x8e\x98\x9b\x13\x12\x45\xbb\xb4\x42\x31\x5a\x57\xdc\x88\xc8\x94\xf8\xd4\xdc\x0a\x02\x72\xb2\x36\x6b\xcb\xeb\x10\x16\x4b\xdf\x6c\xd6\xce\x6e\x69\x16\xec\x3a\x7f\x32\xc1\xa2\x17\x90\x0e\x76\x48\x0b\xcd\xb8\xb4\x11\x12\xd6\x34\xe4\xae\x5e\x0a\xd8\x70\x30\x83\xf2\x96\xa6\x63\x9a\xc1\x8e\xcd\xb4\x3b\x0b\x6d\x2b\x54\xdb\xf8\x58\xc5\x60\xf7\xd6\x80\xc6\x01\xd8\x81\x80\x0b\xc7\xd2\xab\x5a\x03\x9c\x8c\x2f\xaf\xaf\x87\xb7\x0f\x77\xc3\x9b\xc9\xe5\xb7\x71\x44\x3e\xf8\x5d\xbf\x47\x42\xfb\x79\xdf\xea\xea\x92\xf3\x8b\x1f\xe5\xd7\xd1\x68\xd0\xcb\x46\xdd\xab\x8f\x77\xbd\x51\xff\xfe\xfb\xff\xf6\x43\x99\x9b\x9f\x43\xe8\xde\x7f\xb1\x42\xbe\x79\x55\xef\xa9\xeb\x1d\xfc\xd7\xfa\x41\xcd\x92\xbb\xab\x95\x73\xd1\xbf\xed\x0f\xfa\x93\xe1\xc3\xe5\x85\x53\x55\xc7\xce\xaa\x5a\x3f\xaf\x2a\xd7\x9e\xa7\xc1\x86\xcd\x99\xb5\x1e\x55\x5c\xd6\xe7\x1c\x33\x71\xf6\x27\x00\x00\xff\xff\x1e\xa1\x93\x3d\x74\x04\x00\x00")

func server_ui_foot_html_bytes() ([]byte, error) {
	return bindata_read(
		_server_ui_foot_html,
		"server/ui/foot.html",
	)
}

func server_ui_foot_html() (*asset, error) {
	bytes, err := server_ui_foot_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server/ui/foot.html", size: 1140, mode: os.FileMode(420), modTime: time.Unix(1427018937, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_ui_head_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\x3d\x4f\xc4\x30\x0c\x86\xf7\xfc\x8a\x90\xf9\xc2\xb1\xc1\x90\x44\x42\x77\x37\x30\xc1\x00\x12\x8c\x6e\x62\x88\x45\xbe\x68\xdd\x56\xfc\x7b\xaa\x56\xe8\x3a\x59\x8f\xec\xc7\xaf\x5e\x73\x73\x7e\x3e\xbd\x7e\xbc\x5c\x64\xe4\x9c\x9c\x30\xff\x03\x21\x38\x21\xa5\xc9\xc8\x20\x7d\x84\x7e\x40\xb6\x6a\xe4\x4f\xfd\xa0\xae\x8b\xc8\xdc\x34\xfe\x8c\x34\x59\xf5\xae\xdf\x1e\xf5\xa9\xe6\x06\x4c\x5d\x42\x25\x7d\x2d\x8c\x65\xb1\x9e\x2e\x16\xc3\x17\xee\xbc\x02\x19\xad\x9a\x08\xe7\x56\x7b\xde\x9d\xce\x14\x38\xda\x80\x13\x79\xd4\x2b\x1c\x24\x15\x62\x82\xa4\x07\x0f\x09\xed\xdd\xed\xfd\x41\x8e\x03\xf6\x2b\xc3\x92\x64\x4b\xdd\x5e\x33\x71\x42\x77\xae\xfe\xbb\x11\x9b\xe3\x86\x42\x98\xe3\xd6\xc6\x74\x35\xfc\xba\xbf\x00\x00\x00\xff\xff\xd6\x06\x83\xba\xf2\x00\x00\x00")

func server_ui_head_html_bytes() ([]byte, error) {
	return bindata_read(
		_server_ui_head_html,
		"server/ui/head.html",
	)
}

func server_ui_head_html() (*asset, error) {
	bytes, err := server_ui_head_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server/ui/head.html", size: 242, mode: os.FileMode(420), modTime: time.Unix(1426946172, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_ui_list_isolations_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\xc1\x8e\xd3\x30\x10\x3d\x27\x5f\x31\xb2\x2a\xf6\x00\xad\xc5\x1e\x8b\x13\x09\xa9\x12\xda\x03\x5c\xf8\x00\xe4\xd6\x53\x6a\x29\x89\xad\xc4\xd9\x2e\xaa\xf2\xef\xcc\x38\x4d\xd2\x42\x59\x35\x08\x7a\x48\x1d\x7b\xe6\xcd\xbc\xe7\x37\x39\x9d\x20\x60\xe9\x0b\x1d\x10\xc4\x01\xb5\x11\xd0\x75\xa9\x3a\xbc\xcf\x37\x3a\xe8\xad\x6e\x10\x4e\xa7\xd5\x67\x0c\x7a\xf5\xb4\xe9\x3a\x25\xe9\x24\xa5\xf3\xc7\xfc\xa9\x71\x94\x65\x5d\xd5\xd0\xe6\x63\x4e\x7b\xb5\xa4\x67\x5b\xe4\x69\x42\xa8\xb5\xae\xbe\x23\xac\xa6\x28\xc6\x4d\x80\x7e\xaa\xb0\x79\xbf\x4a\xd4\xde\xd5\x25\x94\x18\x0e\xce\x64\xc2\xbb\x26\x08\xd0\x3b\x0e\xcf\x84\xb4\x63\xaa\xa4\x16\xbe\xe8\x12\xbb\x4e\x1a\x2c\x30\xa0\x18\x00\x92\x4f\xf6\x19\x2b\x78\xa0\x82\x31\x82\x8a\x3c\xac\xdf\x54\xdb\xc6\x7f\x50\xdb\x36\x04\x57\x41\xf8\xe1\x31\x13\x4d\xbb\x2d\x6d\x10\x79\x8d\xa5\x7b\x46\x25\xfb\xc3\xb1\x0f\xc9\x8d\x8c\x6f\x91\x43\x8f\x3f\x32\x59\x18\xf4\xef\x60\xd1\x04\x56\x6a\x9d\xc1\xea\x2b\xaf\x26\x56\x49\x92\x30\x31\x8a\xe7\x48\xda\x5e\x03\xaf\xfb\x78\x16\x6e\x62\xcd\xa0\x58\x99\x29\x55\xc9\xa9\xa2\xd2\x70\xa8\x71\xff\x27\xfe\xda\x98\x25\xe1\x8b\xfc\x2d\x7c\x34\x06\x68\xa9\xa4\x3e\xe7\x9e\x6b\x4c\xe8\x3d\x6e\xda\xab\x7c\x43\x57\x41\x81\xbb\x5e\x1f\xed\x7d\x61\x77\x71\x5b\xbe\x2c\x8f\xc7\xe3\x92\x93\x96\x6d\x5d\x50\x88\x33\x48\xbe\xb8\xba\x26\xaa\xa3\x6c\xe5\xdb\x70\xd6\x37\xe0\x0b\xdd\x9d\xa5\xe3\x11\xfe\x5b\x45\x2d\x0b\xe0\x67\x26\xe2\x5a\x72\xd6\xcd\x6b\xd9\xd5\xc8\x3a\x8d\xb9\xd3\x05\x0d\x57\x13\x3d\xb7\x41\x4f\xd4\xa8\x23\x8b\xaf\xba\x8e\xe2\xe6\xfb\x8d\xb4\x7c\xcd\x69\x17\x16\xfb\xd7\x0e\x8b\x8e\x61\x4f\x5d\xb8\x69\x22\x73\xdb\x68\x29\xfd\xdf\x4f\x8a\x2b\x0c\xc4\xa2\x25\x6f\x33\x1d\x20\x2f\x5c\x78\x07\x80\xc8\x2f\xb5\x61\x3f\xce\xd2\x67\xe4\x71\x25\x11\xf5\x30\x7f\x64\x7e\xb9\x42\x1e\x96\xd8\xec\x30\x2e\xf1\x65\xf6\xc0\x30\xea\x7f\x19\x15\x02\xfe\x8b\x21\x89\x33\xff\xdb\x78\x5c\x7d\xc8\xf7\xce\x51\x89\x05\x51\xfa\x19\x00\x00\xff\xff\x03\xef\xe2\x3f\xe0\x05\x00\x00")

func server_ui_list_isolations_html_bytes() ([]byte, error) {
	return bindata_read(
		_server_ui_list_isolations_html,
		"server/ui/list_isolations.html",
	)
}

func server_ui_list_isolations_html() (*asset, error) {
	bytes, err := server_ui_list_isolations_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server/ui/list_isolations.html", size: 1504, mode: os.FileMode(420), modTime: time.Unix(1427018398, 0)}
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
	"server/ui/add_dep.html": server_ui_add_dep_html,
	"server/ui/add_state.html": server_ui_add_state_html,
	"server/ui/edit_state.html": server_ui_edit_state_html,
	"server/ui/foot.html": server_ui_foot_html,
	"server/ui/head.html": server_ui_head_html,
	"server/ui/list_isolations.html": server_ui_list_isolations_html,
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
	"server": &_bintree_t{nil, map[string]*_bintree_t{
		"ui": &_bintree_t{nil, map[string]*_bintree_t{
			"add_dep.html": &_bintree_t{server_ui_add_dep_html, map[string]*_bintree_t{
			}},
			"add_state.html": &_bintree_t{server_ui_add_state_html, map[string]*_bintree_t{
			}},
			"edit_state.html": &_bintree_t{server_ui_edit_state_html, map[string]*_bintree_t{
			}},
			"foot.html": &_bintree_t{server_ui_foot_html, map[string]*_bintree_t{
			}},
			"head.html": &_bintree_t{server_ui_head_html, map[string]*_bintree_t{
			}},
			"list_isolations.html": &_bintree_t{server_ui_list_isolations_html, map[string]*_bintree_t{
			}},
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

