package getrootdir

import (
	"log"
	"os"
)

// GetRootDir returns the root_directory or the path
// where the command is being run.
func GetRootDir() string {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error:", err)
	}

	return rootDir
}
