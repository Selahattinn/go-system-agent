package walker

import (
	"os"
	"path/filepath"

	"github.com/Selahattinn/go-system-agent/pkg/model"
)

type Walker struct {
	rootDirectory string
}

func NewWalker(rootDirectory string) Walker {
	return Walker{rootDirectory: rootDirectory}
}

func (w *Walker) Walk() ([]*model.File, error) {
	var files []*model.File
	err := filepath.Walk(w.rootDirectory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			var file model.File
			file.Path = path
			file.Name = info.Name()
			file.Permision = info.Mode().Perm().String()
			file.Size = info.Size()
			file.Isdir = info.IsDir()
			file.LastUptade = info.ModTime().Unix()
			files = append(files, &file)
		}
		return nil
	})
	return files, err
}

func (w *Walker) GetRootDirectory() string {
	return w.rootDirectory
}
