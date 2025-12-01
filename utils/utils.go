package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller info.")
	}

	dir := filepath.Dir(file)
	path := filepath.Join(dir, filename)

	data, err := os.ReadFile(path)
	Check(err)

	return strings.TrimSuffix(string(data), "\n")
}
