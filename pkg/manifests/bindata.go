// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/defaults/cluster-ingress.yaml (592B)
// assets/router/cluster-role-binding.yaml (329B)
// assets/router/cluster-role.yaml (713B)
// assets/router/deployment.yaml (1.674kB)
// assets/router/namespace.yaml (243B)
// assets/router/service-account.yaml (213B)
// assets/router/service-cloud.yaml (628B)
// assets/router/service-internal.yaml (512B)

package manifests

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

var _assetsDefaultsClusterIngressYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x31\x8f\x9b\x40\x10\x85\x7b\x7e\xc5\x93\x5d\x43\x74\x2d\x5d\x94\xa4\x38\x29\x45\xa4\x48\xe9\xc7\xcb\xc3\x4c\xbc\xcc\xae\x76\x07\x9f\x9c\x5f\x1f\x71\x40\x8a\x8b\x4b\x9e\x86\x6f\xbf\x79\x73\xc6\x97\xb8\x54\x67\x79\xb5\x6b\x61\xad\x78\x53\x9f\x30\x70\x94\x25\x3a\xee\x12\x17\xd6\xe6\x8c\x57\xab\x2e\x31\xb2\x20\x24\x1b\xf5\x8a\x9a\x19\x74\xd4\xb0\x8f\x40\x0a\x21\x39\x47\xe5\x00\x71\x94\xc5\x5c\x67\x76\xcd\x4d\x6d\xe8\x3f\xbc\xd1\x48\xd6\x5f\x2c\x55\x93\xf5\xd0\x2d\xeb\x52\xa6\xd5\x49\x47\xef\x34\x7d\xba\xbf\x48\xcc\x93\xbc\x34\x33\x5d\x06\x71\xe9\x1b\xc0\x64\x66\x7f\xa8\xed\xdf\x35\x4b\x60\x8f\x7f\x3f\xb7\x3b\xae\x4d\x99\x45\x3c\x95\x06\x18\xd5\x24\xea\x1f\x96\xba\x52\xce\xf8\x66\x75\x29\x84\x4f\xe2\x48\x16\x1f\xf0\x89\x38\xe6\x11\xc4\x30\x30\xd2\xf9\x9e\x1f\x4d\x84\x6d\x83\x43\x17\xe9\xf2\x9b\xc1\xbb\xff\xf0\xed\xf3\x85\x76\x4c\xbb\x63\x0e\xcd\x66\xad\x71\xb3\xda\xbb\xf9\x9a\x66\x51\x83\xd6\x27\x6d\x62\xa9\x6a\xd7\x77\x2d\xfd\x70\x8f\x55\xc4\xd2\xc0\x1f\x51\x02\x67\x9a\xaf\xd0\x2d\xfa\xc9\xc8\xe0\xa9\x6c\x09\x30\x8b\x87\xe9\xbb\x5c\x18\xeb\x11\x01\xa7\x75\xb2\x2d\x29\xb2\xbb\x2d\x17\x16\xa3\xb3\xae\xe2\x6f\xa9\xdc\x58\x4e\x3d\x4e\xa7\x06\x98\xf4\x3a\x7d\xbe\x8b\x46\xb9\x68\x54\x7f\x6c\x00\x7f\x64\xae\x37\x4e\xcb\xd0\xfc\x0d\x00\x00\xff\xff\xbb\x91\xf4\xc5\x50\x02\x00\x00")

func assetsDefaultsClusterIngressYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDefaultsClusterIngressYaml,
		"assets/defaults/cluster-ingress.yaml",
	)
}

func assetsDefaultsClusterIngressYaml() (*asset, error) {
	bytes, err := assetsDefaultsClusterIngressYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/defaults/cluster-ingress.yaml", size: 592, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xca, 0x98, 0x9d, 0x1a, 0xf7, 0xdc, 0x5f, 0x53, 0xcd, 0x4d, 0x65, 0xc7, 0xae, 0xa1, 0x5, 0x28, 0x1d, 0x15, 0x36, 0x5e, 0xcc, 0x9c, 0x45, 0x91, 0xbd, 0xca, 0x40, 0xa7, 0xa6, 0xbf, 0xbc, 0x13}}
	return a, nil
}

var _assetsRouterClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x31\x4e\xc4\x40\x0c\x45\xfb\x39\x85\x25\xea\x0c\xa2\x43\xd3\x01\x37\x58\x24\x7a\xef\xc4\xbb\x31\x49\xec\xc8\xf6\xa4\xe0\xf4\x28\x4a\x44\xc3\x4a\x29\x2d\xf9\xbf\xff\xfe\x13\xbc\xb3\xf4\x0e\x31\x10\x98\xb6\x20\x03\xd3\x89\x20\x14\x38\x1c\x3e\xc9\x56\xae\x04\x6f\xb5\x6a\x93\xc8\x69\x64\xe9\x0b\x7c\x4c\xcd\x83\xec\xa2\x13\x6d\x71\x96\x7b\xc2\x85\xbf\xc8\x9c\x55\x0a\xd8\x15\x6b\xc6\x16\x83\x1a\xff\x60\xb0\x4a\x1e\x5f\x3d\xb3\x3e\xaf\x2f\x69\xa6\xc0\x1e\x03\x4b\x02\x10\x9c\xa9\x80\x2e\x24\x3e\xf0\x2d\x3a\x96\xbb\x91\x7b\xb7\x9b\x24\x6f\xd7\x6f\xaa\xe1\x25\x75\xb0\x17\x1f\x3e\x87\xce\x1f\xe1\xf8\xdf\x4f\x5f\xb0\x3e\xa2\xa6\x6d\xd8\x85\x6e\x5b\xf1\xbf\x19\xe7\x32\x27\xf0\xdf\x00\x00\x00\xff\xff\x83\x13\xa9\xa6\x49\x01\x00\x00")

func assetsRouterClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterClusterRoleBindingYaml,
		"assets/router/cluster-role-binding.yaml",
	)
}

func assetsRouterClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsRouterClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/cluster-role-binding.yaml", size: 329, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x12, 0x9a, 0xeb, 0xd3, 0x79, 0x1a, 0xa6, 0x75, 0xb5, 0xda, 0x10, 0x70, 0xf9, 0x80, 0xd1, 0x51, 0x55, 0x98, 0x4d, 0x11, 0x98, 0x79, 0x5c, 0x46, 0x99, 0xa3, 0x39, 0x68, 0xe9, 0xa2, 0x72, 0x56}}
	return a, nil
}

var _assetsRouterClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x31\x4f\xc3\x40\x0c\x85\xf7\xfb\x15\x56\x99\x93\x8a\x0d\x65\x65\x60\xaf\x10\xbb\x73\x71\x89\xe9\xf5\x7c\xb2\x7d\xa9\xc4\xaf\x47\x49\x5a\xa8\x40\x0c\x45\x6c\xf7\x74\x7a\xf6\xf7\x9e\xef\xe0\x31\x55\x73\x52\xb0\x28\x85\x06\x50\x49\x04\x7b\x51\x50\xa9\x4e\x6a\x2d\x3c\x8f\x6c\x60\xa3\xd4\x34\x40\x4f\x80\x06\x4a\xe6\xca\xd1\x79\x5a\x64\x11\x33\xee\x13\xb5\xe1\xc0\x79\xe8\x2e\x13\x77\x92\x28\x60\xe1\x17\x52\x63\xc9\x1d\x68\x8f\xb1\xc5\xea\xa3\x28\xbf\xa3\xb3\xe4\xf6\xf0\x60\x2d\xcb\x76\xba\x0f\x47\x72\x1c\xd0\xb1\x0b\x00\x19\x8f\xd4\x81\x14\xca\x36\xf2\xde\x1b\xce\xaf\x4a\x66\xcd\x8a\x14\xb4\x26\xb2\x2e\x34\x80\x85\x9f\x54\x6a\xb1\xd9\xd4\xc0\x66\x13\x00\xd0\x5d\xb9\xaf\x4e\xbb\x0b\xa4\x64\xeb\x20\xd7\x94\x02\xcc\xe4\x52\x35\xd2\xd9\x41\x79\x28\xc2\xd9\x6d\x51\xf3\x5a\x2b\x18\x69\x95\x46\x3a\xf1\x2a\x26\xd2\xfe\x6c\x49\x6c\xbe\x3c\x4e\xe8\x71\xfc\x09\x31\xe7\xa3\xec\x1c\xaf\x03\xde\xca\xe5\x72\xa0\xac\x34\x31\x9d\xce\x2c\xb5\x7f\xa3\xe8\x18\x23\x99\x7d\x7d\x5c\x71\x45\x25\x74\xfa\xa5\x94\x66\xbd\x66\xfb\x59\xe9\x1f\x98\x96\x09\x37\x96\xf1\xcf\xcb\xb7\xe6\xe8\xf5\x1b\x43\x2d\xc3\x1c\xfc\x23\x00\x00\xff\xff\x13\x49\xf9\x92\xc9\x02\x00\x00")

func assetsRouterClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterClusterRoleYaml,
		"assets/router/cluster-role.yaml",
	)
}

func assetsRouterClusterRoleYaml() (*asset, error) {
	bytes, err := assetsRouterClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/cluster-role.yaml", size: 713, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x17, 0x6f, 0x7f, 0x61, 0x6d, 0xdb, 0x39, 0xe1, 0x3a, 0xf9, 0xf8, 0x54, 0x21, 0x25, 0xf8, 0x6f, 0x6c, 0xd2, 0x2, 0x5, 0x16, 0xf7, 0xdd, 0x81, 0x1f, 0xde, 0xf5, 0x6, 0x66, 0xa, 0x63, 0xd}}
	return a, nil
}

var _assetsRouterDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x41\x4f\xe3\x3c\x10\xbd\xf7\x57\x8c\xda\x73\x28\x7c\xa0\x4f\xbb\xbe\x55\x6d\x58\x55\xa2\x10\xb5\x81\x6b\x65\x9c\x29\xb5\xea\xd8\xd6\x78\x12\xd4\xfd\xf5\x2b\x37\x61\x49\xb2\x05\xb1\xb7\xf5\x29\xf2\xbc\xf7\xe6\x8d\x27\x33\x13\x58\xa0\x37\xee\x58\xa2\x65\x78\xd5\xbc\x87\x02\x77\xb2\x32\x0c\xb5\x34\x15\x86\xd1\x04\xe6\xa6\x0a\x8c\x04\x4b\xfb\x42\x18\x02\x04\x8f\x4a\xef\xb4\x6a\x11\x20\x09\x41\x7a\x6f\x34\x16\x20\x19\xa8\xb2\xac\x4b\xbc\x18\x1d\xb4\x2d\x44\x47\x7e\x24\xbd\x7e\x42\x0a\xda\x59\x11\x09\x61\x5a\x5f\x8d\x4a\x64\x59\x48\x96\x62\x04\x30\x81\x7b\x59\x22\xe8\x00\x01\xb9\x27\x05\x60\x65\x89\xc1\x4b\x85\x02\x9c\x47\x1b\xf6\x7a\xc7\x89\x6e\x1c\x8d\x00\x8c\x7c\x46\x13\xa2\x08\x44\x69\x01\xe4\x2a\x46\x1a\x45\xaf\xf1\x36\xa0\x41\xc5\x8e\x1a\x44\x29\x59\xed\xef\x3a\x94\x3e\x09\x80\xb1\xf4\x46\x32\xb6\xf0\x8e\xc7\x78\x4c\x8f\x39\xe4\x02\xbc\x25\x3d\x7d\x23\xd5\x5a\xe1\x4c\x29\x57\x59\x8e\xf5\xf5\xa0\x00\xd6\x15\xb8\xe9\xb9\x8b\x67\x1c\xaf\x13\x72\x06\x2f\x0e\xd5\x33\x92\x45\xc6\x70\xa1\xdd\xf4\xd5\xd1\x01\x69\x2c\x60\x3c\x6e\xc1\x9e\xb4\x23\xcd\xc7\xb9\x91\x21\x34\x09\xc2\x31\x30\x96\x89\x6a\x1a\x97\x28\xd2\xac\x95\x34\x2d\x41\x39\xcb\x52\x5b\xa4\x4e\x09\xc9\xe9\x81\x07\xd6\xe2\x99\x80\x2e\xe5\xcb\x07\x4d\x79\x3b\x27\x48\x56\x19\x93\x39\xa3\xd5\x51\xc0\x72\x77\xef\x38\x23\x0c\xb1\xed\xef\x38\xef\x88\x3b\x49\xdf\xd3\xee\x99\x7d\xe7\xba\xe3\x31\x73\xc4\x02\xbe\x5d\xf6\xa2\x9e\x1c\x3b\xe5\x8c\x80\x7c\x9e\x7d\x20\x17\x3e\xd3\xbb\xb9\xb9\xfe\x2b\xc1\xc0\x92\x3f\x15\xbc\xfa\x7e\xfd\xff\x97\x14\x27\xb0\x42\x7a\x19\x4c\xca\x7b\x18\x6d\x7d\xee\x7d\x36\xf9\x2c\xdf\x6c\xb3\x87\x75\xde\x4b\x72\x1a\x40\x01\xe3\x98\x7d\x7c\x86\xb6\x7e\x78\xcc\xd3\xf5\x76\x93\xae\x9f\x96\xf3\x74\x7b\x3f\x5b\xa5\x9b\x6c\x36\x4f\xcf\x89\x9c\x9b\xaa\xa1\xde\x22\xbd\x9d\x3d\xde\xe5\xdb\x79\xba\xce\x97\xb7\xcb\xf9\x2c\x4f\xb7\x8b\xe5\xfa\x9c\xdc\x14\x59\x4d\xfd\x41\x4f\xd9\x84\xa9\x27\x5d\x4b\xc6\x0e\xce\xe8\x1a\x2d\x86\x90\x91\x7b\x46\xd1\x13\xd0\x56\xb3\x96\x66\x81\x46\x1e\x37\xa8\x9c\x2d\x82\x80\xab\xfe\x0f\x10\x1b\xfc\x03\xb9\x4f\x04\xf0\x92\xf7\x02\xa6\x7b\x94\x86\xf7\x3f\x87\xc1\x73\x8d\x22\x94\x85\xfe\x17\x8c\xd4\xce\x54\x25\xae\xe2\x92\x18\x4c\x48\x19\xef\xb2\x46\xf0\xf3\x47\x85\xb6\x4d\xed\x06\x4f\x14\x12\xc7\x45\x3d\x44\xc5\xa2\x1f\xac\x39\x0a\x60\xaa\xde\x42\x8d\x81\xdf\xb9\x93\x2f\x68\x05\x54\xd4\xaf\xbd\x45\xaf\x5c\x81\x02\x6e\xfe\xbb\xec\xfd\xf8\x9b\x13\xfc\xcf\x15\x9f\x34\x43\xf0\x2b\x00\x00\xff\xff\x35\xab\x54\xc8\x8a\x06\x00\x00")

func assetsRouterDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterDeploymentYaml,
		"assets/router/deployment.yaml",
	)
}

func assetsRouterDeploymentYaml() (*asset, error) {
	bytes, err := assetsRouterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/deployment.yaml", size: 1674, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4b, 0x27, 0xdb, 0x59, 0xb5, 0x98, 0x2b, 0x77, 0x67, 0x46, 0x3d, 0x6a, 0xd4, 0xf7, 0xfa, 0x1, 0xd1, 0x4d, 0x33, 0x20, 0x8f, 0x35, 0xa8, 0x89, 0x2d, 0xfa, 0xda, 0xcf, 0x48, 0x3d, 0x24, 0xe0}}
	return a, nil
}

var _assetsRouterNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\xbd\x6e\xc3\x30\x0c\xc4\xf1\x5d\x4f\x71\x70\x67\xf7\x63\xd5\x43\x74\xec\xce\x44\xd7\x84\x88\x4c\x0a\x26\xad\xe7\x2f\x0c\x04\xed\xd0\xf9\x7e\xc0\xfd\x1f\x6a\xad\xe2\x53\x36\xc6\x90\x2b\x8b\x0c\xfd\xe2\x1e\xea\x56\x31\x3f\xca\xc6\x94\x26\x29\xb5\x00\x26\x1b\x2b\x7c\xd0\xe2\xae\xdf\xb9\xaa\xdd\x76\x46\x14\x40\xcc\x3c\x25\xd5\x2d\x4e\x88\x3f\xf4\xaa\xfe\x66\xde\xb8\x06\x3b\xaf\xe9\x7b\xc5\xb2\x14\xa0\xcb\x85\xfd\x89\x5f\x10\x4c\x4c\xe9\x07\x91\x0e\x99\xae\x0d\x8d\x83\xd6\xd4\x6e\x70\xc3\xe3\xb8\x10\xd2\x36\x8d\x33\x0c\x79\x97\x7c\x82\x38\xe7\xdf\x37\xc8\xd0\xf8\x1f\xb0\x1f\xb6\x76\x4e\xf6\x8a\xe5\x7d\x29\x3f\x01\x00\x00\xff\xff\x55\xe9\xbf\x8f\xf3\x00\x00\x00")

func assetsRouterNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterNamespaceYaml,
		"assets/router/namespace.yaml",
	)
}

func assetsRouterNamespaceYaml() (*asset, error) {
	bytes, err := assetsRouterNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/namespace.yaml", size: 243, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa7, 0xd4, 0xa1, 0x93, 0xf1, 0x1d, 0xdc, 0xa4, 0xc, 0x91, 0xe5, 0x54, 0xec, 0xa1, 0xd2, 0xb1, 0x34, 0x19, 0x90, 0xd7, 0x34, 0xdc, 0xce, 0xea, 0x67, 0xc7, 0x8a, 0x6c, 0xe6, 0x4b, 0xd1, 0x92}}
	return a, nil
}

var _assetsRouterServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\xb1\x4e\xc4\x30\x10\x84\xe1\xde\x4f\x31\xd2\xd5\x9c\x44\xeb\x8e\x92\x16\x24\x7a\xb3\x99\xbb\x5b\x91\x78\xcd\xee\x3a\x88\xb7\x47\x41\x29\xa7\x98\x5f\xdf\x05\x2f\x22\x36\x7b\xe2\x66\x0e\xb7\x99\xf4\x80\x38\x5b\x72\xc1\xe7\x2f\xf2\x41\xd8\xa0\xb7\x34\xbf\xe2\x35\xf1\xa3\xeb\x0a\xe7\xf7\x54\x27\x64\x9d\x91\x74\x84\xd8\xe0\x52\x2e\x18\xf4\x4d\x23\xd4\x7a\xc0\xb9\xfe\x57\xd2\xf0\x76\x84\x31\xdc\x84\x11\xda\xef\xd7\xf2\xa5\x7d\xa9\x78\xa7\xef\x2a\x3c\x0d\xa5\x0d\xfd\xa0\x1f\xef\x8a\xfd\xb9\x6c\xcc\xb6\xb4\x6c\xb5\x00\xbd\x6d\xac\x27\xf0\x9c\x31\x9a\xb0\x1e\xba\x1e\x0f\xbd\xe5\x93\xf6\xbb\x33\xa2\xfc\x05\x00\x00\xff\xff\x33\xdc\xda\x8c\xd5\x00\x00\x00")

func assetsRouterServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceAccountYaml,
		"assets/router/service-account.yaml",
	)
}

func assetsRouterServiceAccountYaml() (*asset, error) {
	bytes, err := assetsRouterServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-account.yaml", size: 213, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd0, 0xe3, 0x6, 0x3a, 0x88, 0x2e, 0x33, 0xe3, 0x24, 0xf0, 0xf0, 0xe9, 0x43, 0xc8, 0x46, 0x6c, 0x60, 0x9, 0x69, 0x84, 0x3, 0xd8, 0xc3, 0x80, 0xb, 0xab, 0x37, 0x13, 0xce, 0xf2, 0xeb, 0x60}}
	return a, nil
}

var _assetsRouterServiceCloudYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x6b\x14\x41\x10\x85\xef\xfd\x2b\x1e\xe4\x9c\xa0\x98\x83\xcc\x31\x39\x09\x41\x16\x5c\xbc\x57\x7a\x6a\x76\x9a\xf4\x54\x35\x55\x35\xab\xfb\xef\xa5\x7b\x37\xa0\x28\x1e\xe7\x31\xfd\xd5\x7b\xdf\x1d\x5e\x94\x66\x3c\x51\x25\xc9\x6c\xf8\xc6\x76\x2e\x99\x11\x8a\x56\x29\x33\x8a\x60\x31\x95\x80\x2e\x88\x95\x61\xba\x07\x5b\x8f\x73\xd5\x7d\x06\xcb\xb9\x98\xca\xc6\x12\xfe\x90\xee\xf0\x5c\x77\xef\x3f\x7c\x91\x93\xb1\x3b\xbc\x71\x2e\x4b\xc9\x38\x53\xdd\xd9\x41\xc6\xa0\xd6\x6a\xe1\x19\x14\xb0\x5d\xa2\x6c\xfc\x90\xde\x8a\xcc\xd3\xfb\xf9\x44\xad\x7c\x67\xf3\xa2\x32\xe1\xfc\x31\x6d\x1c\x34\x53\xd0\x94\x80\x3b\x7c\xa5\x8d\x51\x1c\xce\xf1\x07\x02\x10\xda\xd8\x1b\x65\x9e\xa0\x8d\xc5\xd7\xb2\xc4\x7d\xb9\x36\x49\x40\xa5\x57\xae\xde\x21\xe8\x1d\xa6\xdb\x98\xd4\x3b\xf6\x34\x2e\x8d\xa7\x21\xe4\xdd\x47\x02\x9c\x2b\xe7\x50\xfb\xfb\x59\xef\x72\x5c\x8b\x83\xaa\x2b\x56\xf2\x21\x88\x97\x85\xf3\xd0\xb5\x91\xbd\x15\x39\xe1\xe5\x09\x4d\xb5\x22\xc8\x4e\x1c\x0e\x72\xec\xb2\x32\xd5\x58\x2f\xf8\xb1\xb2\x40\x74\xc0\x6e\x6e\x9b\xce\x57\x4f\xcd\xd8\xb9\xab\x17\x10\x44\x67\xc6\x2b\xaf\x45\xe6\x71\xc7\xaf\xaa\xfa\x6c\xfe\x19\x6c\x42\xf5\x68\xb4\x2c\x25\x1f\xb4\x96\x7c\xe9\x43\x32\xd5\x04\x34\xb5\x18\xab\xef\x87\xa0\x09\x6b\x44\x1b\x6b\x9a\x69\x68\xd6\x3a\xe1\xf8\x7c\xb8\x26\x6a\x31\xe1\xf3\x87\xf1\x71\x2d\x7c\x18\xd1\xed\xcd\xef\x08\xff\x2f\xe3\xf1\xf1\xd3\x3f\x21\x9e\x7e\x05\x00\x00\xff\xff\x14\xac\xd6\xf5\x74\x02\x00\x00")

func assetsRouterServiceCloudYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceCloudYaml,
		"assets/router/service-cloud.yaml",
	)
}

func assetsRouterServiceCloudYaml() (*asset, error) {
	bytes, err := assetsRouterServiceCloudYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-cloud.yaml", size: 628, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x35, 0x3d, 0xd7, 0x8, 0xf9, 0xba, 0x0, 0x54, 0xbd, 0x2a, 0xeb, 0x98, 0x83, 0x6f, 0x28, 0x5e, 0xda, 0xd8, 0xa9, 0x45, 0x65, 0xe2, 0x35, 0x98, 0xd0, 0x6, 0x64, 0xc4, 0x82, 0x36, 0x14}}
	return a, nil
}

var _assetsRouterServiceInternalYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\x31\x4f\x23\x41\x0c\x05\xe0\x7e\x7e\xc5\x93\x52\xe7\x74\x51\xa2\xd3\x31\x1d\x4a\x95\x06\x45\x02\xd1\x9b\x59\x27\xb1\x98\x9d\x19\xd9\xde\x20\xfe\x3d\xda\x0d\x11\x0b\x34\x29\xd7\xfb\xfc\xf9\xcd\x02\xdb\x3c\x98\xb3\xe2\x91\xf5\x2c\x89\xf1\x26\x7e\x42\xc7\x07\x1a\xb2\xe3\x4c\x79\x60\x0b\x5f\xa9\x5d\x39\x2a\x9b\xc1\x1a\x27\x39\x48\x02\x95\x52\x9d\x5c\x6a\x31\x90\x32\xa8\xb5\x2c\xdc\x81\x1c\x3a\x14\x97\x9e\xff\x84\x57\x29\x5d\xbc\x1e\x08\xd4\xe4\x99\xd5\xa4\x96\x88\xf3\x2a\xf4\xec\xd4\x91\x53\x0c\xc0\x02\x0f\xd4\x33\xa8\x74\xb8\xff\xe1\x1a\xfb\x37\x13\x28\xd4\xb3\x35\x4a\x1c\x51\x1b\x17\x3b\xc9\xc1\x97\x72\xe9\x17\x80\x4c\x2f\x9c\x6d\x54\x31\x96\x8a\xd0\x3a\x38\x6b\x18\x9b\x8f\x53\x7f\x6f\x1c\xaf\xef\xda\xed\x03\x60\x9c\x39\x79\xd5\xdf\x3b\x40\xab\xea\x13\xb6\x9c\xee\x46\x9c\xdc\xdb\x94\x1b\xff\x44\xfc\xff\x7b\xf9\xd0\xea\x35\xd5\x1c\xf1\xb4\xdd\x4f\x13\x27\x3d\xb2\xef\xa7\xd0\xe7\xce\x9c\xb0\x99\xb1\xd9\xac\x6f\x44\x6c\xa6\xf4\xec\x2a\x69\xee\xac\xee\xd6\xff\x6e\x80\xa6\xd8\x47\x00\x00\x00\xff\xff\x30\x30\x02\xf6\x00\x02\x00\x00")

func assetsRouterServiceInternalYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRouterServiceInternalYaml,
		"assets/router/service-internal.yaml",
	)
}

func assetsRouterServiceInternalYaml() (*asset, error) {
	bytes, err := assetsRouterServiceInternalYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/router/service-internal.yaml", size: 512, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0x3, 0x6d, 0x72, 0xb, 0x1e, 0xd8, 0x4, 0x82, 0xd2, 0xb1, 0x70, 0xa9, 0x3f, 0xf, 0x83, 0x3a, 0x2b, 0xeb, 0x18, 0x2b, 0x1e, 0xd2, 0xd6, 0xc3, 0xbe, 0x58, 0x72, 0xaa, 0xee, 0x2f, 0xe3}}
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
	"assets/defaults/cluster-ingress.yaml": assetsDefaultsClusterIngressYaml,

	"assets/router/cluster-role-binding.yaml": assetsRouterClusterRoleBindingYaml,

	"assets/router/cluster-role.yaml": assetsRouterClusterRoleYaml,

	"assets/router/deployment.yaml": assetsRouterDeploymentYaml,

	"assets/router/namespace.yaml": assetsRouterNamespaceYaml,

	"assets/router/service-account.yaml": assetsRouterServiceAccountYaml,

	"assets/router/service-cloud.yaml": assetsRouterServiceCloudYaml,

	"assets/router/service-internal.yaml": assetsRouterServiceInternalYaml,
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
	"assets": {nil, map[string]*bintree{
		"defaults": {nil, map[string]*bintree{
			"cluster-ingress.yaml": {assetsDefaultsClusterIngressYaml, map[string]*bintree{}},
		}},
		"router": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsRouterClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsRouterClusterRoleYaml, map[string]*bintree{}},
			"deployment.yaml":           {assetsRouterDeploymentYaml, map[string]*bintree{}},
			"namespace.yaml":            {assetsRouterNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml":      {assetsRouterServiceAccountYaml, map[string]*bintree{}},
			"service-cloud.yaml":        {assetsRouterServiceCloudYaml, map[string]*bintree{}},
			"service-internal.yaml":     {assetsRouterServiceInternalYaml, map[string]*bintree{}},
		}},
	}},
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
