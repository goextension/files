package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const MaxDepth = 255

func List(path string, ext string, depth int) (files []string, e error) {
	path, e = filepath.Abs(path)
	if e != nil {
		return nil, e
	}
	info, e := os.Stat(path)
	if e != nil {
		return nil, errWrap(e, "fileinfo")
	}

	if info.IsDir() {
		file, e := os.Open(path)
		if e != nil {
			//Ignore error
			return nil, nil
		}
		defer file.Close()
		names, e := file.Readdirnames(-1)
		if e != nil {
			return nil, errWrap(e, "dir")
		}
		var fullPath string
		for _, name := range names {
			fullPath = filepath.Join(path, name)
			if depth > 0 {
				ss, e := List(fullPath, ext, depth-1)
				if e != nil {
					return nil, errWrap(e, "list")
				}
				files = append(files, ss...)
			}
		}
		return files, nil
	}

	if ext != "" && !compareExt(filepath.Ext(path), ext) {
		return nil, nil
	}

	return append(files, path), nil
}

func compareExt(pathext, ext string) bool {
	exts := strings.Split(ext, ",")
	for _, e := range exts {
		if pathext == e {
			return true
		}
	}
	return false
}

func errWrap(e error, msg string) error {
	return fmt.Errorf("%s:%w", msg, e)
}
