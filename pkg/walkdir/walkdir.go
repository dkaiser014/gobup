package walkdir

import (
	"os"
	"path/filepath"

	"bup/pkg/getrootdir"
)

// WalkDir returns a list of all the files
// in a specific directory.
func WalkDir(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(getrootdir.GetRootDir(), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	return files, err
}
