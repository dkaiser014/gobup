package existsdir

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ExistsDir returns true/false depending if the directory
// exists or not.
func ExistsDir(dir string) bool {
	var exists bool

	_, err := ioutil.ReadDir(dir)

	if err != nil {
		err := os.Mkdir(dir, 0755)

		if err != nil {
			log.Fatal("Error:", err)
		}

		exists = false
	} else {
		fmt.Println("The directory already exists...")
		os.Exit(3)
		exists = true
	}

	return exists
}
