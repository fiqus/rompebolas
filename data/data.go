// +build dev

package data

//go:generate go run -tags=dev ../assets_generate.go

import (
	"net/http"
	"path/filepath"
	"runtime"
)

func getAssetsDir() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return basepath + "/assets"
}

var Assets http.FileSystem = http.Dir(getAssetsDir())
