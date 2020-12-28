package diskstore

import (
	_g "io/ioutil"
	_gd "os"
	_b "strings"

	_gc "gitee.com/gooffice/gooffice/common/tempstorage"
)

// Open opens file from disk according to a path
func (_e diskStorage) Open(path string) (_gc.File, error) { return _gd.Open(path) }

// RemoveAll removes all files in the directory
func (_f diskStorage) RemoveAll(dir string) error {
	if _b.HasPrefix(dir, _gd.TempDir()) {
		return _gd.RemoveAll(dir)
	}
	return nil
}

type diskStorage struct{}

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage() { _c := diskStorage{}; _gc.SetAsStorage(&_c) }

// TempFile creates a new temp file by calling ioutil TempFile
func (_cd diskStorage) TempFile(dir, pattern string) (_gc.File, error) {
	return _g.TempFile(dir, pattern)
}

// Add is not applicable in the diskstore implementation
func (_df diskStorage) Add(path string) error { return nil }

// TempFile creates a new temp directory by calling ioutil TempDir
func (_eg diskStorage) TempDir(pattern string) (string, error) { return _g.TempDir("", pattern) }
