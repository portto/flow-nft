// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/ExampleNFT.cdc (7.837kB)
// ../../../contracts/MetadataViews.cdc (9.156kB)
// ../../../contracts/NonFungibleToken.cdc (4.832kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _examplenftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\xcd\x6f\x1b\xbb\x11\xbf\xeb\xaf\x98\xe7\x43\x2a\xa1\xb6\xb6\x05\x8a\x1e\x04\x2b\x4e\x10\x47\xa8\x0f\xcf\x08\x1c\xb5\x3d\x04\x46\x1f\xb5\xa4\x2c\xc2\xbb\xa4\x40\x72\xa5\xb7\xcf\xd1\xff\x5e\x0c\xf7\x8b\x5f\x6b\x39\xcd\xad\x42\x10\xaf\x96\x33\xc3\x99\xdf\x0c\xe7\x83\xca\x32\x58\xef\xb8\x06\xae\x81\x08\x60\xbf\x93\x72\x5f\x30\xe0\xf8\x7f\xc9\x84\x21\x86\x4b\x01\x72\x0b\x04\x56\x85\x3c\xc2\xbd\x14\x57\xab\x4a\x3c\xf1\x4d\xc1\x60\x2d\x9f\x99\x98\x64\x19\xdc\x19\xe4\x17\xd2\xc0\x9e\x28\x83\xe4\x66\xc7\x40\x6e\xb7\x3c\xe7\xa4\x00\x6d\x88\xa0\x44\x51\xd8\x54\x06\xb8\x01\xa2\x75\x55\x32\x0a\x46\xc2\x86\x21\xff\x81\xa9\x1a\x34\x2f\x79\x41\x14\xbe\xdd\xc9\x23\x94\x44\xd4\x70\xbf\x5a\x6b\x38\xca\xaa\xa0\x83\x4a\x56\x76\x2e\x15\x83\x6d\x25\x72\xd4\x8f\x14\xdc\xd4\xf3\xc9\x84\x97\x7b\xa9\x0c\xea\xd8\xa9\x68\x35\x84\xad\x92\x25\x5c\xcc\xb3\x70\x61\x9e\xd3\xfc\xa2\xe3\xfa\x95\x19\x42\x89\x21\xff\xe2\xec\xa8\x7b\x16\xef\x6d\x43\x3f\xd9\x57\x1b\xc8\xa5\x30\x8a\xe4\x06\x3e\x37\x88\xdd\xaf\xd6\x8b\x78\xe3\x97\xc9\x04\x00\x00\x19\x0e\xd6\x32\x43\x8a\xaf\xd5\x7e\x5f\xd4\x0b\xf8\xe7\x9d\x30\x7f\xff\xdb\x40\xc0\x0e\x68\xdb\xa7\x56\xee\x9d\xe0\x86\x93\x82\xff\xc1\xe8\x74\x16\xd0\xfc\x9b\x9b\x1d\x55\xe4\x38\xe5\xb4\x13\x73\x69\x15\x5e\xc0\x47\x4a\x15\xd3\xfa\x26\x64\xb9\x65\x7b\xa9\xb9\xf1\x38\x8c\x74\xe9\x7b\x86\x82\xa1\x16\x45\xc1\x2c\xb4\x5f\x8d\x54\xe4\x89\x7d\x21\x66\xb7\x00\xe7\xcb\x08\xf9\x97\x6a\x53\xf0\xbc\xa1\x1e\x9e\x3d\xe2\x5f\xb9\x30\x4c\x8d\xca\xed\x69\x15\xd3\xb2\x52\x39\x83\x24\xb4\xf3\xbb\xfb\xd5\xfa\xd2\x77\xda\xfc\x81\x69\x59\x1c\x98\x82\x17\x2b\xc5\xdd\x75\x30\x7c\x12\xad\x09\x52\x32\x54\x42\x71\xf1\x14\x2d\x52\xa6\x73\xc5\xf7\x68\xdc\x28\x8d\xd9\x55\xe5\x46\x10\x5e\x44\x14\x24\xcf\x99\xd6\x53\xcd\x8a\xed\xcc\x92\x2a\x59\x93\xc2\x70\xa6\x17\xf0\x2d\x50\xde\xae\xd4\x8f\xb1\x7e\x8c\x72\xdc\xfe\xbe\x2a\x37\x4c\xf5\x66\x84\x54\x25\xf9\xfd\x73\x43\xd8\x91\xdc\x0c\xa2\xb8\xe0\x66\xda\x7f\xb3\x6f\x86\x48\xf0\xde\xbb\x60\xf8\x2b\x09\x24\x7c\x82\x08\x06\x7f\xf9\xbc\xe9\x3e\x7d\xd2\x6c\x9f\x24\xb6\x79\x58\x9f\x39\x51\x80\x1f\xf4\xc1\x9c\x53\x58\x02\xa7\xf1\x02\x9a\x0d\x4b\x6b\x7d\xbc\xe8\x58\x0e\x4b\x17\x87\x98\xb4\xc7\x00\x96\x03\x1e\x31\x59\x8f\x05\x2c\x07\x5c\x62\x32\x0f\x02\x58\xfa\x90\xc4\xe4\x03\x1c\xb0\x74\xb0\xe9\x09\x4f\xf6\xc9\x8b\x9c\x6d\x25\xe0\x89\x19\xeb\x87\xe9\x6c\x01\xdf\xd6\xf5\x9e\x3d\x06\xd0\x29\x66\x2a\x25\xe0\x9b\xf7\x12\x3f\x48\x7c\xed\xfb\xf2\x96\xeb\x7d\x41\xea\xf7\xd3\xd9\xe5\x5b\xc8\x1f\x3a\xe3\xdf\xca\xd0\xda\xf4\xbe\xcd\x8a\xdd\xe7\xd1\x31\x32\x32\x50\x35\x79\x01\x05\x4c\xff\x03\x07\xce\x8e\x0b\x2b\x7a\xb6\x80\x8f\xa2\xfe\x6a\x54\x95\x9b\x9b\x30\x5c\x8e\xdc\xe4\x3b\x4b\x1c\xac\xe0\x27\x27\x9a\xbd\x6e\xfd\x22\xe2\x71\x90\x4c\x32\x4d\x93\x1c\xd0\x9f\xc9\x3e\x4e\x63\x9c\xba\x8f\x77\x44\xc3\xd0\x1d\x67\x73\x0e\xae\xaf\xd9\x3f\xd6\xeb\x2f\x2b\x5e\xb0\x71\xd5\xf0\x53\xa9\x62\x11\x44\xff\x28\xfd\x2c\xb9\x12\xbf\x1d\x03\xd8\x8d\x97\x1f\x80\xb8\x67\x1b\xb7\xc4\x3f\x97\x3f\xa9\xe6\x10\xa5\x3f\xa0\x64\xcb\x34\xae\x62\x90\x11\xe3\x14\x31\xee\x62\x37\x53\x06\xb9\xe2\x0d\xa6\x9e\x52\xe9\x40\x38\x6e\x6e\x08\x4e\x89\xea\x6d\xab\xfd\x96\xe4\xcc\xe9\x94\xc2\x7e\x21\xa8\xd7\x78\x64\x69\xdb\xb0\x18\x2c\xf6\x0b\xf8\x10\xd5\xff\xfb\xd5\x7a\x96\xca\x64\x77\xb7\x4d\x1e\x6b\xca\xc1\x63\x44\xb2\x91\x4a\xc9\xe3\xfd\x6a\xed\xf4\x42\xb3\x05\xbc\x4b\x6d\x30\xc2\x3c\x18\x12\xc8\x18\x16\x90\x3b\xcc\x28\x7b\xa9\x4d\x22\x95\x4c\x15\xd3\x55\x61\x60\xb9\x44\x44\x67\xf0\xfd\x7b\xf7\xea\xc6\x96\x2b\xac\x57\x23\x51\x74\xf1\x89\x08\x6c\xba\x1b\xb5\x1c\x80\x41\xb1\x2d\x53\x4c\xe4\x6c\x61\xbb\xe5\xbb\xdb\xae\x27\x6f\x7c\xc7\xe8\x40\x81\x9d\x3b\x17\xb9\x54\x8a\xe5\xe6\x62\xc4\xed\xe3\xfe\x1d\x7c\xb9\x78\xc5\xc3\x97\x71\x03\xf7\x45\xc9\x03\xa7\x4c\x25\x96\x1e\x58\xce\xf8\x21\xb9\x14\x0b\x4e\xb7\x80\x03\x9d\x03\x79\x96\x01\xe5\xcd\xc4\xa0\x6a\x44\x04\xa1\xca\xa5\xd8\x4a\x55\x72\xf1\x04\x36\xd8\xb4\x4b\x8e\x04\x38\x19\x0d\xf6\x9a\x7a\xcf\xe0\xc8\xcd\x0e\xc7\xa5\xdf\x1a\xdf\xff\x86\x00\x6f\x39\x2b\xa8\x17\x31\xd8\xf2\xcb\xa3\x60\x14\xa7\x98\x05\x7c\x78\x69\xa8\x13\xcd\xec\xfd\x6a\x7d\xf2\x9b\x35\x98\x26\xfb\x97\x5e\x1c\x5c\x5f\xc1\xcb\x29\x55\xf4\xb2\xcc\xaa\x87\x03\x02\x28\x56\xca\x03\xb3\x93\x1d\x5a\x62\x87\x9a\x66\x7a\xea\xd1\x21\x82\x42\x43\xc4\x0d\x8e\x5e\x76\x99\x14\x85\xd3\x5f\x74\xd1\xdf\x89\x9d\x76\x0f\x77\xb7\x4e\xf4\x27\x8f\x68\x60\x83\x6d\x96\xed\x64\x74\x7d\x15\x18\x34\x6f\x74\x9d\x3e\xb3\x7a\x01\xc3\x06\x33\xb8\xb9\x81\x3d\x11\x3c\x9f\x5e\x94\x5c\x6b\x74\xd3\xfd\x6a\x7d\x31\x9b\xf8\xbd\x62\xc9\x83\xb9\xc8\x6e\x33\xe7\xb4\x9b\x8c\xfa\xdd\xd4\xcd\x9c\x34\x53\x4f\x20\xa3\x4d\x6b\xd7\x57\x96\x75\x04\xda\x36\x2f\x81\x21\xcf\x88\xab\x85\x15\x21\x24\x94\x7a\x08\xf6\x00\x6b\x27\xe4\x5c\x41\x3d\x53\x7b\x3e\x5b\x46\x4e\x81\x28\x45\xea\xff\x2d\x21\xbe\x06\x77\xf3\x40\xf4\x2f\xf0\xc1\xcf\x53\x93\x88\x67\xc8\x6a\xd8\xc5\xb6\x40\xfa\x64\x68\x01\xa5\x56\x65\xc1\x8e\xad\xf0\xd6\x06\xe7\x8c\x1d\x77\x3c\xdf\xf5\x61\x68\x2f\x05\x0a\x0a\x52\xb0\x68\x4f\x59\xd0\x75\x3a\x32\xbe\x71\xfa\xd8\x1b\x90\x70\xbb\x3b\xdb\xa2\xbf\x71\xae\x3d\xef\x6d\xca\xb4\x51\xb2\xee\xf7\x1d\xf1\x77\x53\x51\xda\xd8\xb0\x07\xc9\xba\xa7\x4b\xa7\xb8\x66\x76\xc4\x00\x51\x58\xea\x02\xdf\xbf\xa1\x3e\xa5\x3b\xed\xe0\x68\x3c\xb3\x5a\x8f\xe8\xd7\x97\x33\x94\xdd\x24\xaa\x2e\xaf\x1b\xd9\x9d\xfb\x71\xc5\xb2\x0c\xb4\x6c\x2c\x18\x0e\x3e\xe4\x04\x7b\x66\x42\x81\x1b\x0d\x65\x9b\x5f\x6d\xc4\x22\x41\xf7\x76\x27\xa9\xfe\xa9\xf2\x9a\xb6\xfd\x5d\xc2\xfb\x44\x9f\x29\xd0\xa7\x49\x3c\xd6\xfc\x54\xb1\xe6\xdb\x54\x14\xfe\x62\x6b\x74\xa2\x88\x67\x19\x7c\x52\x8c\x18\x66\x23\xa4\x32\x3b\xa9\xf8\x1f\x5e\x95\x45\x6f\x14\x85\x3c\x02\x95\x47\x91\x13\x6d\xdc\x6b\x02\xf7\x20\x28\xb6\x85\xe5\x18\x0a\x28\xfa\x0c\x14\x01\x9c\x28\x0e\x0f\x7d\x60\x6f\x50\xe7\xcf\xf7\x77\x23\xf0\x62\xd5\xed\x6a\x6e\x00\xf0\x47\x51\x3f\xb4\x55\xf3\x25\x5d\xa4\x4f\x89\x7c\x25\xb6\xe6\xa7\xcd\xb7\x97\x26\x43\x33\xb4\xb4\x42\xcf\x81\xd0\x5a\xed\xf0\x61\xd4\xbd\xc1\x88\x14\x48\x6d\x76\x89\x8a\x78\x97\x75\x7c\xf3\xd2\x6d\x56\x96\x21\xd6\xd8\x1e\x77\xf7\x9c\x6d\xaa\x11\xb5\x14\xcc\x1e\x52\x7b\x1c\x8d\x84\xbc\x8d\x3d\x9b\x8b\x59\xb9\x37\x75\x78\xd8\x3b\xaf\x35\x94\x9f\x91\x64\x68\x91\xa6\xc9\xf2\x9d\x6c\xa1\xfa\x22\xd9\xed\xe9\x4a\x09\xb4\x7f\xe8\x7b\xa6\x46\x6d\x20\xb4\xe4\x02\xa4\x02\x2d\x31\x7f\x60\x2d\xef\x2e\x7d\x9b\x3b\x5e\x79\x14\xed\xa5\x70\x2b\x82\x6c\x0a\x7b\x74\x4a\x2e\x8c\x35\xae\x87\x2b\xcb\x92\x37\x85\xcd\xed\x62\x77\xf1\xda\x4a\x41\x6e\x74\x28\xfe\xd5\x2d\x4a\xf8\xbd\x69\xe3\xec\xd7\xbb\xdb\xb0\x38\x77\x95\x1e\xff\x89\xb6\x77\xce\xf9\x9e\x33\x94\xe1\x34\x50\x95\xed\x49\xcc\x8e\x71\xe5\xbe\xee\x4f\x7e\x74\x70\x5a\x6d\xa6\x41\xf4\xb5\xb2\x17\xf0\xee\xe5\x6c\xd7\x7b\xfa\x3f\xbe\xb3\x0b\xdb\x0c\x2f\xb6\xc3\x43\x8b\x2d\xb6\x60\xb6\xfe\x0d\x21\x19\x81\x0b\xed\x7d\xa7\x73\xfa\x9d\xab\xf8\x78\x68\x6e\xf0\x4c\x5f\xb5\x78\x88\xbe\x7a\xbb\xe2\x40\xdb\x3f\xc6\x54\x0e\xc2\xfd\x63\x4c\x15\xe0\x7a\x66\xe8\x77\x21\x1e\x9e\x7d\xba\x59\x04\xf4\x2b\x01\xff\x27\x0d\x24\xcf\x65\x25\x8c\x17\xee\x71\x8c\x83\x1b\xca\xf3\xa0\x67\xbd\xbe\x6a\x5c\x15\x6c\x9d\xf6\x0a\x2c\xc7\x16\xfe\xdc\x86\xcd\xf4\xaf\xb3\x74\xe2\xb4\x97\xdd\x33\x7f\xee\x1b\x7e\x4f\xb1\x96\x59\x79\xa0\xad\xc0\x9e\xac\xb9\xc2\xf2\x54\xf8\x8b\x97\x47\xbe\xb2\xa6\x4d\xc2\xc8\xa0\xb0\x27\x66\xa7\x7d\xe6\xe4\xef\x26\xb0\x84\x4c\x37\x5f\x33\x96\x98\x90\xc7\x44\x0c\xbf\x9f\xa0\x84\xa6\x14\xbc\x41\x40\xf4\xfb\x4a\x7a\xff\x86\xcc\x33\xaf\xeb\x60\x9c\x9c\x3e\xe4\x56\x4c\x89\x9a\x1c\x58\x3b\xe4\xb4\x02\x7b\x76\x2c\xb9\x4e\xf6\x7b\xa5\x3e\xf4\x8a\xb6\x11\x35\x47\xa9\xd3\xeb\xab\x81\xdb\x69\xe1\x93\x80\xce\x3c\xad\xfb\xfc\xd0\x16\xcb\x9c\xec\xc9\x86\x17\xdc\xd4\xb0\x95\x6a\xac\xf1\xf5\x34\x28\xb8\x78\xbe\x76\x7b\x83\x61\xdb\xf3\xb9\xf8\xd2\x8d\xd3\xf1\x1b\x90\xd3\xfb\x69\x3c\xce\xa7\x9c\x1d\xe4\x67\xa2\x9e\x98\x79\x0d\x8d\x49\xe2\x44\xbb\xce\x6c\x2b\xe2\x8f\x38\xb2\x6c\x58\xbc\x8c\xda\x88\x39\xe3\xc3\x86\xd1\xf1\x5f\x14\x8c\x8e\x92\x76\x78\x1b\xff\xbd\xf3\x34\x39\x4d\xfe\x1b\x00\x00\xff\xff\xdb\x72\x05\xa0\x9d\x1e\x00\x00"

func examplenftCdcBytes() ([]byte, error) {
	return bindataRead(
		_examplenftCdc,
		"ExampleNFT.cdc",
	)
}

func examplenftCdc() (*asset, error) {
	bytes, err := examplenftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ExampleNFT.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0xb4, 0x78, 0x0, 0xf9, 0xbc, 0x81, 0xeb, 0x48, 0xf0, 0xe1, 0xc9, 0x5e, 0x1b, 0x2, 0xcc, 0xc4, 0x98, 0xdc, 0x83, 0x33, 0xcc, 0x8c, 0x7c, 0x9d, 0x22, 0xc5, 0xa, 0x82, 0x49, 0x54, 0xd6}}
	return a, nil
}

var _metadataviewsCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x18\x64\x81\x22\xee\xd9\x92\x93\xeb\x06\x7b\xc6\x7a\x17\xd9\xb6\xe9\xe6\xd0\xf6\x82\x34\xdd\x7b\x58\x14\x0d\x25\x8d\x6c\x5e\x29\x52\x47\x52\x71\x7c\x45\xbf\xfb\x61\xf8\x47\x96\x2c\x25\x97\x5d\x5c\x1f\x52\x4b\x22\x87\x33\xbf\xf9\x3f\x4c\x9f\x3f\x9f\x4c\x6e\x36\xdc\x40\xae\xa4\xd5\x2c\xb7\xc0\xab\x5a\x60\x85\xd2\x1a\xb0\x1b\x84\x0a\x2d\x2b\x98\x65\x60\x2c\x93\x05\xd3\x05\xd4\x5a\xd5\xca\x60\x31\xe1\x12\x2e\xde\x5e\x5e\xcd\x17\x67\x7f\x3d\x4b\x26\x93\x6b\x2c\x97\xb0\xb1\xb6\x36\xcb\x34\x5d\x73\xbb\x69\xb2\x24\x57\x55\xaa\x64\x29\xd4\x36\x75\x7f\x32\xa1\xb2\xb4\x62\xc6\xa2\x4e\x4b\xc1\x6b\x93\x9e\x2e\x4e\x4f\x16\x7f\x3b\x39\x9b\xcb\xd2\xce\xe3\x61\x49\x55\x4c\x26\x1f\xac\x6e\x72\x6b\x80\xc9\x02\x34\x1a\xd5\xe8\x1c\x0d\xe4\x4c\xee\x59\x04\x25\x11\x94\x86\x4a\x69\x9c\xb4\x9c\xda\x5d\x8d\x66\x06\x39\x13\x02\x0b\xb8\xe3\xb8\x35\x09\xbc\x66\xf9\xc6\xfd\x76\x9f\x41\x63\xad\xd1\x90\x94\x13\x06\x05\x2f\x4b\xd4\x44\xef\x0b\x97\x05\xa8\xb2\x95\x7a\x06\xa6\xc9\x37\xc0\x0c\x30\xc8\x35\x32\xab\x34\x64\x5c\xad\x35\xab\x37\xbb\x89\xd2\xc0\xe0\xef\x57\xaf\xdf\x00\xaf\xd8\x1a\xa1\xe4\x02\x93\xc9\xf3\x74\x32\xe1\x55\xad\xb4\x85\x8b\x46\xae\x79\x26\xf0\x46\x7d\x41\x09\xa5\x56\x15\x2c\xee\x11\x7f\x38\xfd\xe1\xfb\xb3\xac\x3c\x5d\xe0\x29\x63\x67\x93\x49\xdd\x64\x7b\xfc\xdf\x85\xa3\x7f\x23\xbe\xe1\xeb\x64\x02\x00\x90\xa6\x29\x9c\xc3\x35\x1a\x25\xee\x50\x93\x0a\xee\x78\x81\x06\x58\x9e\xa3\x31\x60\x15\x30\x30\x68\xbb\xac\x07\xc1\xe3\xf6\x0e\x19\xe3\x80\x25\xdc\x22\xac\x70\x8c\xc9\x3a\x01\x26\xe1\xfd\xc5\xcd\xf4\x00\x63\x4b\xe6\xc1\xa5\x45\x5d\xb2\x1c\x5b\x3a\x56\x45\x36\x3a\x5c\x90\xc5\xb8\x73\xc1\x6e\x98\x05\x6e\xc1\x34\x35\x41\x71\xc0\x08\x49\xdc\x1e\xde\xd2\xde\x0b\xf8\xd5\xad\x8a\x2b\xcb\x46\xc2\x1a\xad\x43\xe4\x78\xba\x84\xdf\x6f\x76\x35\x7e\x1a\x2c\xd1\x7e\x37\x2d\x3b\xfe\xec\xd8\x58\x02\xad\x9c\x2e\xe1\x5c\xee\xbc\x39\xfd\xec\x76\x7d\x1b\x43\xf5\xa5\x12\x02\x73\xcb\x95\x04\x4e\xfa\x5e\x6b\xd5\xd4\x84\xa8\xb3\x9a\x40\x5c\x13\x14\x05\xde\x43\xb6\x83\xcb\x57\x7f\x48\xa8\x0e\xfd\xa1\x78\x99\xd2\x5a\x6d\x89\xf5\xb8\xfc\x98\x17\x4b\xf8\x78\x29\xed\xd9\x8b\xe9\x12\x9e\x7d\x8d\xef\xbf\x8d\x41\x73\xf9\xca\x03\xe3\xd7\x7f\x3a\x14\xf2\x15\x37\xb5\x60\x3b\x2f\x57\xc6\x0c\xcf\x83\x2b\x38\x25\xc9\x5c\x34\x64\x4c\xa4\x3c\xc9\x2a\x9c\x41\x81\x26\xd7\xbc\x76\xbc\x32\x59\xec\x75\xbe\x69\xaa\x4c\x32\x2e\xa0\x24\xdb\x97\xa0\xb2\x7f\x61\x6e\x13\x78\xa7\x8c\x0d\x0f\x06\xcc\x46\x35\xa2\x38\xb4\x20\x3a\x70\x88\x57\xb0\xc5\xc8\x60\x30\xf7\x78\xde\x4d\xe0\x88\xb4\x40\xdc\xc5\xe3\xba\x8b\x0e\x36\x70\x03\x25\x47\x51\xc0\x96\x0b\x01\x19\x42\xe1\x49\x63\x01\x5c\x82\xe0\x26\x44\x14\xbb\x41\x8d\xa5\xd2\x18\xd8\xed\x91\xc9\xdc\x5b\x6d\x49\xc4\x5c\xc9\x9c\x1b\x4c\x46\xcf\x24\x11\x04\x5a\xc7\xe4\x12\x3e\x58\xcd\xe5\xba\x2f\xc2\x39\x6c\x35\xb7\x16\x65\x0f\xd4\xff\x97\x3c\x0c\x0a\xb4\x8c\xc7\x38\xd7\xa7\x3b\xeb\x91\x32\xca\xf9\x75\x86\x2e\x5a\xc2\x1d\xea\x4c\x99\xd6\xf3\xa1\x66\x9a\xb9\xb0\x06\x5c\x1a\x8b\xcc\x85\x41\x06\x86\xcb\xb5\x40\x10\x5c\xe2\xf4\x71\x08\x3a\xe2\x3d\x84\x84\xa9\x98\x10\x1d\x23\x6a\x83\x30\x1b\x01\xe5\x29\x98\x04\x4b\xcb\x10\x18\x6c\x31\x9b\x97\x9a\xa3\x2c\xc4\xce\x45\x62\x38\xe6\x09\xba\xf0\x3c\x83\xab\xf7\x6f\xa6\x3d\x22\xce\xf2\x03\x1e\x43\x0b\x99\x91\xc0\x5f\xa0\xd6\xe8\x82\xd9\x0c\xd0\xe6\x8f\x4b\xdf\x0a\xd5\x89\x35\x5f\x2f\xb8\xc0\x6f\x7b\x10\xb8\xe4\xf6\xb8\x7d\xa2\x7f\x5d\xb3\x99\xf5\xbe\x8c\xa0\xd9\x5f\xf0\xc8\x81\x71\xc9\xb4\x13\x67\xe8\x9f\x41\x51\x26\xce\x9d\x56\xee\xe4\xe1\xc7\xae\x89\xae\xba\x3c\x0c\x97\xee\xb5\xb8\xda\xf3\xd2\x2e\xfb\x76\x18\x81\x88\xb3\x10\x56\x51\xa2\xe6\x79\x27\x40\x3a\x5d\xec\x13\x32\x30\xaf\x3e\x63\x95\xc6\x02\xc8\x30\x34\xa8\xb2\x84\x7c\xc3\xb8\x1c\x26\x34\x22\x6d\xa2\x2e\x1b\x83\x05\x25\x22\x8d\x2e\xa3\x53\xc5\xe0\x72\xb3\x99\x01\xa5\x2a\xe5\x9d\x5f\x91\xf7\x43\x85\x05\x67\x0f\x86\xa4\x3d\x7f\x8e\xf7\x61\xc8\x6e\x34\xa7\x98\x1b\x4c\xfd\x40\xde\x5f\x6f\x6e\xae\xf6\x32\x3b\x79\x7c\xb0\x8d\x39\x9b\xca\x02\x60\x2e\xc2\xd0\x5a\x38\x56\xda\xfd\xf8\x30\x85\x8f\xd7\x6f\x43\x48\x18\x61\x2b\x12\x5e\x8e\xb1\x45\x96\xd8\x68\x31\xf4\x3f\x67\x7a\x9d\x2f\xa3\xa6\xd1\x68\x52\x66\xa3\xbb\x6a\x7c\x5c\xea\x03\x2a\x1a\x6d\xa3\x65\x4b\xec\x61\x6b\xb8\xbc\xba\xf8\xe0\xd8\xf7\x3b\x08\xa2\xbd\x3d\x85\x5a\xaa\x9b\x60\xda\x8d\xc1\x26\xa8\x1e\x93\x9d\xa2\x8b\x1c\x97\x68\x0e\x8d\x83\xde\x06\x0b\x00\xa6\x71\x6f\x17\x05\x65\x70\xbb\x41\xae\x5d\xe9\x45\x69\x8a\x17\x28\x2d\x2f\x39\x6a\x38\x7e\x79\xf9\x6a\xda\x12\xd1\xcc\xd9\x8b\xdd\x30\x17\x71\xb9\xc6\xdc\xc2\xc7\xeb\xcb\x04\xce\x21\x17\x9c\xf6\xb2\xba\x16\x3c\xf7\x71\x8c\x4c\xb1\x31\xe8\xd3\xde\xcb\xcb\x57\xdd\x92\xa9\xa4\x02\x93\x4c\x50\x28\xe6\x92\x50\x90\xe2\x8e\x33\x12\xc9\xb1\xbb\x66\x16\xb7\x6c\xf7\xa0\x65\x46\xf4\x5a\x13\xe8\x45\xb6\x97\x97\xaf\xc8\xca\x88\xf4\x88\x60\x04\xab\xe3\xcb\x9d\xe4\x0b\xd6\xce\xee\x1e\xa5\x5e\x41\x5f\xa8\xdc\x24\xbc\x2e\x4d\xc2\x55\x4a\x39\x11\x6b\x6b\xd2\x70\xc2\x9c\x15\x85\x26\xa3\x96\xeb\xf4\xd1\x30\x99\x53\x51\x33\x96\x1c\xae\x98\xdd\x38\xe7\x90\xa0\x5c\xc4\x61\x02\x6a\x7a\x17\x8a\xca\xd2\x5b\x4b\x5b\x5c\xb5\x60\x79\x6d\x28\xbd\x7b\x52\xc2\xe0\x06\x94\x14\x3b\x90\x88\x05\xc5\xfb\x72\x4f\xdc\x15\xb9\xc6\x95\xb3\x4f\x21\xfa\x04\x70\x88\xec\xdc\xec\x8c\xc5\xca\x3c\x0e\x0b\x49\x1a\x71\xf9\xf9\xc0\x6b\x3b\x90\xcd\xfa\x0b\x47\x9d\x38\xe7\x05\xac\x08\xe7\xe1\x27\x87\xe7\xca\xd1\x18\xf3\xf0\x3d\x54\x8d\xf4\x65\x6a\xf4\x4e\xc2\xc8\x81\x2d\x99\xe5\x77\x48\x01\x6a\x6f\x48\x7f\xd6\x86\x36\x6a\x3b\xb7\x2a\x0d\x96\x33\xa7\xd7\x73\x25\xe7\x5b\xcc\xd2\xef\xfc\x39\xf3\x46\x0b\xf3\x20\x6e\xff\x2b\x18\xf1\xb2\x85\x16\x56\x1d\x00\xfa\xab\x60\x1f\xb6\x8e\x88\x85\x65\x9a\x1e\x25\xa4\x41\x66\x8f\x23\x9e\xd3\xf8\xe2\x28\x3d\x6a\x7f\x13\xad\x69\x8f\x54\x07\xca\x27\x51\x7d\x20\x3c\x3e\x77\xff\x3d\x07\xb8\x56\x3b\x26\xec\x0e\x5c\xdb\x13\x5f\xbe\xc2\x92\x4b\x8c\xee\x5d\xd5\xca\x30\x4a\x24\x3a\xac\x6d\x9b\x74\x97\x6c\xd6\xfc\x0e\x0d\x54\x4c\x7f\x41\x5b\x0b\x46\xc9\x90\x41\x23\x29\x12\x14\x07\x1d\xdd\x73\x20\x4f\x0b\xcd\x1a\x75\x81\x81\x24\xc7\xd0\xb9\xb5\x6c\xbd\xeb\x92\xa3\x50\xf7\xef\x06\xf5\xce\xdb\xc2\xed\x75\xdc\x74\x1b\x83\x95\x6b\x7a\xdf\x5f\xdc\x18\x88\x04\x28\xfa\x51\x28\xc6\xfb\x1a\x73\xeb\xd3\x75\xcd\x76\xfb\x03\xa9\x43\xf1\x89\xdf\x6e\xd0\x20\x98\x1a\x73\x5e\x86\xe8\xda\x67\x67\x10\x1b\x5b\x06\x0e\xc3\xe2\xb9\xd6\x6c\xe7\x61\xa1\x3e\xfb\x8b\x87\x90\xcb\x82\xdf\xf1\xa2\x61\x62\x7f\x7c\xbb\xcd\x27\x6a\xa7\xb0\xa9\x8f\x5e\x8d\xbd\x94\xa5\x32\x4b\xf8\x3d\x28\xe7\x53\x3f\x41\x3a\x87\xfd\x3c\xb6\xee\xd0\x57\xd3\x14\x7e\x63\x82\x17\xcc\x86\xc2\xc0\x34\x95\xab\xb8\x85\xa0\xed\x50\x35\xc2\xf2\x5a\x70\x6a\x38\x43\x95\x2b\x95\xa5\x0a\x67\xad\x91\xd9\x98\x8d\x4e\x92\x45\x8f\xec\x1d\xd3\x60\x95\x65\xe2\x65\x63\x61\x05\x8b\x83\xcf\xe4\xb4\xd1\x54\xb8\x6c\xf9\x1c\xf1\x89\x0e\x91\xf6\xe7\x5f\xe2\xde\x24\x6f\xec\x81\xe9\x77\x9f\x98\x31\xa8\xed\x71\xbb\xef\xc7\x15\xf1\x39\x83\x0a\x8d\x61\x6b\x5c\xc2\xd1\x07\x2f\x6c\x7b\xfe\xd3\xa5\x3d\x9a\x1e\xc2\x78\x6e\x0c\x5f\x4b\xef\x10\x81\xde\x48\x54\x8c\x27\xad\x86\x8b\x0e\x42\xe0\xb5\xf7\xdb\x2e\x3d\xd7\x14\x8c\xb5\xdd\xad\xb5\xb9\xe6\x3b\xea\xfa\x91\xaa\x68\xe4\x70\xe8\x97\x46\xbe\x92\x77\xce\x48\xb5\x4e\xe8\xef\x4c\xaf\x1b\x8b\x2a\x24\x43\x71\x75\x92\xf3\x74\x37\xbc\x19\xf7\x87\x41\x63\xfd\x26\x14\xe1\xfd\x01\xd5\x35\xe6\xc8\xef\xda\x2a\x01\x21\x43\x89\x25\xcf\x39\xd3\xbb\xd8\x9d\x85\xb3\xfb\x25\x07\x73\x68\xc4\x9a\x23\xd7\x48\x56\xbd\xab\xdb\xbe\x5d\x47\xc2\x5b\x6e\x37\xed\x53\xb2\x46\x7b\xb3\xab\xf1\x78\x7a\xa0\x80\x5c\x55\x15\xca\xc2\x57\x53\x73\xf8\x68\x3a\x76\xe1\xa6\x70\x94\xa3\x25\x6e\x7d\x97\xe6\x11\xb8\x10\x6a\xeb\xa5\xd0\x7d\x29\xb8\x81\x86\x70\x83\xdb\x56\x61\xbb\x28\xe8\x55\x93\x09\x9e\x53\xed\x71\x3c\xbd\xed\xb7\xcb\x14\xa5\xc8\x06\x7d\x21\x47\x7a\x28\x59\x23\xec\xc8\x39\xc9\x30\x83\xba\x5e\x9d\x09\xa1\xb6\xb4\x5f\xbb\xd1\x58\x53\x07\x67\x47\xc8\x59\xcd\x32\x2e\xb8\xf7\x42\x57\x81\x34\xb6\xd1\xe8\x96\x19\x12\xce\xb5\xe8\xb1\x53\xda\x2f\x1f\x14\x0e\x91\x87\x25\xbc\x6c\x17\xfd\xf8\xec\x5c\xee\xae\x43\xa9\xf4\xb5\xa7\xe1\x24\x0a\xfe\xed\xa7\xbe\x3d\xbc\x6b\x1d\xb0\xed\xa0\x72\x26\xf2\x46\x44\x96\x59\xa5\x1a\xe9\xe6\x8b\x86\x09\x84\x3b\x26\x1a\xa4\x40\x2a\x4d\x89\x5a\x87\x9e\x2b\xd8\xda\x38\x30\xef\x95\x45\x98\xc3\xa5\xed\x74\xed\x19\xda\x2d\xa2\xa4\x38\xe5\x00\x3f\x49\x16\xbd\x59\x08\xbc\xbe\xa7\x2d\xde\x88\x3a\x07\x73\x03\xf7\x6e\xc3\x3e\x70\xd0\xbb\x45\xf2\xfd\x19\x2d\x95\x5d\x4b\x0d\x5b\xb6\xf1\x4c\xb7\xe8\x39\xdc\x3f\x5c\xb3\x38\xe7\x60\x42\xec\xa0\x46\x9d\xa3\xb4\x54\xa0\xaf\xb1\xd3\xa3\xfa\x51\x81\x45\x5d\x39\xc7\xcc\x98\xe1\x06\x6a\xc5\xa5\xed\xd5\x2b\xb4\xc8\x28\xc1\x0b\xd2\xb4\x4f\x6a\xa6\x62\xda\xb6\xc3\x5e\x03\xdb\x0d\xd5\x9e\x39\x2b\xa8\x27\xa1\x2e\x97\xac\xe5\xf6\xe3\x05\xbf\x3f\x7b\x71\xeb\x33\x03\x13\x1a\x59\xb1\x6b\x27\xa9\x07\x73\x0c\xec\x1d\xef\xec\x27\x67\x86\xb0\xcd\x19\x3d\x70\x6b\x40\xd5\xa8\x7d\xf6\xec\xdb\x38\x65\x61\x69\xb9\x46\xb1\xa3\x40\x83\xba\xe2\x92\x1b\x1b\xba\xf3\x35\xea\xce\x4e\x87\x77\x2c\x10\x9a\x9a\x14\xfe\x43\x3c\x54\x95\x50\x6b\xcc\xb9\xe1\x4a\x26\xc3\xa2\xbf\xb1\x4b\xf0\x22\xf5\xad\xee\x1f\xb1\xd2\xef\x4d\x3b\xfc\x75\x84\x6f\xe7\xbd\xb7\x90\x18\x74\x04\xdb\x91\x1f\x77\x74\x3b\x1b\x60\xa1\x51\x78\x76\x37\xbc\x6e\xcd\x8b\x3e\xdc\x6e\x99\x10\x68\x6f\xe3\xf0\x8f\x82\xe5\x0c\x5c\x93\xb9\xb3\x1b\xa2\x8b\xc2\x84\x5c\xec\xa6\x50\x5b\x89\x1a\x2a\xbe\xde\x58\xd8\x32\xe9\x23\xb2\x2b\x44\x86\x5e\xf8\xe8\xe8\xcb\xd5\x04\xe4\x12\x35\x35\x8a\x7f\xdc\x51\x67\x5d\xfc\x66\x63\x67\x1d\xd6\x16\xb5\xc6\x91\x8c\x4e\xc9\xe2\x27\x57\x12\xc0\xb3\x67\xee\xc9\xe7\x65\x58\xc2\x11\x65\x69\xef\x26\x7b\xdf\xe4\x92\x5e\xf1\x02\x34\x93\x6b\x04\x9e\x20\xfc\xbe\x98\x9d\x7c\x3a\x7a\x24\xf3\xbb\x1c\xd7\xc6\xdf\x15\xb4\x62\x8f\x66\x64\x9f\x8c\xff\xcc\x14\x6a\x90\x35\xdf\x84\xec\xe3\x5b\x46\x57\x4f\xb5\x82\xb8\x80\xe6\x2a\x1f\xc7\x18\xa9\xba\x5f\xeb\xb5\x61\xdb\xcd\x6a\xf6\x24\x5c\x14\xc7\x3b\x94\xb6\x71\xd1\xa0\x4b\x6b\x3f\xc7\x32\x5b\x6e\xf3\x4d\xa6\xa8\xea\x8e\xa2\xcf\x5a\xba\x1b\xe7\xd7\x71\xb2\x0e\x59\x13\xc8\xba\x2e\xb4\xc7\x5c\x5b\xb8\xd2\x93\x54\x61\x56\x3e\x28\x36\x46\x73\xd7\x12\xf6\x4f\x1d\xd5\x87\xca\x23\xad\xdd\xc7\x34\x24\xfd\x8b\x9b\x48\xa2\x8f\x22\x9c\x87\x4b\x01\xb5\x0f\x74\xf0\x0e\x0b\xce\x66\x71\x84\xe5\xd2\xb7\x9b\x89\x6b\x8d\xa6\x56\xbc\x20\x34\xdd\x1c\x8d\x12\x79\x72\x58\x7d\xb8\xdd\x23\x33\xaa\xd2\x4d\x2f\x1e\x1c\x94\xa6\xa9\xa7\x39\x77\x35\x44\xae\x2a\x34\xa1\x1b\x20\x70\x5c\xe9\x48\x5f\x52\xd3\x64\x6e\x05\x33\xc1\x4a\x32\x2c\x60\x83\x1a\xf7\x1d\x27\xde\xa1\xa0\x18\x96\x54\xea\x3f\x5c\x08\x96\x28\xbd\x4e\x51\xce\x3f\x7e\x70\xdd\x68\xfa\x4f\xcc\xd2\x5f\x6f\x6e\xae\xd2\x5f\x98\xe1\xb9\xf9\xac\xca\xcf\xee\xf1\xdd\xe5\xbb\xd7\x9f\xdd\x45\xe2\x80\xf9\x56\xdc\x07\x7c\x7d\x54\xb8\xd9\x70\x5b\xdf\x6d\x9d\xe5\xd3\xd6\x15\xfd\x39\xfc\xd0\x6e\x5e\xb5\xbf\x1e\x72\x06\x38\x07\xc1\x73\x94\xc6\xdd\xcb\x29\xed\x54\x64\x55\x0b\x89\xa9\x8b\x7b\x87\x42\x58\x65\xe2\xb0\x2c\xee\xbf\x89\xb7\x35\xbd\x99\x6a\x18\x95\xc4\x91\x12\x95\xa3\xae\xe0\x24\xa7\x09\x94\x8a\x30\xb5\xec\x18\xc0\xdb\xc0\xc9\xd0\x04\x88\x8d\xcb\x76\x2e\xf5\x00\x94\x9f\x3b\xa3\xab\x47\x47\x97\x7d\x6a\xb0\xea\xec\x7b\x04\xa7\x68\xee\x78\x5f\x2b\x57\x76\x7d\xbc\x7e\xeb\x07\x4e\x14\x0c\x2c\x56\xe0\xee\xbf\x00\xef\x2d\x6a\xca\x53\x86\x5b\x4c\xc6\xe0\xea\x22\x95\xed\xba\xd3\x40\x42\xe7\x0b\x42\xd2\x0e\xfe\x7e\x11\x2a\xa7\x33\x54\x1c\x24\xb6\x15\xa2\x4b\x3b\x9a\xaf\x39\x1d\xb5\x2f\x6d\x1d\xcc\x03\xd7\x7a\x1d\x78\x22\x96\xff\xc8\x10\xf8\x33\x3c\x61\x0c\xbc\x7a\x64\x7a\x0b\xaf\x0b\xde\x5e\x92\x76\xa6\x75\x0e\x4c\x62\x38\x6f\xef\x39\xe3\x2d\xb0\x31\x0d\xb6\x4d\x5e\xbc\xe0\x89\xb1\xd7\x07\x15\x5f\xe0\x55\x0f\x5c\x40\x7a\xe7\xf1\xa3\xa6\x10\xaa\x25\x45\x02\x16\x19\x89\xb4\x1a\x83\x65\xe3\x6f\x27\x7d\xbe\xf7\x65\x4b\x77\x04\x12\xf5\x37\xb0\xd5\x28\x57\xaf\x55\x72\x57\x90\x18\xbe\xc8\xa6\xca\xa8\x26\x7a\xe8\x9e\xaa\xbb\xef\xc2\xab\xee\xe8\xf4\x05\xad\x3f\x59\x2c\xe0\xf8\xbb\xd3\x17\xe9\xc9\x62\x31\x3d\x72\xb6\x35\x73\x44\x02\xe9\xf7\x9e\x32\x37\x70\xfa\xa2\x77\x1d\x38\xd0\x6c\x6f\x43\xbc\x20\x1e\x70\x5c\xb1\xfb\x51\xae\xb9\x69\xfb\xb2\xa0\x85\x2e\xff\xd0\x27\x73\x78\xcd\x16\x73\x56\xb8\xf6\xf7\x79\x4b\xf0\x8a\x5b\x2c\xe6\xe1\x34\x2c\x46\x09\x3f\x01\x8d\x8a\xdd\x77\x2c\xeb\x64\xb1\x18\x23\xe0\x9c\xd2\x19\x4c\x23\xc3\xc9\x51\xce\x1e\x85\x7d\xf6\x37\xe8\x4a\x37\xc9\xc5\xb8\xa4\x6d\x54\x6f\x37\x47\x50\x0f\x07\xb1\xa3\xc8\xcf\xc6\x36\x8e\xba\x55\x5f\xd3\xab\xbe\x22\x87\xcb\x3b\xc2\xac\x3a\x67\x0c\xbc\xf2\xdb\xe4\xbf\x01\x00\x00\xff\xff\xcb\x1e\x8e\xac\xc4\x23\x00\x00"

func metadataviewsCdcBytes() ([]byte, error) {
	return bindataRead(
		_metadataviewsCdc,
		"MetadataViews.cdc",
	)
}

func metadataviewsCdc() (*asset, error) {
	bytes, err := metadataviewsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "MetadataViews.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd0, 0x4, 0x17, 0x2e, 0x78, 0x55, 0xaf, 0x16, 0x60, 0xc2, 0x80, 0xe6, 0x72, 0x9f, 0x3, 0xdf, 0x52, 0xeb, 0x39, 0x68, 0xbb, 0x43, 0x17, 0xd6, 0x8c, 0xbd, 0x4e, 0xf2, 0xa5, 0xd7, 0x91, 0xc}}
	return a, nil
}

var _nonfungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x41\x8f\xdb\xba\x11\xbe\xeb\x57\xcc\xcb\x03\x9a\xdd\xc0\x6b\xf7\x50\xf4\x60\x20\x68\xda\xb7\x6f\x01\x5f\xb6\x0f\x5b\x17\x3d\x04\x01\x4c\x8b\x23\x9b\x08\x45\x2a\x24\x65\xc7\x0d\xf6\xbf\x17\x33\x24\x25\xca\xf6\x26\x9b\x5b\x73\x89\x57\x22\xbf\x99\xf9\xe6\x9b\x8f\xd4\xe2\xdd\xbb\xaa\xfa\xf5\x57\x58\xef\x11\x1e\xb4\x3d\xc2\xa3\x35\x77\x0f\xbd\xd9\xa9\xad\x46\x58\xdb\xcf\x68\xc0\x07\x61\xa4\x70\x92\x17\x6e\x1e\xad\xc9\xef\xf9\xf5\x06\x6a\x6b\x82\x13\x75\x00\x65\x02\xba\x46\xd4\x58\x55\x84\x37\xfc\x09\x61\x2f\x02\x08\xad\xc1\x58\x73\xd7\x64\xf4\xc0\xe8\x79\xb7\x87\xda\xf6\x5a\xd2\xdf\x8d\x75\x2d\x04\x3b\xaf\x56\x0d\x08\xe8\x3d\x3a\x38\x0a\x13\x3c\x04\x0b\x12\x3b\x6d\x4f\x20\xc0\xe0\x11\x4c\x13\x86\xfd\x33\x08\x7b\x54\x6e\xcc\xe6\xc8\x70\x06\x51\x56\xc1\x82\x6a\x3b\x8d\x2d\x9a\x40\xcb\xe0\xbc\x88\x31\xd7\x39\xe7\x7e\x89\xb3\x17\x07\xca\x18\x1a\xab\x89\x26\x2a\x86\x80\x5c\xaf\xd1\x83\x30\x12\x8c\x68\x95\xd9\x55\x5c\x6a\x98\x54\xef\x3b\xac\x55\xa3\xd0\xcf\x13\x83\x0f\xeb\x0d\x38\xf4\xb6\x77\x99\xaa\xda\x3a\x1c\x1e\x41\x38\x75\x89\x33\x87\x9d\x43\x8f\x54\xbb\x30\xf0\xf8\xb0\x06\x65\x18\xdd\xb7\xc2\x8d\xb5\x27\xe0\xdf\xac\xd6\x58\x07\x65\xcd\x06\x9e\x26\xf8\x23\x34\xa1\xfa\x60\x1d\x65\xcd\xd4\xbe\xf5\x8c\x5b\x0f\x7b\xe7\xd5\x8a\x5a\x59\xeb\x5e\xf2\xa2\x06\x8f\xd0\xf4\x86\xdf\x71\x0b\x04\x33\x40\x59\xd8\xa3\x41\x47\x8f\x50\x78\xa5\x4f\x55\x6b\x0f\xa9\xad\x9e\x12\x25\x5a\x6c\x1f\xc0\x36\xbc\xba\x0c\xc1\xf9\xfe\xe1\xec\x41\x49\x74\x1b\x5e\xb9\x79\xc2\x1a\xd5\x81\xfe\x1c\xd2\x1d\x48\xf4\x5c\x87\x2f\x9f\x80\xc4\x5a\x0b\x87\x45\x72\x47\x15\xf6\xe0\x6d\x8b\xd0\x39\x64\xd0\xce\x7a\xa6\x49\x2a\x5e\x51\x25\x56\xbf\xf4\xca\x21\x27\x35\x72\x56\x74\xb7\x46\x17\x84\x32\xa9\xa7\x0c\xb4\xc5\xbd\x38\x28\xeb\x86\x69\xf0\x51\x29\x27\xa0\x14\x3c\x76\xc2\x89\x80\xb0\xc5\x5a\xf4\x94\x66\x80\x9d\x3a\xa0\xe7\x18\xac\x60\xfa\x21\xb6\x4a\xab\x70\xa2\x48\x7e\x4f\xfb\x04\x38\x6c\xd0\xa1\xa9\x91\x44\x1a\x15\x5c\xa6\x44\xe9\x5a\xa3\x4f\x80\x5f\x3b\xeb\x13\x5e\xa3\x50\xcb\xa8\xba\xb1\x76\x65\xc0\x1a\x04\xeb\xa0\xb5\x0e\xab\xc4\xf9\x48\xd7\x1c\x56\x34\x83\xde\xa6\xc4\x28\x29\x7f\x9e\x55\x2b\x3e\x23\xd4\xbd\x0f\xb6\x1d\x9a\x90\x48\x9b\x0c\xd0\xb4\x11\x34\x96\x16\x0e\xc2\x29\xdb\x13\xa4\x32\xbb\xd4\x0b\x82\x8f\x7a\x98\x57\xd5\x3f\x4e\xd0\x7b\xe2\x73\x40\xe6\x12\x46\xa0\x59\x4a\xca\x36\x2c\xc9\xa9\xc6\x3d\xd4\xc2\x80\x47\x23\x2b\xda\xe5\xa2\x58\xb2\xda\x3a\x44\x77\x17\xec\x1d\xfd\x3f\xe3\xd8\x24\x3c\x6a\x99\xd9\x51\x7e\x1c\x84\xa7\x99\xd2\x12\x50\x23\xa1\x6a\xd0\x28\x77\xe8\xaa\x8b\x71\x5a\x5b\x0e\x95\xa7\x8e\x54\x6f\x6c\xd8\xa3\xe3\x14\x67\x83\x2d\xb1\x37\x78\xe2\xe6\xc4\xd0\xd2\x89\x38\x1a\x8f\x0f\xeb\xaa\x71\xb6\xbd\xe8\x29\xfb\x94\x81\x3a\x3b\x88\xc4\xce\x7a\x15\x86\x4e\x82\x35\x93\x58\x6f\x7d\x35\xd5\x68\x6d\xa9\x13\x21\xca\x37\x38\x61\x7c\x83\x6e\x5e\x55\xef\x16\x55\xb5\x58\xb0\x93\xb7\x24\xde\x38\xd5\xe7\xd6\x3c\x87\x7f\x32\x74\xf9\x96\x9a\xa5\x35\x6d\x56\x6d\x67\x5d\x88\x6d\x29\xfa\xad\x7c\xe1\xed\x8b\x45\xd5\xf5\xdb\x2b\xd0\x97\xae\xfa\xad\xaa\x00\x00\x52\x56\xc1\x06\xa1\xc1\xf4\xed\x16\x1d\x7b\x42\x6c\x1d\x2b\x55\xf9\xe8\x7a\xca\x00\x7e\x55\x3e\xf0\x44\xd0\x5e\x0a\x75\x10\x2e\x6e\xfe\x57\xdf\x75\xfa\xb4\x84\x7f\xaf\x4c\xf8\xeb\x5f\x06\xf0\xdf\x0f\x31\x4d\x11\x00\x5b\x15\x02\x4a\x38\x12\xc7\xa9\x0f\x45\xaa\x54\x87\x0a\x4a\x68\xf5\x5f\x94\x69\xfb\x10\x06\x19\xe6\xb7\xb4\x78\x35\x2e\xbc\xb9\xbd\x16\x4a\xf9\x69\x34\x91\x0e\x34\xe5\x07\x25\x98\x59\xde\xa7\x8c\x54\xb5\x08\xac\xc6\xc1\x38\x2f\x7c\x31\x01\x07\x38\x8a\x02\x04\x48\x47\xf3\x32\xdb\xc5\x02\x56\x17\x7b\x95\x07\x63\x43\xf4\x5d\x10\x75\x6d\x7b\x13\xde\x7a\x36\x7b\xb1\xc3\x19\x6c\x08\x66\xc3\xad\x86\x2d\xc2\xc6\x28\xbd\x99\x5f\xe7\xe0\x3f\x29\xf4\x8d\x92\x99\xec\x19\x67\xb1\x84\xbf\x4b\xe9\xd0\xfb\xbf\x5d\xa5\xe4\x25\x3e\x92\xc6\x51\xf2\x20\x4d\x0e\x82\xb3\xaa\x42\x66\x2a\x59\xdd\x6b\x88\x2a\xd1\x5f\x28\xe8\x3e\x2e\x99\xd4\x13\xec\xb5\x6a\x56\xd3\x4b\x4b\x92\x90\x1f\xce\xff\xf1\x7a\x72\x1e\xe9\xf2\xd0\x82\x15\xa9\xef\x1b\xaf\x28\xe6\xa0\x37\xea\x4b\x8f\xb0\xba\x4f\xa4\x89\x7a\xcf\x32\xdd\x0b\x3f\x2c\x25\x40\x8d\x01\xc6\x84\xf9\xd5\xf3\x90\xe7\x53\x3c\xc3\xda\x81\x7b\xf2\x93\x94\x1c\xa9\xec\x9a\x81\x52\x0d\x79\x3f\x5f\xa5\x1a\x65\xe2\x19\x94\x32\x27\x53\x42\x19\x1d\x8f\x30\x13\x1e\x3b\x3c\xd5\x72\x59\xeb\xe3\xc3\x7a\x79\x5e\xe6\x0f\x73\x2f\x38\xb6\xd0\xa2\x54\x74\x72\x66\xb9\x7b\xc8\xb6\x59\x98\xe6\x2b\xb8\xce\x97\x89\x29\xdf\x83\x27\x3b\xa4\xcb\xc9\x70\x8d\x1a\x62\x14\x9a\x22\xd7\x8b\x8b\x54\x80\x78\x1a\x47\x46\xdc\xa4\xb4\xa6\x37\x03\xec\x4d\xfe\xb1\xba\xcf\xb5\xde\x2e\xe1\xc3\x94\x0f\xde\x48\xf7\x90\xe9\x23\xfa\xe7\xd0\xf7\x3a\xcc\x95\x84\xf7\xef\xa1\xc4\x7a\x43\x42\x59\xdd\x67\xe5\x8f\x5e\x10\x67\xaa\xed\x7d\xa0\x21\xe6\xab\xa0\x68\x11\x44\x1c\x17\xba\xd9\xa0\xa7\x51\x58\xdd\xbf\x99\x44\x7b\xae\xa6\xbf\x7e\xd0\x8d\x34\x53\x3e\xf3\xf0\x53\xad\xc8\x17\xb9\xec\xff\x29\x50\x3e\xe9\x82\xf8\x3c\x36\x42\xf0\x2f\xe1\x76\x3d\x4b\x99\x7a\x20\xa4\x2c\x5b\x70\x16\xba\x08\x5f\x76\x24\x81\xdf\x30\x3f\xb1\x05\xb7\x2f\x17\xca\x03\x33\xb8\x64\x3a\xc6\x6b\xdb\xb6\x7c\xd7\xca\x1b\xba\x7e\xab\x95\xdf\x43\x63\xdd\xf0\x71\x31\xc9\xe5\x85\xfa\xc7\x8c\xff\x20\x84\xfa\x6c\x36\xbe\x9b\x6e\xb9\x68\x87\x61\x75\xef\x6f\x6e\x97\xf0\x31\x6a\xeb\xd3\xc5\x92\xad\x75\xce\x1e\x1f\x1f\xd6\x85\xb5\xdd\x2e\xe1\x4f\x79\x58\xaf\x1b\x46\x2a\x28\x0d\x80\xa9\x1d\x5d\x27\x26\x9f\x1f\x85\x4d\x6c\x31\xdf\xb4\x65\xfe\xfa\x18\xee\x06\xe4\x34\xd9\x5f\x5e\x14\xc6\x48\xc7\x72\x98\xd2\xd9\x20\x92\xd9\x35\xba\x4a\xd9\xdc\x2b\x7e\x27\x1c\xdf\x50\xf7\x56\xcb\xd1\x95\x53\x3e\x57\x24\x92\xef\x0d\x74\x80\x48\x5a\xbb\x84\x0f\xdf\x22\x3f\x4b\xda\xfb\x5c\xfd\x5f\xd8\xc4\xf7\x06\x24\xce\xc7\xe5\x40\x8c\xb9\x78\x90\x03\x39\x25\xd0\xb0\x29\x44\x17\x49\x1b\x95\x04\xe1\x9c\x38\xbd\x4e\x8d\x25\x60\x54\x22\x38\x0c\xbd\x33\x69\x62\x9d\x38\x65\x7b\xa2\x77\x71\xa6\x1c\xe6\x9e\xd4\xd7\x7b\xf2\x82\xae\xcb\x60\x4f\x39\x4a\x52\x37\xca\xf1\x2b\x29\xde\xc4\xcb\x2f\xe1\x2b\x71\x16\x0b\xf0\x76\x3c\xbf\x63\x73\xf8\xf3\xc1\xa1\x90\x20\x45\x10\x4c\x11\xdf\xc1\x5b\x0c\x7b\x2b\xd3\xa9\xa3\xc2\xcf\x4c\xd8\xb9\xc7\x3b\xbc\x62\xf1\x1e\x75\x33\x1f\x54\xf8\x51\xc9\x4f\xf0\xcb\x7b\x30\x4a\x2f\xe1\x0d\x61\x48\x8b\xf1\xe2\xc6\xf7\xde\xcb\xaa\x7e\x79\xad\x8f\xd7\x0e\x45\xc0\xdf\xdb\x2e\x9c\x8a\x0f\x86\xf8\x94\x5b\x86\xf4\xea\xd2\xc9\x21\x7e\x4e\x45\xce\xcf\x25\x5d\x12\x79\x62\x0a\xed\x91\xe9\xf7\x55\x49\xd2\xd5\xd8\xd4\xe0\x0f\x45\x2a\x85\x0b\x5e\x9e\x86\xe9\x24\xcc\xd2\x98\x6b\x34\xbb\xb0\xa7\x63\xf1\xcf\xe9\x34\x8c\x31\x64\x39\x8a\xf9\x18\xe4\xca\x0a\xa2\x32\x35\xcf\xd5\xff\x02\x00\x00\xff\xff\x33\x4d\x81\x27\xe0\x12\x00\x00"

func nonfungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_nonfungibletokenCdc,
		"NonFungibleToken.cdc",
	)
}

func nonfungibletokenCdc() (*asset, error) {
	bytes, err := nonfungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NonFungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x61, 0xca, 0x9d, 0xaa, 0x66, 0x36, 0xdf, 0xbc, 0x51, 0xdb, 0x7b, 0x51, 0xd8, 0x3d, 0x6f, 0x4e, 0x9c, 0x8e, 0x50, 0x28, 0x7c, 0x18, 0x1d, 0x2, 0xb2, 0xc2, 0x2b, 0x26, 0xa1, 0xfe, 0x2d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"ExampleNFT.cdc":       examplenftCdc,
	"MetadataViews.cdc":    metadataviewsCdc,
	"NonFungibleToken.cdc": nonfungibletokenCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"ExampleNFT.cdc": {examplenftCdc, map[string]*bintree{}},
	"MetadataViews.cdc": {metadataviewsCdc, map[string]*bintree{}},
	"NonFungibleToken.cdc": {nonfungibletokenCdc, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
