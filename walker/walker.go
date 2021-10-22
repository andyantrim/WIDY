package walker

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Walker struct {
	BaseDir string
}

func NewWalker(baseDir string) Walker {
	return Walker{
		BaseDir: baseDir,
	}
}

func (w Walker) Walk() (paths []string, err error) {
	// Check the base dir exists
	fileInfo, err := os.Stat(w.BaseDir)

	// If nothing exists, return error
	if os.IsNotExist(err) {
		return paths, err
	}

	// If it exists but is not a dir return an error
	if !fileInfo.IsDir() {
		err = errors.New("provided file is not a dir")
		return paths, err
	}

	//List the dirs under it
	files, err := ioutil.ReadDir(w.BaseDir)
	if err != nil {
		return paths, err
	}

	for _, f := range files {
		if f.IsDir() {
			location := filepath.Join(w.BaseDir, f.Name())
			paths = append(paths, location)
		}
	}
	// Parse and return

	return paths, err
}
