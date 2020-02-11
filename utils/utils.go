package utils

import (
	"errors"
	"path/filepath"
	"runtime"
)

var (
	ErrFailedGetRoot = errors.New("can't get buzz root path")
)

func getUtilsPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	return file
}
func GetRootPath() (string, error) {
	cp := getUtilsPath()
	if cp == "" {
		return "", ErrFailedGetRoot
	}
	return filepath.Join(filepath.Dir(cp), ".."), nil
}
