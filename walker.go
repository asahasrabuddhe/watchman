package watchman

import (
	"os"
	"path/filepath"
)

type Walker struct {
	path        string
	directories []string
}

func NewWalker(path string) *Walker {
	return &Walker{path: path}
}

func (w *Walker) GetDirectories() ([]string, error) {
	err := filepath.Walk(w.path, w.walker)
	return w.directories, err
}

func (w *Walker) walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info == nil {
		info, err = os.Stat(path)
		if err != nil {
			return err
		}
	}

	mode := info.Mode()

	if mode.IsDir() {
		w.directories = append(w.directories, path)
	}

	return nil
}

