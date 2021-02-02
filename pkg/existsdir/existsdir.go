package existsdir

import (
	"fmt"
	"log"
	"os"
)

// ExistsDir returns true/false depending if the directory
// exists or not.
func ExistsDir(dir string) bool {
	var exists bool

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal("Error:", err)
		}

		exists = false
	}

	if _, err := os.Stat(dir); os.IsExist(err) {
		fmt.Println("The directory already exists... exiting...")
		os.Exit(3)

		exists = true
	}

	return exists
}
