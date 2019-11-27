package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func List(path string, ext string, depth int) (files []string, e error) {
	path, e = filepath.Abs(path)
	if e != nil {
		return nil, e
	}
	info, e := os.Stat(path)
	if e != nil {
		return nil, fileWrap(e, "file info")
	}

	if info.IsDir() {
		file, e := os.Open(path)
		if e != nil {
			return nil, fileWrap(e, "open file")
		}
		defer file.Close()
		names, e := file.Readdirnames(-1)
		if e != nil {
			return nil, fileWrap(e, "read dir")
		}
		var fullPath string
		for _, name := range names {
			fullPath = filepath.Join(path, name)
			if depth > 0 {
				ss, e := List(fullPath, ext, depth-1)
				if e != nil {
					return nil, fileWrap(e, "list dir")
				}
				files = append(files, ss...)
			}
		}
	}

	if ext != "" && filepath.Ext(path) == ext {
		files = append(files, path)
	}

	return files, nil
}

func fileWrap(e error, msg string) error {
	return fmt.Errorf("%s:%w", msg, e)
}
