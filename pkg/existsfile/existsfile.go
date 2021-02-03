package existsfile

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ExistsFile returns true/false depending if the file
// exists or not.
func ExistsFile(file string) bool {
	var exists bool

	_, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("The config.json file doesn't exist, be sure to run ~bup start~ before ~bup add~")
		os.Exit(3)

		exists = false
	} else {
		exists = true
	}

	return exists
}
