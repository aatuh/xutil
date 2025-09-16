package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Returns absolute directory of the caller. Panics on failure.
func MustGetMyDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Error getting runtime caller information")
	}
	return filepath.Dir(filename)
}

// Returns absolute directory of the caller. Panics on failure.
func DirMustExist(path string) string {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		panic(fmt.Sprintf(
			"directory does not exist: %s",
			path,
		))
	}
	if err != nil {
		panic(fmt.Sprintf(
			"error while checking if directory exists: %s",
			err,
		))
	}

	return path
}

// Creates a directory. Panics on failure.
func MustCreateDir(directory string) string {
	err := os.MkdirAll(
		directory,
		os.FileMode(0755),
	)

	// Ignore "directory already exists" error
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	return directory
}

// Returns absolute file of the caller. Panics on failure.
func FileMustExist(path string) string {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		panic(fmt.Sprintf(
			"file does not exist: %s",
			path,
		))
	}
	if err != nil {
		panic(fmt.Sprintf(
			"error while checking if file exists: %s",
			err,
		))
	}
	return path
}
