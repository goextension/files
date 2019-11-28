package files

import (
	"fmt"
	"os"
	"path/filepath"
)

const MaxDepth = 255

func List(path string, ext string, depth int) (files []string, e error) {
	path, e = filepath.Abs(path)
	if e != nil {
		return nil, e
	}
	info, e := os.Stat(path)
	if e != nil {
		return nil, fileWrap(e, "fileinfo")
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
			return nil, fileWrap(e, "dir")
		}
		var fullPath string
		for _, name := range names {
			fullPath = filepath.Join(path, name)
			if depth > 0 {
				ss, e := List(fullPath, ext, depth-1)
				if e != nil {
					return nil, fileWrap(e, "list")
				}
				files = append(files, ss...)
			}
		}
		return files, nil
	}

	if ext != "" && filepath.Ext(path) != ext {
		return nil, nil
	}

	return append(files, path), nil
}

func fileWrap(e error, msg string) error {
	return fmt.Errorf("%s:%w", msg, e)
}
